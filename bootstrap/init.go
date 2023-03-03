package bootstrap

import (
	"gin-api-frame/app/global/variable"
	"gin-api-frame/app/model"
	"gin-api-frame/utils/app_log"
	"gin-api-frame/utils/cache"
	"gin-api-frame/utils/env_config"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	// 1.初始化常量（参见variable常量包），相关路径：app\global\variable\variable.go

	// 2.接收参数初始化变量
	variable.ServerCmd.Flags().StringVarP(&variable.BasePath, "base_path", "b", "./", "项目更目录")
	variable.ServerCmd.Flags().StringVarP(&variable.EnvFile, "env_file", "e", ".env", ".env文件名称")
	variable.ServerCmd.Flags().StringVarP(&variable.HttpPort, "http_port", "p", "26000", "http端口")
	variable.RootCmd.AddCommand(variable.ServerCmd)
	if err := variable.RootCmd.Execute(); err != nil {
		panic(err)
	}

	// 3.全局配置文件
	variable.EnvConfig = env_config.CreateEnvFactory()

	// 4.设置运行模式
	variable.AppRunMode = variable.EnvConfig.GetString("RUN_MODE")
	gin.SetMode(variable.AppRunMode)

	// 5.初始化全局日志句柄，并载入日志钩子处理函数
	variable.LogPath = "/data/logs/" + variable.EnvConfig.GetString("APP_NAME") + "/"
	if variable.AppRunMode == gin.DebugMode {
		variable.LogPath = variable.BasePath + "/storage/logs/"
	}
	variable.AppLogger = app_log.InitAppLog(variable.LogPath)

	// 6.根据配置初始化全局*gorm.Db
	err := model.InitDB()
	if err != nil {
		variable.AppLogger.Error(err)
	}

	// 7.初始化redis
	cache.InitRedis()

}
