package controllers

import (
	"api/models"
	"encoding/json"
	"time"

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
// @Success 200 {object} models.Token
// @router /login [post]
func (c *UserController) Login() {
	var email = c.GetString("email")
	var password = c.GetString("password")
	//json.Unmarshal(c.Ctx.Input.RequestBody, &login)

	user, ok := models.CheckUser(email, password)
	var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	var tokenModel = models.Token{
		UserID:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Token:     token,
	}

	var res models.JSONResult
	if ok {
		res.Status = true
		res.Error = ""
		res.Data = tokenModel
	} else {
		res.Status = false
		res.Error = "Eposta adresi veya parola hatalı. Lütfen kontrol ederek tekrar deneyin."
		res.Data = nil
	}

	c.Data["json"] = res
	c.ServeJSON()
}

// Update function
// @Title Update
// @Description Mevcut kullanici bilgisini düzenler.
// @Param firstName query string true "kullanici adı"
// @Param lastName	query string true "kullanici soyadı"
// @Param phone		query string true "kullanici telefon numarası"
// @Success 200 {object} models.User
// @router /update/:id [post]
func (c *UserController) Update() {
	//var user models.User
	//json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	var userID, _ = c.GetInt(":id")

	var user = models.User{
		ID:        int64(userID),
		FirstName: c.GetString("firstName"),
		LastName:  c.GetString("lastName"),
		Phone:     c.GetString("phone"),
	}

	err := models.UpdateUser(user)

	var res models.JSONResult
	if err != nil {
		res.Status = false
		res.Error = err.Error()
		res.Data = nil
	} else {
		res.Status = true
		res.Error = ""
		res.Data = nil
	}

	c.Data["json"] = res
	c.ServeJSON()
}

// Add function
// @Title Add
// @Description Yeni kullanici bilgisi ekler.
// @Param body body models.User true "kullanici model bilgisi"
// @Success 200 {object} models.JSONResult
// @router /add [post]
func (c *UserController) Add() {
	var user models.User
	json.Unmarshal(c.Ctx.Input.RequestBody, &user)

	user.Created = time.Now()
	userID, err := models.AddUser(user)

	var res models.JSONResult
	if userID > 0 {
		res.Status = true
		res.Error = ""
		res.Data = models.GetUser(int64(userID))
	} else {
		res.Status = false
		res.Error = "Kullanıcı kayıt işleminde hatalar oluştu. Hata: " + err.Error()
		res.Data = ""
	}

	c.Data["json"] = res
	c.ServeJSON()
}

// Get function
// @Title Get
// @Description Kullanici bilgilerini getirir.
// @Param id path string true "Kullanici ID bilgisi"
// @Success 200 {object} models.User
// @router /:id [get]
func (c *UserController) Get() {
	var userID, _ = c.GetInt(":id")

	var user = models.GetUser(int64(userID))

	var res models.JSONResult
	res.Status = true
	res.Error = ""
	res.Data = user

	c.Data["json"] = res
	c.ServeJSON()
}
