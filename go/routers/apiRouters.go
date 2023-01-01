package routers

import (
	"go/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/")
	apiRouters.GET("/getData", api.ApiController{}.GetData)
	apiRouters.POST("/register", api.ApiController{}.Register)
	apiRouters.POST("/loginUser", api.ApiController{}.LoginUser)
	apiRouters.POST("/addUser", api.ApiController{}.AddUser)
	apiRouters.POST("/editUser", api.ApiController{}.EditUser)
	apiRouters.POST("/deleteUser", api.ApiController{}.DeleteUser)
}
