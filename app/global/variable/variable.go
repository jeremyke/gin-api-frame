package variable

import (
	"gin-api-frame/utils/app_log"
	"gin-api-frame/utils/env_config/env_config_intf"

	"github.com/spf13/cobra"
)

var (
	BasePath   string                  // 定义项目的根目录
	LogPath    string                  // 定义项目的日志目录
	EnvFile    string                  // 定义配置文件名称
	HttpPort   string                  // 服务端口
	AppRunMode string                  // 程序运行模式
	DateFormat = "2006-01-02 15:04:05" // 配置文件键值缓存时，键的前缀

	//命令行参数
	RootCmd   = &cobra.Command{}
	ServerCmd = &cobra.Command{
		Use:   "run",
		Short: "启动web服务",
		Long:  "启动web服务",
		Run:   func(cmd *cobra.Command, args []string) {},
	}

	//日志
	AppLogger *app_log.AppLogger

	// 全局配置文件
	EnvConfig env_config_intf.EnvConfigInterf
)
