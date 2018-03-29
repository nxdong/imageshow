package routers

import (
	"github.com/pr1n4ple/imageshow/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
