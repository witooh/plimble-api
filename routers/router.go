package routers

import (
	"github.com/witooh/plimble-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.RESTRouter("/object", &controllers.ObjectController{})
}
