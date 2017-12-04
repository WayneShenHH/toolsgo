package middleware

import (
	"github.com/WayneShenHH/toolsgo/repository"
	"github.com/WayneShenHH/toolsgo/repository/repositoryimpl"
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
)

// Store is a middleware function that initializes the datastore and attaches to
// the context of every request context.
func Store(cli *cli.Context) gin.HandlerFunc {
	v := setupStore(cli)
	return func(c *gin.Context) {
		repository.ToContext(c, v)
		c.Next()
	}
}

// setupStore is the helper function to create the datastore from the CLI context.
func setupStore(c *cli.Context) repository.Repository {
	// connectString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=%s",
	// 	c.String("database-user"),
	// 	c.String("database-password"),
	// 	c.String("database-host"),
	// 	c.String("database-name"),
	// 	c.String("database-charset"),
	// 	c.Bool("database-parse_time"),
	// 	c.String("database-loc"))
	// dbConfig := app.Configuration()
	// connectString := dbConfig.Afu
	return repositoryimpl.New()
}
