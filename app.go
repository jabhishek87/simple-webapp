package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func getBkColor() string {

	// Set white as default color if env var is not set
	appColor := "white"
	if ac := os.Getenv("APP_COLOR"); ac != "" {
		appColor = ac
	}

	return appColor
}

// template data
var data = make(map[string]interface{})

func main() {

	myfile, _ := os.Create("app.log")
	gin.DefaultWriter = io.MultiWriter(myfile, os.Stdout)
	router := gin.Default()

	// data["name"] = ctx.DefaultQuery("name", "Guest !")
	data["name"], _ = os.Hostname()
	// get app color
	data["color"] = getBkColor()

	//new template engine
	router.HTMLRender = gintemplate.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{"data": data})
	})

	router.GET("/color/:color", func(ctx *gin.Context) {
		color := ctx.Param("color")
		data["color"] = color
		ctx.HTML(http.StatusOK, "index.html", gin.H{"data": data})
	})

	router.GET("/readlogs/", func(ctx *gin.Context) {

		content, _ := ioutil.ReadFile("app.log")
		data["log"] = template.HTML(strings.Replace(string(content), "\n", "<br>", -1))

		ctx.HTML(http.StatusOK, "index.html", gin.H{"data": data})
	})

	router.Run("0.0.0.0:9090")
}
