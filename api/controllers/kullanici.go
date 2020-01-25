package controllers

import (
	"api/models"
	"encoding/json"

	"github.com/astaxie/beego"
)

// KullaniciController struct
type KullaniciController struct {
	beego.Controller
}

// Giris function
// @Title Giris
// @Description Kullanicinin sisteme giris islemlerini yapar.
// @Param body body models.Kullanici true "kullanici model bilgisi"
// @Success 200 {object} models.Kullanici
// @router /giris [post]
func (c *KullaniciController) Giris() {
	var kullanici models.Kullanici
	json.Unmarshal(c.Ctx.Input.RequestBody, &kullanici)

	var res models.JSONResult
	if kullanici.Eposta == "demo@demo.com" && kullanici.Parola == "demo" {
		res.Status = true
		res.Error = ""
		res.Data = kullanici
	} else {
		res.Status = false
		res.Error = "Eposta adresi veya parola hatalı. Lütfen kontrol ederek tekrar deneyin."
		res.Data = nil
	}

	c.Data["json"] = res
	c.ServeJSON()
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

	var res models.JSONResult
	res.Status = true
	res.Error = ""
	res.Data = kullanici

	c.Data["json"] = res
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

	var res models.JSONResult
	res.Status = true
	res.Error = ""
	res.Data = kullanici

	c.Data["json"] = res
	c.ServeJSON()
}

// Bilgi function
// @Title Bilgi
// @Description Kullanici bilgilerini getirir.
// @Param id path string true "Kullanici ID bilgisi"
// @Success 200 {object} models.Kullanici
// @router /bilgi/:id [get]
func (c *KullaniciController) Bilgi() {
	var kullaniciID, _ = c.GetInt(":id")

	var kullanici = models.KullaniciBilgi(kullaniciID)

	var res models.JSONResult
	res.Status = true
	res.Error = ""
	res.Data = kullanici

	c.Data["json"] = res
	c.ServeJSON()
}
