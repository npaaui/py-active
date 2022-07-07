package wxhelper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//http请求
func HttpGet(url string) (res string, code int) {
	code = -1
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	code = resp.StatusCode
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	res = string(body)
	return
}

func JsonRpc(url string, method string, username, pass string, params []interface{}) string {
	requestParam := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	}
	bytesData, err := json.Marshal(requestParam)
	fmt.Println(string(bytesData))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	if username != "" {
		request.SetBasicAuth(username, pass)
	}
	request.Header.Set("Content-Type", "application/json;")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	//byte数组直接转成string，优化内存
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
	return string(respBytes)
}

//http请求参数
type HttpReq struct {
	Url      string
	IsPost   bool
	Data     []byte
	Timeout  int
	AuthUser string
	AuthPass string
	IsJson   bool
	Header   map[string]string
	UrlQuery map[string]string
}

func HttpRequest(q HttpReq) (respData []byte, err error) {
	var req *http.Request
	i := 0
	for {
		if i != 0 {
			time.Sleep(time.Millisecond * 300)
			fmt.Println("HttpRequest 启动 try ", i, " times ")
		}
		if i > 100 {
			return
		}
		i++
		method := "GET"
		if q.IsPost {
			method = "POST"
		}
		req, err = http.NewRequest(method, q.Url, bytes.NewReader(q.Data))
		if err != nil {
			return
		}
		req.Close = true
		if q.AuthUser != "" {
			req.SetBasicAuth(q.AuthUser, q.AuthPass)
		}
		for key, value := range q.Header {
			req.Header.Set(key, value)
		}
		if len(q.UrlQuery) > 0 {
			uq := req.URL.Query()
			for key, value := range q.UrlQuery {
				uq.Add(key, value)
			}
			req.URL.RawQuery = uq.Encode()
		}

		if q.IsJson {
			req.Header.Set("Content-Type", "application/json")
		} else if q.IsPost {
			req.Header.Set("Content-Type", "x-www-form-urlencoded")
		}
		timeout := 30
		if q.Timeout > 0 {
			timeout = q.Timeout
		}
		client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
		resp, err1 := client.Do(req)
		if err1 != nil {
			fmt.Println("HttpRequest err:", err1)
			err = err1
			return
		}
		defer resp.Body.Close()
		respData, err = ioutil.ReadAll(resp.Body)
		//不是200，报错
		if resp.StatusCode != http.StatusOK {
			err = fmt.Errorf(method+" "+req.URL.String()+" http status:", resp.StatusCode, " ,res : ", string(respData))
		}
		return
	}
	return
}
