package env_config

import (
	"gin-api-frame/app/global/consts"
	"gin-api-frame/app/global/variable"
	"gin-api-frame/utils/env_config/env_config_intf"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

// 由于 vipver 包本身对于文件的变化事件有一个bug，相关事件会被回调两次
// 常年未彻底解决，相关的 issue 清单：https://github.com/spf13/viper/issues?q=OnConfigChange
// 设置一个内部全局变量，记录配置文件变化时的时间点，如果两次回调事件事件差小于1秒，我们认为是第二次回调事件，而不是人工修改配置文件
// 这样就避免了 vipver 包的这个bug

var lastChangeTime time.Time

type envConfig struct {
	viper *viper.Viper
}

func init() {
	lastChangeTime = time.Now()
}

// 创建一个yaml配置文件工厂
// 参数设置为可变参数的文件名，这样参数就可以不需要传递，如果传递了多个，我们只取第一个参数作为配置文件名
func CreateEnvFactory(fileName ...string) env_config_intf.EnvConfigInterf {

	viperObj := viper.New()
	// 配置文件所在目录
	viperObj.AddConfigPath(variable.BasePath)
	// 需要读取的文件名,默认为：config
	if len(fileName) == 0 {
		// todo 抽离出外部
		viperObj.SetConfigName(variable.EnvFile)
	} else {
		viperObj.SetConfigName(fileName[0])
	}
	//设置配置文件类型(后缀)为 yml
	viperObj.SetConfigType("env")

	if err := viperObj.ReadInConfig(); err != nil {
		log.Fatal(consts.ErrorsConfigInitFail + err.Error())
	}

	return &envConfig{
		viperObj,
	}
}

//监听文件变化
func (e *envConfig) ConfigFileChangeListen() {
	e.viper.OnConfigChange(func(changeEvent fsnotify.Event) {
		if time.Now().Sub(lastChangeTime).Seconds() >= 1 {
			if changeEvent.Op.String() == "WRITE" {
				lastChangeTime = time.Now()
			}
		}
	})
	e.viper.WatchConfig()
}

// 允许 clone 一个相同功能的结构体
func (e *envConfig) Clone(fileName string) env_config_intf.EnvConfigInterf {
	// 这里存在一个深拷贝，需要注意，避免拷贝的结构体操作对原始结构体造成影响
	var ymlC = *e
	var ymlConfViper = *(e.viper)
	(&ymlC).viper = &ymlConfViper

	(&ymlC).viper.SetConfigName(fileName)
	if err := (&ymlC).viper.ReadInConfig(); err != nil {
		//variable.ZapLog.Error(my_errors.ErrorsConfigInitFail, zap.Error(err))
	}
	return &ymlC
}

// Get 一个原始值
func (e *envConfig) Get(keyName string) interface{} {
	value := e.viper.Get(keyName)
	return value
}

// GetString
func (e *envConfig) GetString(keyName string) string {
	value := e.viper.GetString(keyName)
	return value
}

// GetBool
func (e *envConfig) GetBool(keyName string) bool {
	value := e.viper.GetBool(keyName)
	return value
}

// GetInt
func (e *envConfig) GetInt(keyName string) int {
	value := e.viper.GetInt(keyName)
	return value
}

// GetInt32
func (e *envConfig) GetInt32(keyName string) int32 {
	value := e.viper.GetInt32(keyName)
	return value
}

// GetInt64
func (e *envConfig) GetInt64(keyName string) int64 {
	value := e.viper.GetInt64(keyName)
	return value
}

// float64
func (e *envConfig) GetFloat64(keyName string) float64 {
	value := e.viper.GetFloat64(keyName)
	return value
}

// GetDuration
func (e *envConfig) GetDuration(keyName string) time.Duration {
	value := e.viper.GetDuration(keyName)
	return value

}

// GetStringSlice
func (e *envConfig) GetStringSlice(keyName string) []string {
	value := e.viper.GetStringSlice(keyName)
	return value
}
