// Package router gin http router
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"

	"github.com/WayneShenHH/toolsgo/pkg/module/memcache"
	"github.com/WayneShenHH/toolsgo/pkg/module/mq"
	"github.com/WayneShenHH/toolsgo/pkg/transport/websocket/sample"
)

// router 設定 route 對應 endpoint
type router struct {
	cache    memcache.MemCache
	mq       mq.MessageQueueService
	sampleWS sample.Hub
}

// NewRouter 設定 route 對應 endpoint 建構子
func NewRouter(
	cache memcache.MemCache,
	mq mq.MessageQueueService,
	sampleWS sample.Hub,
) router.Router {
	return &router{
		cache:    cache,
		mq:       mq,
		sampleWS: sampleWS,
	}
}

// LoadServiceRouter 獨立載入 websocket 因為要從不同 container 啟動
func (r *Router) LoadServiceRouter(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	r.loadWebsocket(g)
	return g
}

// ProviderSet transport wire 建構子集合
var ProviderSet = wire.NewSet(
	NewRouter,
	sample.NewHub,
)
