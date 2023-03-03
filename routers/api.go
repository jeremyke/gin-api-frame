package routers

import (
	"fmt"
	"gin-api-frame/app/http/controller"
	"gin-api-frame/app/http/middleware/cors_iddleware"
	"gin-api-frame/app/http/middleware/log_middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

//         .............................................
//
//                            _ooOoo_
//                           o8888888o
//                           88" . "88
//                           (| -_- |)
//                            O\ = /O
//                        ____/`---'\____
//                      .   ' \\| |// `.
//                       / \\||| : |||// \
//                     / _||||| -:- |||||- \
//                       | | \\\ - /// | |
//                     | \_| ''\---/'' | |
//                      \ .-\__ `-` ___/-. /
//                   ___`. .' /--.--\ `. . __
//                ."" '< `.___\_<|>_/___.' >'"".
//               | | : `- \`.;`\ _ /`;.`/ - ` : | |
//                 \ \ `-. \_ __\ /__ _/ .-` / /
//         ======`-.____`-.___\_____/___.-`____.-'======
//                            `=---='
//
//         .............................................
//							永 不 报 错

// InitApiRouter ...
func InitApiRouter(test bool) *gin.Engine {
	router := gin.Default()

	//全局中间件
	if !test {
		router.Use(log_middleware.AccessLog()) //日志中间件
		router.Use(cors_iddleware.Next())      //跨域中间件
	}

	//探针
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello gin-api-frame")
	})

	//健康检查
	healthApi := router.Group("/health/")
	{
		healthApi.GET("check", controller.HealthCheck)
	}

	// 业务路由
	exampleGroup := router.Group("/")
	{
		// 获取自提点详情
		exampleGroup.POST("/xx.get.detail/1.0.0", controller.GetExampleDetail)
	}

	// 输出 Git 佛祖
	fmt.Println("    .............................................")
	fmt.Println("")
	fmt.Println(`                       _ooOoo_`)
	fmt.Println(`                      o8888888o`)
	fmt.Println(`                      88" . "88`)
	fmt.Println(`                      (| -_- |)`)
	fmt.Println(`                       O\ = /O`)
	fmt.Println(`                   ____/'---'\____`)
	fmt.Println(`                 .   ' \\| |// '.`)
	fmt.Println(`                  / \\||| : |||// \`)
	fmt.Println(`                / _||||| -:- |||||- \`)
	fmt.Println(`                  | | \\\ - /// | |`)
	fmt.Println(`                | \_| ''\---/'' | |`)
	fmt.Println(`                 \ .-\__ '-' ___/-. /`)
	fmt.Println(`              ___'. .' /--.--\ '. . __`)
	fmt.Println(`           ."" '< '.___\_<|>_/___.' >'"".`)
	fmt.Println(`          | | : '- \'.;'\ _ /'';.'/ - ' : | |`)
	fmt.Println(`            \ \ '-. \_ __\ /__ _/ .-' / /`)
	fmt.Println(`    ======'-.____'-.___\_____/___.-'____.-'======`)
	fmt.Println(`                       '=---='`)
	fmt.Println(``)
	fmt.Println(`    .............................................`)
	fmt.Println("                     永 不 报 错\n ")

	// 开启服务
	fmt.Println("    --------------   起飞！！！   ---------------")
	return router
}
