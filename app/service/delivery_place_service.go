package service

import (
	"gin-api-frame/app/entity"
	"gin-api-frame/app/global/variable"
	"gin-api-frame/app/model"
	"github.com/gin-gonic/gin"
)

type ExampleService struct {
	ctx *gin.Context
}

func ExampleServiceNew(ctx *gin.Context) *ExampleService {
	svc := ExampleService{ctx: ctx}
	return &svc
}

func (s *ExampleService) GetExampleDetail(req *entity.GetExampleReq) (*entity.GetExampleRes, error) {
	m := model.Example{}
	// 读取数据库获取详情信息
	example, err := m.GetExampleDetail(req.PlaceId)
	if err != nil {
		variable.AppLogger.Error(err)
		return nil, err
	}

	exampleRes := entity.GetExampleRes{
		Id:         example.Id,
		PlaceName:  example.PlaceName,
		PlacePhone: example.PlacePhone,
		Province:   example.Province,
		City:       example.City,
		County:     example.County,
		Detail:     example.Detail,
	}

	return &exampleRes, nil
}
