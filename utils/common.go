package utils

import (
	"archive/zip"
	"crypto/md5"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"gin-api-frame/app/global/variable"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const DateTimeFormat = "2006-01-02 15:04:05"

var fuckStrs = `·！@#￥%……&*——+=-【】、{}|；‘：”，。/《》？?><./';:"][{}\|+_-=*&^%$#@!` + "`"

func HostPort(host, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}

// StringTimeFormat 时间转换,string to string
func StringTimeFormat(timeString string) (string, error) {
	dateTime, err := time.Parse(time.RFC3339, timeString)
	// 如果有错误就不转换直接返回
	if err != nil {
		return timeString, err
	}
	dateTimeString := dateTime.Format(DateTimeFormat)
	return dateTimeString, nil
}

// []model 转 []byte再存入到mysql
func ModelToString(model interface{}) (string, error) {
	modelByte, err := json.Marshal(model)
	if err != nil {
		return "", err
	}
	return string(modelByte), nil
}

// Str2MD5 returns a MD5 hash in string form of the passed-in `s`
func Str2MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// StringConcat returns the concatenation of `base` and `strs` strings seperated by `sep`
func StringConcat(sep string, strs ...string) string {
	return strings.Join(strs, sep)
}

// StringContainsShit returns true if `str` contains any of the strings in `fuckStrs`
func StringContainsShit(str string) bool {
	for _, s := range fuckStrs {
		if strings.Contains(str, string(s)) {
			return true
		}
	}
	return false
}

// Void send everything passed in to the void and returns nil
func Void(args ...interface{}) {
}

// RecursiveZip recursively zips the specified directory into the specified destination zip file
func RecursiveZip(pathToZip, destinationPath string) error {
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		variable.AppLogger.Error(err)
		return err
	}
	myZip := zip.NewWriter(destinationFile)
	err = filepath.Walk(pathToZip, func(filePath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			variable.AppLogger.Error(err)
			return err
		}
		relPath := strings.TrimPrefix(filePath, filepath.Dir(pathToZip))
		zipFile, err := myZip.Create(relPath)
		if err != nil {
			variable.AppLogger.Error(err)
			return err
		}
		fsFile, err := os.Open(filePath)
		if err != nil {
			variable.AppLogger.Error(err)
			return err
		}
		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			variable.AppLogger.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		variable.AppLogger.Error(err)
		return err
	}
	err = myZip.Close()
	if err != nil {
		variable.AppLogger.Error(err)
		return err
	}
	return nil
}

//MyTime 自定义时间
type MyTime time.Time

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

func InArray(needle interface{}, hyStack interface{}) bool {
	switch key := needle.(type) {
	case string:
		for _, item := range hyStack.([]string) {
			if key == item {
				return true
			}
		}
	case int:
		for _, item := range hyStack.([]int) {
			if key == item {
				return true
			}
		}
	case int32:
		for _, item := range hyStack.([]int32) {
			if key == item {
				return true
			}
		}
	case int64:
		for _, item := range hyStack.([]int64) {
			if key == item {
				return true
			}
		}
	default:
		return false
	}
	return false
}
