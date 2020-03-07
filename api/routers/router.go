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

	// Filter Token
	//beego.InsertFilter("/v1/user/profile", beego.BeforeRouter, filterToken)
}

/*
func filterToken(ctx *context.Context) {
	var authToken = ctx.Input.Header("Authorization")
	// Token seklen kontrol edilmeli

	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("TEKLIF_SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}
}
*/
