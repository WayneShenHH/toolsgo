// Package bigquery gcp bigquery client
package bigquery

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/s535504/bqmigrate"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
)

// Conn bigquery connection
type Conn struct {
	ctx     context.Context
	client  *bigquery.Client
	dataset *bigquery.Dataset
}

// New instance bigquery Conn
func New(cfg environment.Conf) *Conn {
	ctx := context.Background()
	instance, err := bigquery.NewClient(ctx, cfg.BigQuery.ProjectID)
	if err != nil {
		panic(err)
	}
	return &Conn{
		ctx:     ctx,
		client:  instance,
		dataset: instance.Dataset(cfg.BigQuery.DatasetID),
	}
}

// Migrate to bigquery
func (conn *Conn) Migrate(versions []*bqmigrate.MigrateVersion) {
	m := bqmigrate.New(conn.client, conn.dataset, &bqmigrate.DefaultOption, versions)
	if err := m.Migrate(); err != nil {
		panic(err)
	}
}

// // Connect implement connect to bigquery service
// func (conn *Conn) Connect() {
// 	logger.Info("Connect to bigquery service")
// 	ctx := context.Background()
// 	instance, err := bigquery.NewClient(ctx, environment.Setting.BigQuery.ProjectID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// logger.Info(fmt.Sprintf(`bigquery connected %v`, conn.Connected()))
// 	conn.ctx = ctx
// 	conn.client = instance
// 	conn.dataset = instance.Dataset(environment.Setting.BigQuery.DatasetID)
// }

// Connected check bigquery is connected
func (conn *Conn) Connected() bool {
	return conn.client != nil
}

// Insert data to bigquery service
func (conn *Conn) Insert(data interface{}, tableName string) error {
	table := conn.dataset.Table(tableName)
	u := table.Uploader()
	return u.Put(conn.ctx, data)
}

// GetPriceUpdateLatestTs from bigquery
func (conn *Conn) GetPriceUpdateLatestTs() (ts int64, err error) {
	q := conn.client.Query(fmt.Sprintf(`
	SELECT
		Ts
	FROM
		%s.PriceUpdate
	ORDER BY Ts DESC
	LIMIT 1`, environment.Setting.BigQuery.DatasetID))

	it, err := conn.run(q)
	if err != nil {
		return
	}

	var latestTs struct{ Ts int64 }
	err = it.Next(&latestTs)

	return latestTs.Ts, err
}

// GetAntepostLatestTime from bigquery
func (conn *Conn) GetAntepostLatestTime(offerIDs string) (t []APLatestTime, err error) {
	q := conn.client.Query(fmt.Sprintf(`
	SELECT
		offerid,
		MAX(time) Time
	FROM
		%s.PersistenceAntepost
	WHERE offerid IN (%s)
	GROUP BY offerid`, environment.Setting.BigQuery.DatasetID, offerIDs))

	it, err := conn.run(q)
	if err != nil {
		return
	}

	var latestTimes = make([]APLatestTime, it.TotalRows)

	for i := it.StartIndex; i < it.TotalRows; i++ {
		if err = it.Next(&latestTimes[i]); err != nil {
			return
		}
	}

	return latestTimes, nil
}

// PrintErr detail error message
func PrintErr(err error) {
	if multiError, ok := err.(bigquery.PutMultiError); ok {
		for _, err1 := range multiError {
			for _, err2 := range err1.Errors {
				logger.Error(err2)
			}
		}
	} else {
		logger.Error(err)
	}
	os.Exit(1)
}

// executes a query then return row iterator
func (conn *Conn) run(q *bigquery.Query) (it *bigquery.RowIterator, err error) {
	job, err := q.Run(conn.ctx)
	if err != nil {
		return
	}
	status, err := job.Wait(conn.ctx)
	if err != nil {
		return
	}
	if err = status.Err(); err != nil {
		return
	}
	return job.Read(conn.ctx)
}
