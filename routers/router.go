// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego/context"
	"github.com/yanluwei/qraphql-go/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.Get("/v1/version", func(ctx *context.Context) {
		ctx.Output.Body([]byte("固定路由"))
	})

	// 不写找方法名为GET的方法，动作与方法名对应
	//beego.Router("/v1/student",&controllers.StudentController{})
	beego.Router("/v1/student",&controllers.StudentController{},"get:ListStudent")

	beego.AddNamespace(ns)
}
