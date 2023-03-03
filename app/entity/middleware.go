package entity

type AppRequest struct {
	AppId  string `json:"app_id"`
	UserId string `json:"user_id"`
}

type AppResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
