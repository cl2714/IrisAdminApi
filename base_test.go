package main

import (
	"IrisApiProject/caches"
	"IrisApiProject/models"
	"flag"
	"fmt"
	"os"
	"testing"

	"IrisApiProject/config"
	"IrisApiProject/database"
	"github.com/iris-contrib/httpexpect"
	"github.com/kataras/iris"
	"github.com/kataras/iris/httptest"
)

var (
	app *iris.Application // iris.Applications
)

//单元测试基境
func TestMain(m *testing.M) {

	// 初始化app
	app = newApp()

	flag.Parse()
	exitCode := m.Run()

	// 删除测试数据表，保持测试环境
	database.DB.DropTable(
		&models.User{},
		&models.OauthToken{},
		&models.Role{},
		&models.Permission{},
	)

	os.Exit(exitCode)
}

// 单元测试 post 方法
func login(t *testing.T, url string, Object interface{}, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
	if Data != nil {
		e.POST(url).WithJSON(Object).
			Expect().Status(StatusCode).
			JSON().Object().Values().Contains(Status, Msg, Data)
	} else {
		e.POST(url).WithJSON(Object).
			Expect().Status(StatusCode).
			JSON().Object().Values().Contains(Status, Msg)
	}

	return
}

// 单元测试 post 方法
func create(t *testing.T, url string, Object interface{}, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
	at := GetLoginToken()

	ob := e.POST(url).WithHeader("Authorization", "Bearer "+at.Token).WithJSON(Object).
		Expect().Status(StatusCode).JSON().Object()

	ob.Value("status").Equal(Status)
	ob.Value("msg").Equal(Msg)

	for k, v := range Data {
		ob.Value("data").Object().Value(k).Equal(v)
	}

	return
}

// 单元测试 post 方法
func update(t *testing.T, url string, Object interface{}, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
	at := GetLoginToken()

	ob := e.PUT(url).WithHeader("Authorization", "Bearer "+at.Token).WithJSON(Object).
		Expect().Status(StatusCode).JSON().Object()

	ob.Value("status").Equal(Status)
	ob.Value("msg").Equal(Msg)

	for k, v := range Data {
		ob.Value("data").Object().Value(k).Equal(v)
	}

	return
}

// 单元测试 get 方法
func getOne(t *testing.T, url string, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
	at := GetLoginToken()
	if Data != nil {
		e.GET(url).WithHeader("Authorization", "Bearer "+at.Token).
			Expect().Status(StatusCode).
			JSON().Object().Values().Contains(Status, Msg, Data)
	} else {
		e.GET(url).WithHeader("Authorization", "Bearer "+at.Token).
			Expect().Status(StatusCode).
			JSON().Object().Values().Contains(Status, Msg)
	}

	return
}

// 单元测试 get 方法
func getMore(t *testing.T, url string, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
	at := GetLoginToken()
	if Data != nil {
		e.GET(url).WithHeader("Authorization", "Bearer "+at.Token).
			Expect().Status(StatusCode).
			JSON().Object().Values().Contains(Status, Msg, Data)
	} else {
		e.GET(url).WithHeader("Authorization", "Bearer "+at.Token).
			Expect().Status(StatusCode).
			JSON().Object().Values().Contains(Status, Msg)
	}

	return
}

// 单元测试 get 方法
func delete(t *testing.T, url string, StatusCode int, Status bool, Msg string, Data map[string]interface{}) (e *httpexpect.Expect) {
	e = httptest.New(t, app, httptest.Configuration{Debug: config.Conf.Get("app.debug").(bool)})
	at := GetLoginToken()

	e.DELETE(url).WithHeader("Authorization", "Bearer "+at.Token).
		Expect().Status(StatusCode).
		JSON().Object().Values().Contains(Status, Msg)

	return
}

/**
*设置测试数据表
*@param tn stirng 数据表名称
 */
func SetTestTableName(tn string) {
	err := caches.Cache.Set("test_table_name", tn, 0).Err()
	if err != nil {
		panic(err)
	}
}

/**
*登陆用户
*@return   Token 返回登陆后的token
 */
func GetLoginToken() models.Token {
	response, status, msg := models.CheckLogin(
		config.Conf.Get("test.LoginUserName").(string),
		config.Conf.Get("test.LoginPwd").(string),
	)

	// 打印错误信息
	if !status {
		fmt.Println(msg)
	}

	return response
}