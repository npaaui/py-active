package common

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
)

// 字段名称转换
type ValidTransfer interface {
	Translates() map[string]string
}

// 验证param参数提交
func ValidateParam(g *gin.Context, rule map[string]string, as interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	for k := range rule {
		if query, ok := g.Params.Get(k); ok && query != "" {
			query = strings.TrimSpace(query)
			data[k] = query
		}
	}

	return ValidateData(data, rule, as)
}

// 验证query参数提交
func ValidateQuery(g *gin.Context, rule map[string]string, as interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	for k := range rule {
		if query, ok := g.GetQuery(k); ok && query != "" {
			query = strings.TrimSpace(query)
			data[k] = query
		}
	}

	// 分页参数处理
	if rule["page"] != "" && rule["page_size"] != "" {
		if data["page"] == nil || data["page_size"] == nil {
			data["page"], data["page_size"] = "1", "10"
		}
		data["offset"], data["limit"] =
			PartPage(StrToInt(data["page"].(string), 1), StrToInt(data["page_size"].(string), 10))
		delete(rule, "page")
		delete(rule, "page_size")
		rule["limit"], rule["offset"] = "int", "int"
	}

	return ValidateData(data, rule, as)
}

// 验证 post form 参数提交
func ValidatePostForm(g *gin.Context, rule map[string]string, as interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	for k := range rule {
		if query, ok := g.GetPostForm(k); ok && query != "" {
			query = strings.TrimSpace(query)
			data[k] = query
		}
	}

	return ValidateData(data, rule, as)
}

// 验证post json数据
func ValidatePostJson(g *gin.Context, rule map[string]string, as interface{}) map[string]interface{} {
	data := map[string]interface{}{}
	jsonByte, err := g.GetRawData()
	if err != nil {
		panic(NewValidErr(err))
	}

	// 记录到请求日志
	ReqLogChan <- &ReqLogForChan{
		ReqNo: g.GetString("req_no"),
		Param: string(jsonByte),
	}

	err = json.Unmarshal(jsonByte, &data)
	if err != nil {
		panic(NewValidErr(err))
	}
	return ValidateData(data, rule, as)
}

// map数据校验
// @param rule map[string]string{
// 		"name": "string|required||字段名称",
// }
func ValidateData(data map[string]interface{}, format map[string]string, as interface{}) map[string]interface{} {
	var filter = map[string]string{}   // 过滤配置
	var rule = map[string]string{}     // 校验配置
	var transfer = map[string]string{} // 字段名转换配置
	filter, rule, transfer = getValidateMaps(format)

	// int, float类型值传了空字符串过不了filter，暂时做一下处理
	for k, v := range filter {
		if v == "int" && data[k] == "" {
			data[k] = 0
		}
		if v == "float" && data[k] == "" {
			data[k] = "0"
		}
	}

	defer func() {
		validErr := recover()
		switch validErr.(type) {
		case string:
			panic(NewValidErr(errors.New(validErr.(string))))
		case error:
			panic(NewValidErr(validErr.(error)))
		}
	}()

	// 参数校验
	va := validate.Map(data)
	va.FilterRules(filter)
	va.StringRules(rule)
	if len(transfer) > 0 {
		va.AddTranslates(transfer)
	} else if trans, ok := as.(ValidTransfer); ok {
		va.AddTranslates(trans.Translates())
	}

	if !va.Validate() {
		panic(NewValidErr(errors.New(va.Errors.One())))
	}

	// 参数赋值到结构体
	err := va.BindSafeData(&as)
	if err != nil {
		panic(NewValidErr(err))
	}
	return data
}

// 获取过滤配置，校验规则，字段名转换配置
func getValidateMaps(format map[string]string) (map[string]string, map[string]string, map[string]string) {
	var filter, rule, transfer = map[string]string{}, map[string]string{}, map[string]string{}
	for k, v := range format {
		tmpSlice := strings.Split(v, "||")
		if len(tmpSlice) == 0 {
			panic(NewValidErr(errors.New("校验配置有误")))
		}
		rule[k] = tmpSlice[0]
		typ := strings.Split(tmpSlice[0], "|")[0]
		if typ != "" {
			filter[k] = typ
		}
		if len(tmpSlice) == 2 {
			transfer[k] = tmpSlice[1]
		}
	}
	return filter, rule, transfer
}

func LoadPostFile(g *gin.Context, fileKey string, resType string) string {
	header, err := g.FormFile(fileKey)
	if err != nil {
		panic(NewValidErr(err))
	}
	fileName := header.Filename

	fileExt := path.Ext(fileName)
	fileTime := time.Now().Format("20060102150405")
	fileRand := IntToStr(rand.Intn(100))
	name := fileTime + fileRand + "[" + strings.TrimRight(fileName, "."+fileExt) + "]" + fileExt

	postFile, _ := header.Open()
	defer func() {
		err = postFile.Close()
		if err != nil {
			panic(NewSysErr(fmt.Errorf("postFile文件句柄关闭失败:%w", err)))
		}
	}()

	src := "upload/" + resType
	upload, err := os.Create(src + "/" + name)
	defer func() {
		err = upload.Close()
		if err != nil {
			panic(NewSysErr(fmt.Errorf("upload文件句柄关闭失败:%w", err)))
		}
	}()

	if os.IsNotExist(err) {
		// 判断upload文件夹是否存在
		_, err = os.Stat(src)
		if os.IsNotExist(err) {
			err = os.MkdirAll(src, os.ModePerm)
			if err != nil {
				panic(NewSysErr(fmt.Errorf("upload文件夹创建失败:%w", err)))
			}
			upload, _ = os.Create(src + "/" + name)
		} else {
			panic(NewSysErr(fmt.Errorf(name+"文件创建失败:%w", err)))
		}
	}

	_, err = io.Copy(upload, postFile)
	if err != nil {
		panic(NewSysErr(fmt.Errorf(name+"文件创建失败:%w", err)))
	}
	return resType + "/" + name
}

func PartPage(page, pageSize int) (offset, limit int) {
	if page < 0 {
		page = 1
	}
	if pageSize > 200 || pageSize < 1 {
		pageSize = 10
	}
	limit = pageSize
	offset = limit * (page - 1)
	return offset, limit
}
