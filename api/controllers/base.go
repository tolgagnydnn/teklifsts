package controllers

import (
	"os"

	"github.com/astaxie/beego"
)

// SigningKey secret key for Token
var SigningKey = []byte(os.Getenv("TEKLIF_SECRET_KEY"))

// BaseController struct
type BaseController struct {
	beego.Controller
}
