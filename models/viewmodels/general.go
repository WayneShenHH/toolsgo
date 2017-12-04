package viewmodels

// APIResult Http API 回應外層
type APIResult struct {
	Success bool
	Code    uint
	Data    interface{}
	Message string
	// 在分頁的狀況下使用，回傳總資料筆數
	Count uint
}
