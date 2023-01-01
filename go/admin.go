package main

import (
	"go/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../dist/index.html")
	r.StaticFS("/dist", http.Dir("../dist"))
	routers.DefRoutersInit(r)
	routers.ApiRoutersInit(r)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
