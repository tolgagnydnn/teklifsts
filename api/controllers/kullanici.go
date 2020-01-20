package controllers

import (
	"api/models"

	"github.com/astaxie/beego"
)

// KullaniciController struct
type KullaniciController struct {
	beego.Controller
}

// Bilgi function
// @Title Bilgi
// @Description Kullanici bilgilerini getirir.
// @Success 200 {object} models.Kullanici
// @router / [get]
func (c *KullaniciController) Bilgi() {
	var kullanici = models.KullaniciBilgi()

	c.Data["json"] = kullanici
	c.ServeJSON()
}
