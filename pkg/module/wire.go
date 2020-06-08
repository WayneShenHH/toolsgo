// Package module 外部相依介面
package module

import (
	"github.com/google/wire"

	"github.com/WayneShenHH/toolsgo/pkg/module/bigquery"
	"github.com/WayneShenHH/toolsgo/pkg/module/gsuite"
	"github.com/WayneShenHH/toolsgo/pkg/module/memcache"
	"github.com/WayneShenHH/toolsgo/pkg/module/mq/nsq"
	"github.com/WayneShenHH/toolsgo/pkg/module/nosql/firestore"
	"github.com/WayneShenHH/toolsgo/pkg/module/orm"
	"github.com/WayneShenHH/toolsgo/pkg/module/stomp"
)

// ProviderSet module 外部相依性 instance 建構子集合
var ProviderSet = wire.NewSet(
	nsq.New,
	firestore.New,
	orm.New,
	memcache.New,
	gsuite.New,
	bigquery.New,
	stomp.NewClient,
)
