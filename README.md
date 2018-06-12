## beego admin

基于beego,bootstarp的一个后台管理系统


## 获取安装

您需要安装 Go 1.1+ 以确保所有功能的正常使用。

你需要安装或者升级 Beego 和 [Bee](http://beego.me/docs/install/bee.md) 的开发工具:

进入`GOPATH/src`目录

	$ go get github.com/fromiuan/cstore

进入cstore目录：

	$ bee run

## QS

1. 如果执行`bee run`失败,请确保安装上述bee工具(为了更加方便的操作，请将 `$GOPATH/bin` 加入到你的 `$PATH` 变量中)
2. 学习使用[beego](https://beego.me/docs/intro/) （官网http://beego.me ）获取更多帮助

## 初次使用

把项目中`cstore.sql`文件导入数据库,数据库名称为`cstore`,或者通过`conf/app.conf`修改数据库配置链接

默认用户名和密码都为`admin`

