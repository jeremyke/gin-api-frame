package controller

import (
	"gin-api-frame/app/global/consts"
	"gin-api-frame/utils/response"
	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {
	response.Success(ctx, consts.CurdStatusOkMsg, "OK!!!")
}
