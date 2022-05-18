package main

import (
	_ "fc-bee/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
