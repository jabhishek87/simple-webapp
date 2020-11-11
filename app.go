package main

import (
	"io"
	"math/rand"
	"net/http"
	"os"
	"reflect"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

var colorMap = map[string]string{
	"red":      "#e74c3c",
	"green":    "#16a085",
	"blue":     "#2980b9",
	"blue2":    "#30336b",
	"pink":     "#be2edd",
	"darkblue": "#130f40",
}

func getBkColor() string {
	// get keys from colormap
	keys := reflect.ValueOf(colorMap).MapKeys()

	// get random color from colormap
	defaultAppColor := keys[rand.Intn(len(keys))]

	appColor := os.Getenv("APP_COLOR")

	// if env var no set fall back o default
	if val, ok := colorMap[appColor]; ok {
		appColor = val
	} else {
		appColor = defaultAppColor.String()
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
		data["colors"] = reflect.ValueOf(colorMap).MapKeys()
		ctx.HTML(http.StatusOK, "index.html", gin.H{"data": data})
	})

	router.Run(":9090")
}
