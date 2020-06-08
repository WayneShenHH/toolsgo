// Package firestore stores db connections
package firestore

import (
	"context"
	"fmt"
	"regexp"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/option"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/module/nosql"
)

// New instance
func New() nosql.Handler {
	ctx := context.Background()
	cfg := environment.Setting.Nosql
	opt := option.WithGRPCConnectionPool(cfg.MaxConns)
	client, err := datastore.NewClient(ctx, cfg.ProjectID, opt)
	if err != nil {
		panic(err)
	}
	// Verify that we can communicate and authenticate with the firestore service.
	t, err := client.NewTransaction(ctx)
	if err != nil {
		panic(fmt.Errorf("firestoreDB: could not connect: %v", err))
	}
	if err := t.Rollback(); err != nil {
		panic(fmt.Errorf("firestoreDB: could not connect: %v", err))
	}
	return &firestoreDB{
		client: client,
	}
}

// firestoreDB persists entity to Cloud Datastore.
type firestoreDB struct {
	client *datastore.Client
}

func (db *firestoreDB) datastoreKey(entityName string, id int64) *datastore.Key {
	return datastore.IDKey(entityName, id, nil)
}

func (db *firestoreDB) datastoreKeys(entityName string, ids []int64) []*datastore.Key {
	keys := []*datastore.Key{}
	for _, id := range ids {
		keys = append(keys, db.datastoreKey(entityName, id))
	}
	return keys
}

// Get retrieves a entity by its ID. entity should be a pointer of entity, ex: &entity{}
func (db *firestoreDB) Get(entityName string, id int64, entity interface{}) error {
	ctx := context.Background()
	k := db.datastoreKey(entityName, id)
	if err := db.client.Get(ctx, k, entity); err != nil {
		return fmt.Errorf("firestoreDB: could not get entity: %v", err)
	}
	return nil
}

// Put saves a given entity, assigning it with ID. entity should be a pointer of entity, ex: &entity{}
func (db *firestoreDB) Put(entityName string, id int64, entity interface{}) error {
	ctx := context.Background()
	var err error
	k := db.datastoreKey(entityName, id)
	k, err = db.client.Put(ctx, k, entity)
	if err != nil {
		return fmt.Errorf("firestoreDB: could not put entity: %v", err)
	}
	return nil
}

// PutMulti saves a given slice of entity, input slice should be a slice of pointer, ex: []*entity{}
func (db *firestoreDB) PutMulti(entityName string, ids []int64, slice interface{}) error {
	ctx := context.Background()
	var err error
	keys := db.datastoreKeys(entityName, ids)
	keys, err = db.client.PutMulti(ctx, keys, slice)
	if err != nil {
		return fmt.Errorf("firestoreDB: could not put entity slice: %v", err)
	}
	return nil
}

// List returns a list of entities, input slice should be a pointer of slice, ex: &[]entity{}
func (db *firestoreDB) List(entityName string, filter map[string]interface{}, slice interface{}) error {
	ctx := context.Background()

	query, err := parseQuery(entityName, filter)
	if err != nil {
		return fmt.Errorf("firestoreDB: could not parseQuery: %v", err)
	}

	_, err = db.client.GetAll(ctx, query, slice)
	if err != nil {
		return fmt.Errorf("firestoreDB: could not list entities: %v", err)
	}

	return nil
}

// Delete removes a given entity by its ID.
func (db *firestoreDB) Delete(entityName string, id int64) error {
	ctx := context.Background()
	k := db.datastoreKey(entityName, id)
	if err := db.client.Delete(ctx, k); err != nil {
		return fmt.Errorf("firestoreDB: could not delete entity: %v", err)
	}
	return nil
}

// Delete removes a given entities by its IDs.
func (db *firestoreDB) DeleteMulti(entityName string, ids []int64) error {
	ctx := context.Background()
	keys := db.datastoreKeys(entityName, ids)
	if err := db.client.DeleteMulti(ctx, keys); err != nil {
		return fmt.Errorf("firestoreDB: could not delete entities: %v", err)
	}
	return nil
}

func parseQuery(entityName string, filter map[string]interface{}) (*datastore.Query, error) {
	query := datastore.NewQuery(entityName)
	for op, val := range filter {
		if ok, _ := regexp.MatchString("^[a-z,A-Z,0-9]+ (=|>=|<=|>|<)$", op); !ok {
			return nil, fmt.Errorf("input operator invalid: %v", op)
		}
		query = query.Filter(op, val)
	}
	return query, nil
}
