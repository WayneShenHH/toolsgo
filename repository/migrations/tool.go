package migrations

import (
	"fmt"

	"github.com/WayneShenHH/toolsgo/repository"
)

type process struct {
	ID      int
	User    string
	Host    string
	DB      string
	Command string
	Time    int
	State   string
	Info    string
}

// KillProcess kill process too long
func KillProcess() {
	db := repository.DBConnect()
	ids := []process{}
	rows, _ := db.Raw(`show processlist;`).Rows()
	for rows.Next() {
		p := process{}
		rows.Scan(&p.ID, &p.User, &p.Host, &p.DB, &p.Command, &p.Time, &p.State, &p.Info)
		ids = append(ids, p)
	}
	for _, p := range ids {
		if p.Time > 200 {
			fmt.Println(p.ID, p.Time, p.State)
			db.Exec(`kill ?;`, p.ID)
		}
	}
}
