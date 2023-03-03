## 日志说明

#### 1.请求日志

日志级别：info

日志文件名称：request_logistics.log

#### 2.调试日志

日志级别：debug

日志文件名称：biz_logistics.log

注意：打日志的时候，请带上channel,方便后续快速从es中查询到该类别的日志，例如：
```go
variable.AppLogger.WithFields(logrus.Fields{
	"channel": "HealthCheck",
	"app_id": "11111111",
	"msg":    "cscscscsc",
}).Debug("c.ShouldBindJSO")
```
