package main

import (
	_ "cstore/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.RunSyncdb("default", false, false)
	orm.Debug = true
	beego.Run()
}
