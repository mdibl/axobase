package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.RenderTemplate("app/index.html")
}

func (c App) Databases() revel.Result {
	return c.RenderTemplate("app/databases.html")
}

func (c App) Nomenclature() revel.Result {
	return c.RenderTemplate("app/nomenclature.html")
}

func (c App) Lines() revel.Result {
	return c.RenderTemplate("app/lines.html")
}
