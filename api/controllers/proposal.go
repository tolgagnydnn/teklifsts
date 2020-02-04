package controllers

import (
	"api/models"

	"github.com/astaxie/beego"
)

// ProposalController struct
type ProposalController struct {
	beego.Controller
}

// GetAll function
// @Title GetAll
// @Description tum teklif kayitlarini getirir.
// @Success 200 {object} models.Proposal
// @router /list [get]
func (c *ProposalController) GetAll() {
	var teklifler = models.GetProposals()

	c.Data["json"] = teklifler
	c.ServeJSON()
}
