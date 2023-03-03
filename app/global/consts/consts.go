package consts

const (
	//系统配置

	ErrorsConfigInitFail string = "初始化配置文件发生错误"

	// 表单验证器前缀
	ValidatorPrefix              string = "Form_Validator_"
	ValidatorParamsCheckFailCode int    = -400300
	ValidatorParamsCheckFailMsg  string = "参数校验失败"

	// CURD 常用业务状态码
	CurdStatusOkCode      int    = 0
	CurdStatusOkMsg       string = "Success"
	CurdCreatFailCode     int    = -400200
	CurdCreatFailMsg      string = "新增失败"
	CurdUpdateFailCode    int    = -400201
	CurdUpdateFailMsg     string = "更新失败"
	CurdDeleteFailCode    int    = -400202
	CurdDeleteFailMsg     string = "删除失败"
	CurdSelectFailCode    int    = -400203
	CurdSelectFailMsg     string = "查询无数据"
	CurdRegisterFailCode  int    = -400204
	CurdSelectErrorMsg    string = "查询失败"
	CurdRegisterErrorCode int    = -400205

	// 请求业务
	CallOtherServerErr int64 = -400206
	// 请求业务
	UnknownErrorMsg = "未知错误"
)
