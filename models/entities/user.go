package entities

import (
	"time"

	"github.com/jinzhu/gorm"
)

/*User 所有使用者基本資料
包含會員 代理 控盤人員
*/
type User struct {
	gorm.Model
	UserAncestorID      uint `sql:"default:null"`
	EncryptedPassword   string
	Username            string
	FirstName           string
	LastName            string
	ResetPasswordToken  string    `sql:"default:null"`
	ResetPasswordSentAt time.Time `sql:"default:null"`
	RememberCreatedAt   time.Time `sql:"default:null"`
	SignInCount         int
	CurrentSignInAt     time.Time `sql:"default:null"`
	LastSignInAt        time.Time `sql:"default:null"`
	CurrentSignInIP     string    `sql:"default:null"`
	LastSignInIP        string    `sql:"default:null"`
	BankID              uint      `sql:"default:null"`
	Tier                uint      `sql:"default:null"`
	Identity            string    `sql:"default:null"`
	ForkID              uint      `sql:"default:null"`
	AccessToken         string    `sql:"default:null"`
	Online              bool
	IsAdmin             bool
	Tags                string `sql:"not null"`
}

// CreateOperatorParams CreateAdmin所需參數
type CreateOperatorParams struct {
	Username string
	Password string
}

// UpdateOperatorParams UpdateAdmin所需參數
type UpdateOperatorParams struct {
	TargetID uint
	Password string
}

// DeleteOperatorParams RemoveAdmin所需參數
type DeleteOperatorParams struct {
	TargetID uint
}

// SignUpParams SignUp所需參數
type SignUpParams struct {
	UserName    string
	Password    string
	UserProfile UserProfileParams
	IsTier0     bool // 第0層產生默認資料，其餘則繼承
	IsAdmin     bool // admin不產生setting及allotter
}

// UserProfileParams SignUp user_profile所需參數
type UserProfileParams struct {
	NickName      string
	Note          string
	Quota         float64
	DelayOriginal int
	DelayRunning  int
	Parlay        int
	Status        string
	Accessable    bool
	CurrentQuota  float64
}
