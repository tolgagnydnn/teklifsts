package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["api/controllers:KullaniciController"] = append(beego.GlobalControllerRouter["api/controllers:KullaniciController"],
        beego.ControllerComments{
            Method: "Bilgi",
            Router: `/bilgi/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:KullaniciController"] = append(beego.GlobalControllerRouter["api/controllers:KullaniciController"],
        beego.ControllerComments{
            Method: "Duzenle",
            Router: `/duzenle`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:KullaniciController"] = append(beego.GlobalControllerRouter["api/controllers:KullaniciController"],
        beego.ControllerComments{
            Method: "Ekle",
            Router: `/ekle`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:KullaniciController"] = append(beego.GlobalControllerRouter["api/controllers:KullaniciController"],
        beego.ControllerComments{
            Method: "Giris",
            Router: `/giris`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["api/controllers:TeklifController"] = append(beego.GlobalControllerRouter["api/controllers:TeklifController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
