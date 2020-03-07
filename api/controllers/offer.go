package controllers

import (
	"api/models"
)

// OfferController struct
type OfferController struct {
	BaseController
}

// GetAll function
// @Title GetAll
// @Description tum teklif kayitlarini getirir.
// @Success 200 {object} models.Offer
// @router /list [get]
func (c *OfferController) GetAll() {
	var offers = models.Offers()

	c.Data["json"] = offers
	c.ServeJSON()
}
