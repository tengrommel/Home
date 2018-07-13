package main

import (
	_ "loveHome/routers"
	"github.com/astaxie/beego"
	_ "loveHome/models"
)

func main() {
	beego.Run()
}

