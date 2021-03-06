<h1 align="center">IrisAdminApi</h1>

<div align="center">
    <a href="https://travis-ci.org/snowlyg/IrisAdminApi"><img src="https://travis-ci.org/snowlyg/IrisAdminApi.svg?branch=master" alt="Build Status"></a>
    <a href="https://codecov.io/gh/snowlyg/IrisAdminApi"><img src="https://codecov.io/gh/snowlyg/IrisAdminApi/branch/master/graph/badge.svg" alt="Code Coverage"></a>
    <a href="https://goreportcard.com/report/github.com/snowlyg/IrisAdminApi"><img src="https://goreportcard.com/badge/github.com/snowlyg/IrisAdminApi" alt="Go Report Card"></a>
    <a href="https://godoc.org/github.com/snowlyg/IrisAdminApi"><img src="https://godoc.org/github.com/snowlyg/IrisAdminApi?status.svg" alt="GoDoc"></a>
    <a href="https://github.com/snowlyg/IrisAdminApi/blob/master/LICENSE"><img src="https://img.shields.io/github/license/snowlyg/IrisAdminApi" alt="Licenses"></a>
    <h5 align="center">IrisAdminApi</h5>
</div>

项目功能升级迁移到 [snowlyg/go-tenancy](https://github.com/snowlyg/go-tenancy) （开发中）,本项目将停止更新。

#### 项目介绍
- `iris-go` 框架后台接口项目
- `gorm` 数据库模块 
- `jwt` 的单点登陆认证方式
- `cors` 跨域认证
- 数据支持 `mysql`，`sqlite3` 配置; `sqlite3` 需要下载 `gcc`, 并且在 `/temp` 目录下新建文件 `gorm.db` ,  `tgorm.db`。  [gcc 下载地址](http://mingw-w64.org/doku.php/download)
- 使用了 [https://github.com/snowlyg/gotransformer](https://github.com/snowlyg/gotransformer) 转换数据，返回数据格式化，excel 导入数据转换，xml 文件生产数据转换等 
- 增加了 `excel` 文件接口导入实例
- 前端采用了 `element-ui` 框架,代码集成到 `front` 目录
- 使用 `casbin` 做权限控制, `config/rbac_model.conf` 为相关配置。系统会根据路由名称生成对应路由权限，并配置到管理员角色。
- 增加系统日志记录 `/logs` 文件夹下，自定义记录，控制器内 `ctx.Application().Logger().Infof("%s 登录系统",aul.Username)`
- 增加多商户模式，分为管理端和商户端

 **注意：**
 - 默认数据库设置为 `DriverType = "Sqlite"` ，使用 mysql 需要修改为 `DriverType = "Mysql"` ,在 `config/conf.tml` 文件中
 - `permissions.xlsx` 权限导入测试模板文件，仅供测试使用; 权限会自动生成，无需另外导入。
 
 - 为了兼容测试 `files/file.go` 内增加了 `GetAbsPath()` 方法去获取配置文件的绝对路径，没有测试可以不需要此方法。
    如果你使用了此方法，需要将此方法中的路径修改为你项目的路径，而且你的项目必须要在 `gopath` 的 `/src` 目录下。
 
---

#### 项目开发过程详解

[Iris-go 项目登陆 API 构建细节实现过程](https://learnku.com/articles/39551)

---


#### 更新日志

[UPDATE](UPDATE.MD)
---

#### 问题总结

[ERRORS](ERRORS.MD)



#### 项目初始化

>拉取项目

```shell script
git clone https://github.com/snowlyg/IrisAdminApi.git

# github 克隆太慢可以用 gitee 地址:

git clone https://gitee.com/snowlyg/IrisAdminApi.git

```

>加载依赖管理包 (解决国内下载依赖太慢问题)
>使用国内七牛云的 go module 镜像。
>
>参考 https://github.com/goproxy/goproxy.cn。
>
>阿里： https://mirrors.aliyun.com/goproxy/
>
>官方： https://goproxy.io/
>
>中国：https://goproxy.cn
>
>其他：https://gocenter.io
>
>golang 1.13 可以直接执行：
```shell script
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

```

>项目配置文件 /config/conf.tml

```shell script
cp conf.tml.example conf.tml
```

>打包前端代码 
```shell script
 cd front               # 进入前端代码目录
 npm install            #加载依赖
 npm run-script build   #打包前端代码

 # 如果是开发前端代码,使用热加载
 npm run dev  
```


>运行项目 

[gowatch](https://gitee.com/silenceper/gowatch)
```shell script
go get github.com/silenceper/gowatch

# 安装 gowatch 后才可以使用
gowatch 

# go 命令
go run main.go 
```

---
##### 单元测试 
>http test

```shell script
 go test -v  //所有测试
 
 go test -run TestUserCreate -v //单个方法


// go get github.com/rakyll/gotest@latest 增加测试输出数据颜色

 gotest 
 
```

---

##### 接口文档
自动生成文档 (访问过接口就会自动成功)
因为原生的 jquery.min.js 里面的 cdn 是使用国外的，访问很慢。
有条件的可以开个 vpn ,如果没有可以根据下面的方法修改一下，访问就很快了
>打开 /resource/apiDoc/index.html 修改里面的

```shell script
https://ajax.googleapis.com/ajax/libs/jquery/2.1.3/jquery.min.js

国内的 cdn


https://cdn.bootcss.com/jquery/2.1.3/jquery.min.js
```

>访问文档，从浏览器直接打开 http://localhost:8081/apiDoc

---

#### 登录项目
- http://localhost:8080


#### 演示地址
[http://112.74.61.105:8087/](http://112.74.61.105:8087)


###### Iris-go 学习交流QQ群 ：676717248
<a target="_blank" href="//shang.qq.com/wpa/qunwpa?idkey=cc99ccf86be594e790eacc91193789746af7df4a88e84fe949e61e5c6d63537c"><img border="0" src="http://pub.idqqimg.com/wpa/images/group.png" alt="Iris-go" title="Iris-go"></a>

