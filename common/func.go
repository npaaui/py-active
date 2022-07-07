package common

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/shopspring/decimal"
)

/****************************
 * 时间相关
 ****************************/
func GetNow() (date string) {
	date = time.Now().Format("2006-01-02 15:04:05")
	return
}

func GetForever() (date string) {
	date = "2099-01-01 00:00:00"
	return
}

func GetBegin() (date string) {
	date = "0000-00-00 00:00:00"
	return
}

func GetTomorrowBegin() (date string) {
	date = time.Now().AddDate(0, 0, 1).Format("2006-01-02 00:00:00")
	return
}

func GetAfterHour(hour int) (date string) {
	date = time.Now().Add(time.Hour * time.Duration(hour)).Format("2006-01-02 15:04:05")
	return
}

/****************************
 * 加密相关
 ****************************/
// 获取hash
func GetHash(str string) (hash string) {
	h := md5.Sum([]byte(str))
	hash = fmt.Sprintf("%x", h)
	return
}

// 获取随机数
func RandNumString(length int) string {
	rand.Seed(time.Now().UnixNano())
	rs := make([]string, length)
	for start := 0; start < length; start++ {
		rs = append(rs, strconv.Itoa(rand.Intn(10)))
	}
	return strings.Join(rs, "")
}

func GetUniqueId() string {
	return Int64ToStr(UniqueIdWorker.GetId())
}

/****************************
 * 数据格式转换
 ****************************/
func StrToInt(str string, def int) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(fmt.Errorf("字符串「"+str+"」转整数失败: ", err))
		return def
	}
	return i
}

func StrToInt64(str string, def int64) int64 {
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println(fmt.Errorf("字符串「"+str+"」转整数失败: ", err))
		return def
	}
	return i
}

func Float64ToInt(f float64) int {
	i, _ := strconv.Atoi(fmt.Sprintf("%1.0f", f))
	return i
}

func IntToStr(i int) string {
	str := strconv.Itoa(i)
	return str
}

func Int64ToStr(i int64) string {
	str := strconv.FormatInt(i, 10)
	return str
}

func Float64ToString(f float64) string {
	str := strconv.FormatFloat(f, 'g', -1, 64)
	decimalNum, err := decimal.NewFromString(str)
	if err != nil {
		return str
	}
	str = decimalNum.String()
	return str
}

func StrToFloat64(s string, def float64) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println(fmt.Errorf("字符串「"+s+"」转浮点数失败: ", err))
		return def
	}
	return f
}

/****************************
 * 函数调用
 ****************************/
func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("the number of params is not adapted")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}


/****************************
 * 一些常用的工具函数
 ****************************/
// 数字+字母随机
func RandomStr(n int) string {
	seed := rand.NewSource(time.Now().UnixNano())
	letterIdxBits := uint(6)
	letterIdxMask := 1<<letterIdxBits - 1
	letterIdxMax := 63 / letterIdxBits
	letterBytes := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i, cache, remain := n-1, seed.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = seed.Int63(), letterIdxMax
		}
		if idx := int(cache & int64(letterIdxMask)); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}