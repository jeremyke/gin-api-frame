package model

import (
	"gin-api-frame/app/global/variable"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	MysqlCoreDb      *gorm.DB
	GoodsCenterDb    *gorm.DB
	GoodsServiceDb   *gorm.DB
	ExBusinessDb     *gorm.DB
	OrderDb          *gorm.DB
	ExampleDB        *gorm.DB
	OrdersAttachedDb *gorm.DB
)

func InitDB() error {
	var err = error(nil)
	ExampleDB, err = databaseConnect("DB_RW", "db_example")
	if err != nil {
		return err
	}
	return nil
}

// databaseConnect 连接指定数据库
func databaseConnect(connectName string, dbName string) (*gorm.DB, error) {
	dbLogLevel := logger.Silent
	if variable.AppRunMode == gin.DebugMode {
		dbLogLevel = logger.Info
	}
	_ = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  dbLogLevel,  // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
		},
	)

	dsn := setConnectInfo(connectName, dbName)
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	Logger: newLogger,
	//})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	// todo 添加连接池
	//sqlDB, _ := db.DB()
	////设置数据库连接池参数
	//sqlDB.SetMaxOpenConns(100)   //设置数据库连接池最大连接数
	//sqlDB.SetMaxIdleConns(20)   //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关

	return db, err
}

// setConnectInfo 设置db连接信息
func setConnectInfo(connectName string, dbName string) string {
	return variable.EnvConfig.GetString(connectName+"_USERNAME") + ":" +
		variable.EnvConfig.GetString(connectName+"_PASSWORD") + "@tcp(" +
		variable.EnvConfig.GetString(connectName+"_HOST") + ":" +
		variable.EnvConfig.GetString(connectName+"_PORT") + ")/" +
		dbName + "?charset=utf8&parseTime=True&loc=Local"
}
