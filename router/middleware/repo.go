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
	repo := repositoryimpl.New()
	return func(c *gin.Context) {
		repository.ToContext(c, repo)
		c.Next()
	}
}
