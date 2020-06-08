package jose

import (
	"time"
)

const (
	// AffiliateKind 代理之身份類別代號
	AffiliateKind uint8 = 1
	// PlayerKind 玩家之身份類別代號
	PlayerKind uint8 = 2
	// StationKind 站台之身份類別代號
	StationKind uint8 = 3
)

// UserToken JWT 中的簡單版 user
type UserToken struct {
	ID           uint64
	IdentityID   uint64
	StationID    uint64
	Username     string
	OwnerType    string
	LastSignInAt *time.Time
	LogoutURL    string `json:",omitempty"`
	Kind         uint8
	Status       IdentityStatus
}

// IdentityStatus 身份狀態定義格式
type IdentityStatus uint8

// Status 身份的狀態
const (
	ValidStatus     IdentityStatus = 1
	InvalidStatus   IdentityStatus = 2
	SupsupendStatus IdentityStatus = 3
	ClosedStatus    IdentityStatus = 4
)
