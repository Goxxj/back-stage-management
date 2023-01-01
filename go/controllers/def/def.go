package def

import (
	"go/models"

	"github.com/gin-gonic/gin"
)

type DefController struct {
}

func (con DefController) Index(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{})
}
func (con DefController) GetUser(c *gin.Context) {
	var UserDetails []models.UserDetails
	if c.Query("name") == "all" {
		models.DB.Find(&UserDetails)
		c.JSON(200, UserDetails)
	} else {
		models.DB.Where("name=?", c.Query("name")).First(&UserDetails)
		c.JSON(200, UserDetails)
	}

}
