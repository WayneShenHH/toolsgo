// Package transport ...
package transport

import (
	"github.com/google/wire"
)

// ProviderSet transport wire 建構子集合
var ProviderSet = wire.NewSet(
	NewAPIServer,
	NewWebsocketServer,
)
