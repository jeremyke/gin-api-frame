package controller

import (
	"gin-api-frame/app/entity"
	"gin-api-frame/app/global/consts"
	"gin-api-frame/app/service"
	"gin-api-frame/utils/response"
	"github.com/gin-gonic/gin"
)

// GetExampleDetail 获取详情
func GetExampleDetail(c *gin.Context) {
	req := entity.GetExampleReq{}

	// 参数验证
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Fail(c, consts.ValidatorParamsCheckFailCode, err.Error(), err)
		return
	}

	// 调用service
	svc := service.ExampleServiceNew(c)
	data, err := svc.GetExampleDetail(&req)
	if err != nil {
		response.Fail(c, consts.CurdSelectFailCode, err.Error(), err)
		return
	}
	response.Success(c, consts.CurdStatusOkMsg, data)
}
