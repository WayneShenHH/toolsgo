//Package gsuite google gsuite service client
package gsuite

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	sheets "google.golang.org/api/sheets/v4"

	"github.com/WayneShenHH/toolsgo/pkg/environment"
	"github.com/WayneShenHH/toolsgo/pkg/module/logger"
)

// Client client for google sheet
type Client interface {
	// 取得試算表資料
	GetSheetRows(sheetID, sheetTitle string) (*[][]interface{}, error)
	GetSheetRowsToMap(sheetID, sheetTitle string) (*[]map[string]interface{}, error)
}

type client struct {
	sheetSvc *sheets.Service
}

// New gsuite client ctor
func New(c *environment.GSuiteConfig) Client {
	ctx := context.Background()
	googleShtsvc, err := sheets.NewService(ctx, option.WithCredentialsJSON([]byte(c.Cert)))
	if err != nil {
		logger.Fatal(fmt.Sprintf("Unable to retrieve Sheets Client %v", err))
		return nil
	}
	return &client{
		sheetSvc: googleShtsvc,
	}
}

func (c *client) GetSheetRows(sheetID, sheetTitle string) (*[][]interface{}, error) {
	resp, err := c.sheetSvc.Spreadsheets.Values.Get(sheetID, sheetTitle).Do()
	if err != nil {
		logger.Debug(fmt.Sprintf("GetRowsFromSheet Error %v", err))
		return nil, err
	}
	return &resp.Values, nil
}

func (c *client) GetSheetRowsToMap(sheetID, sheetTitle string) (*[]map[string]interface{}, error) {
	var m []map[string]interface{}
	resp, err := c.GetSheetRows(sheetID, sheetTitle)
	if err != nil {
		logger.Debug(fmt.Sprintf("GetSheetRowsToMap Error %v", err))
		return nil, err
	}
	rows := *resp
	cols := len(rows[0]) // 標題欄位數量
	for i := 1; i < len(rows); i++ {
		r := make(map[string]interface{})
		if cols != len(rows[i]) {
			logger.Fatalf(`row(%v) column counts not match with title`, i)
		}
		for j, title := range rows[0] {
			r[fmt.Sprintf("%v", title)] = rows[i][j]
		}
		m = append(m, r)
	}
	return &m, nil
}
