package controllers

import (
	"api/models"

	"github.com/astaxie/beego"
)

// TeklifController struct
type TeklifController struct {
	beego.Controller
}

// GetAll function
// @Title GetAll
// @Description tum teklif kayitlarini getirir.
// @Success 200 {object} models.Teklif
// @router / [get]
func (c *TeklifController) GetAll() {
	var teklifler = models.TeklifListe()

	c.Data["json"] = teklifler
	c.ServeJSON()
}
