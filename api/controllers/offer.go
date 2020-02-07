package controllers

import (
	"api/models"

	"github.com/astaxie/beego"
)

// OfferController struct
type OfferController struct {
	beego.Controller
}

// GetAll function
// @Title GetAll
// @Description tum teklif kayitlarini getirir.
// @Success 200 {object} models.Proposal
// @router /list [get]
func (c *OfferController) GetAll() {
	var offers = models.Offers()

	c.Data["json"] = offers
	c.ServeJSON()
}
