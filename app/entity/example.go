package entity

// GetExampleReq 请求结构体
type GetExampleReq struct {
	AppId   string `json:"app_id" binding:"required"`
	PlaceId int64  `json:"place_id" `
}

// GetExampleRes 获取详情的返回结构体
type GetExampleRes struct {
	Id         int64  `json:"place_id" `
	AppId      string `json:"app_id,omitempty"`
	PlaceName  string `json:"place_name"`
	PlacePhone string `json:"place_phone"`
	Province   string `json:"province"`
	City       string `json:"city"`
	County     string `json:"county"`
	Detail     string `json:"detail"`
	IsDefault  int8   `json:"is_default,omitempty" `
}
