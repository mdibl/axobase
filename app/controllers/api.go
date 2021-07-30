package controllers

// import (
// 	"biocore_resource_server/app/utils"
// 	"fmt"
// 	"os"

// 	"github.com/revel/revel"
// )

// type APIV1 struct {
// 	App
// }

// func (c APIV1) RenderImg(category, topic, dir string) revel.Result {
// 	img := c.Params.Query.Get("image")
// 	pathString := "%s/%s/%s/%s/%s"
// 	imgPath := fmt.Sprintf(
// 		pathString,
// 		utils.ReadTutorialRepoPath(),
// 		category,
// 		topic,
// 		dir,
// 		img,
// 	)
// 	f, err := os.Open(imgPath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return c.RenderFile(f, revel.Inline)
// }
