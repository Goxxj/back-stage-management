package routers

import (
	"go/controllers/def"

	"github.com/gin-gonic/gin"
)

func DefRoutersInit(r *gin.Engine) {
	defRouters := r.Group("/")
	defRouters.GET("/", def.DefController{}.Index)
	defRouters.GET("/getUser", def.DefController{}.GetUser)
}
