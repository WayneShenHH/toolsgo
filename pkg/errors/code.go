// Package errors 錯誤處理
//nolint:unused // 預先定義錯誤類型接受沒有使用的定義，因此加入 nolint
package errors

// 預先定義的身份驗證錯誤代碼
const (
	Token Code = 2001 // Token 錯誤
)

// 預先定義的 Order 錯誤代碼
const (
	InsufficientQuota      Code = 4111 // 信用金餘額不足
	OddsChange             Code = 4112 // 賠率跳動
	LineChange             Code = 4113 // 球頭跳動
	ProportionChange       Code = 4114 // 比例跳動
	OnceLimit              Code = 4115 // 超過單注最大限額
	OfferLimit             Code = 4116 // 超過玩法限額
	SetLimit               Code = 4117 // 超過場次限額
	MatchLimit             Code = 4118 // 超過比賽限額
	PositionLimit          Code = 4119 // 超過部位限額
	ParamWrong             Code = 4120 // 參數錯誤
	PathWrong              Code = 4121 // 路徑錯誤
	NeedParam              Code = 4122 // 需要參數
	MatchQuota             Code = 4123 // 超過 Setting 單場限額
	OfferQuota             Code = 4124 // 超過 Setting 玩法限額
	MatchError             Code = 4125 // Match 狀態異常
	ItemCount              Code = 4126 // Item 數量異常
	ParlayCount            Code = 4127 // 串關數量異常
	ParlaySameMatch        Code = 4128 // 串關同比賽
	InsufficientBank       Code = 4129 // 銀行金餘額不足
	LockFail               Code = 4130 // 上鎖失敗
	TargetNotAccept        Code = 4131 // Offer 不接受下注對象
	ScoreChange            Code = 4132 // 比分跳動
	PositionStationLimit   Code = 4133 // 超過代理部位限額
	PositionPlayerLimit    Code = 4134 // 超過單使用者部位限額
	ScoreNotFound          Code = 4135 // 走地 Offer 無即時比分
	OfferError             Code = 4136 // Offer 狀態異常
	TargetUndefined        Code = 4137 // Target 未定義
	AmountTooLow           Code = 4138 // 下注金額過低
	OddsTypeWrong          Code = 4139 // 盤種錯誤
	RateModeWrong          Code = 4140 // 賠率模式錯誤
	OfferNotExist          Code = 4141 // Offer 不存在
	OddsNotExist           Code = 4142 // Odds 不存在
	OnceQuota              Code = 4143 // 超過 Setting 單注限額
	OrderTimeIsUp          Code = 4144 // 超過 Setting 賽前下注時間（單式）
	WinnableLimit          Code = 4145 // 超過單注可贏限額
	MultipleFail           Code = 4146 // 多筆注單下注時,當中有失敗
	DuplicateNaming        Code = 4147 // 名稱已經被使用過
	BeingUsed              Code = 4148 // 該目標被使用中
	OrderDataInvalid       Code = 4149 // 注單資料檢查不完整 會員，子站有誤
	CurrencyFail           Code = 4150 // 匯率找不到
	OddsTransferFail       Code = 4151 // 賠率轉換錯誤,不符合港盤規則
	JuiceParlaySettingFail Code = 4152 // 賠率轉換錯誤,水位不符規則
)

// 預先定義的錯誤代碼
const (
	BindFail    Code = 3001 // Gin 參數綁定失敗
	ServerError Code = 3002 // Server 錯誤
	GrpcFail    Code = 3003 // 呼叫grpc失敗
)

// 控盤賽事設定錯誤
const (
	MatchTeamDataErr Code = 3301 // 派彩隊伍對應錯誤
)

// 預先定義的 User 錯誤代碼
const (
	Validation   Code = 2401 // 驗證失敗
	UserNotExist Code = 2402 // 找不到 User
)

// 比對 & 賠率資料處理錯誤
const (
	SportCategoryNil    Code = 1201 // 聯盟分類為空
	SportLeagueNil      Code = 1202 // 聯盟為空
	SportGroupNil       Code = 1203 // 組別為空
	SportTeamNil        Code = 1204 // 隊伍為空
	SportMatchNil       Code = 1205 // 隊伍為空
	SportMatchTeamCount Code = 1206 // 賽事隊伍數量小於二
	BookMakerEnable     Code = 1207 // bookmaker 啟用錯誤
	SportOfferTypeMap   Code = 1208 // 運動玩法對應
	OfferMatchNotFound  Code = 1209 // 賽事還未被建立
	OfferMessageInvalid Code = 1210 // 訊息內容不正確
)

// 預先定義的注單風險部位檢查錯誤代碼
const (
	Server            Code = 8000 // 伺服器錯誤 (go-mts 專用)
	OrderException    Code = 8001 // MTS系統錯誤。(下注時，注單進風控系統檢查部位時發生錯誤。正常不會有此情況)
	OrderRejectSystem Code = 8002 // 系統自動拒單。(風控人員超過10秒無人處理，系統自動拒單)
	OrderRejectOp     Code = 8003 // 風控人員拒單。(由風控人員手動手動拒單)
)
