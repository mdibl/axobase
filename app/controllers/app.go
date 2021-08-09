package controllers

import (
	"axobase/app/utils"

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

func (c App) Antibodies() revel.Result {
	records, err := utils.ParseAntibodiesCSV("antibodies.csv")
	if err != nil {
		//TODO: Log and handle the error accordingly.
		panic(err)
	}
	c.ViewArgs["records"] = records
	return c.RenderTemplate("app/antibodies.html")
}

func (c App) ResearchLabs() revel.Result {
	return c.RenderTemplate("app/research-labs.html")
}
