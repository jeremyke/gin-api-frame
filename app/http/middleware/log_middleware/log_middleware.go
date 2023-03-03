/**
* @Author: maxwell
* @description:日志中间件
* @Date: 2021/9/22 8:23 下午
 */
package log_middleware

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"gin-api-frame/app/entity"
	"gin-api-frame/app/global/variable"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		logFields := make(map[string]interface{}, 20)

		//请求前
		c.Writer = bodyWriter //将其赋予当前的writer写入流
		beginTime := time.Now().UnixNano() / 1e3
		logFields["begin_time"] = beginTime
		requestBody, _ := ioutil.ReadAll(c.Request.Body)
		//再重新写回请求体body中，ioutil.ReadAll会清空c.Request.Body中的数据
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))

		logFields["params"] = string(requestBody)
		logFields["client_ip"] = strings.Split(c.Request.RemoteAddr, ":")[0]
		logFields["target_url"] = c.Request.RequestURI
		requestBodyStruct := entity.AppRequest{}
		json.Unmarshal(requestBody, &requestBodyStruct)
		appId := requestBodyStruct.AppId
		userId := requestBodyStruct.UserId
		logFields["app_id"] = appId
		logFields["user_id"] = userId

		//请求后
		c.Next()
		responseBody := bodyWriter.body.String()
		logFields["content"] = responseBody
		responseBodyStruct := entity.AppResponse{}
		json.Unmarshal([]byte(responseBody), &responseBodyStruct)
		responseCode := responseBodyStruct.Code
		responseData := responseBodyStruct.Data
		responseMsg := responseBodyStruct.Msg
		// fmt.Println(responseBodyStruct)
		logFields["code"] = responseCode
		logFields["message"] = responseMsg
		logFields["data"] = responseData
		httpCode := bodyWriter.Status()
		logFields["http_code"] = httpCode
		endTime := time.Now().UnixNano() / 1e3
		logFields["end_time"] = endTime
		costTime := endTime - beginTime
		logFields["cost_time"] = costTime

		//记录日志
		variable.AppLogger.WithFields(logFields).Info("log middleware")
	}
}
