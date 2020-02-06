package controllers

import (
	"api/models"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// UserController struct
type UserController struct {
	beego.Controller
}

// Login function
// @Title Login
// @Description Kullanicinin sisteme giris islemlerini yapar.
// @Param email query string true "kullanici eposta bilgisi"
// @Param password query string true "kullanici parola bilgisi"
// @Success 200 {object} models.User
// @router /login [post]
func (c *UserController) Login() {
	var email = c.GetString("email")
	var password = c.GetString("password")
	//json.Unmarshal(c.Ctx.Input.RequestBody, &login)

	fmt.Println("----> ", email)

	user, ok := models.CheckUser(email, password)

	var res models.JSONResult
	if ok {
		res.Status = true
		res.Error = ""
		res.Data = user
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
// @Param body body models.User true "kullanici model bilgisi"
// @Success 200 {object} models.User
// @router /duzenle [post]
func (c *UserController) Duzenle() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	var res models.JSONResult
	res.Status = true
	res.Error = ""
	res.Data = user

	c.Data["json"] = res
	c.ServeJSON()
}

// Ekle function
// @Title Ekle
// @Description Yeni kullanici bilgisi ekler.
// @Param body body models.User true "kullanici model bilgisi"
// @Success 200 {object} models.User
// @router /ekle [post]
func (c *UserController) Ekle() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	var res models.JSONResult
	res.Status = true
	res.Error = ""
	res.Data = user

	c.Data["json"] = res
	c.ServeJSON()
}

// Bilgi function
// @Title Bilgi
// @Description Kullanici bilgilerini getirir.
// @Param id path string true "Kullanici ID bilgisi"
// @Success 200 {object} models.User
// @router /bilgi/:id [get]
func (c *UserController) Bilgi() {
	var userID, _ = c.GetInt(":id")

	var user = models.GetUser(userID)

	var res models.JSONResult
	res.Status = true
	res.Error = ""
	res.Data = user

	c.Data["json"] = res
	c.ServeJSON()
}
