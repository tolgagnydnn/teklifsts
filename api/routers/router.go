// Package routers module
// @APIVersion 0.4.0
// @Title TeklifSTS Sistem API
// @Description TeklifSTS sistemi icin gerekli metodlari icerir.
// @Contact zafercelenk@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/offer",
			beego.NSInclude(
				&controllers.OfferController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
