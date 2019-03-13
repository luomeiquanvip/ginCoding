package handler

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)


var router *gin.Engine

func init()  {
	// 初始化路由
	router = gin.Default()
	router.GET("/", IndexApi)
}

func Get(url string, router *gin.Engine) []byte {
	// 构造get请求
	req := httptest.NewRequest("GET", url, nil)
	// 初始化响应
	w := httptest.NewRecorder()

	// 调用相应的handler接口
	router.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()
	defer result.Body.Close()

	// 读取响应body
	body,_ := ioutil.ReadAll(result.Body)
	return body
}

func TestIndexApi(t *testing.T) {
	// 初始化请求地址
	url := "/"

	// 发起Get请求
	body := Get(url, router)
	fmt.Printf("response:%v\n", string(body))

	// 判断响应是否与预期一致
	if string(body) != "It works" {
		t.Errorf("响应字符串不符，body:%v\n",string(body))
	}
}
