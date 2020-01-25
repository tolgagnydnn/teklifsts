package controllers

import (
	"api/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// KullaniciController struct
type KullaniciController struct {
	beego.Controller
}

// Duzenle function
// @Title Ekle
// @Description Yeni kullanici bilgisi ekler.
// @Param body body models.Kullanici true "kullanici model bilgisi"
// @Success 200 {object} models.Kullanici
// @router /duzenle [post]
func (c *KullaniciController) Duzenle() {
	var kullanici models.Kullanici
	json.Unmarshal(c.Ctx.Input.RequestBody, &kullanici)

	fmt.Println("--> Kullanici: ", kullanici)

	c.Data["json"] = kullanici
	c.ServeJSON()
}

// Ekle function
// @Title Ekle
// @Description Yeni kullanici bilgisi ekler.
// @Param body body models.Kullanici true "kullanici model bilgisi"
// @Success 200 {object} models.Kullanici
// @router /ekle [post]
func (c *KullaniciController) Ekle() {
	var kullanici models.Kullanici
	json.Unmarshal(c.Ctx.Input.RequestBody, &kullanici)

	fmt.Println("--> Kullanici: ", kullanici)

	c.Data["json"] = kullanici
	c.ServeJSON()
}

// Bilgi function
// @Title Bilgi
// @Description Kullanici bilgilerini getirir.
// @Param id path string true "Kullanici ID bilgisi"
// @Success 200 {object} models.Kullanici
// @router /bilgi/:id [get]
func (c *KullaniciController) Bilgi() {
	var kullaniciID = c.GetString(":id")

	fmt.Println("--> Kullanici ID: ", kullaniciID)

	var kullanici = models.KullaniciBilgi()

	c.Data["json"] = kullanici
	c.ServeJSON()
}
