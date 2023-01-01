package api

import (
	"fmt"
	"go/models"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
}

func (con ApiController) GetData(c *gin.Context) {
	type Data struct {
		TableData  []models.TableData `json:"tableData"`
		CountData  []models.CountData `json:"countData"`
		TableLable models.TableLable  `json:"tableLable"`
	}

	var data = Data{
		TableData:  []models.TableData{},
		TableLable: models.TableLable{},
		CountData:  []models.CountData{},
	}
	models.DB.Find(&data.CountData)
	models.DB.Find(&data.TableData)
	models.DB.First(&data.TableLable)
	c.JSON(200, data)
}

//注册用户
func (con ApiController) Register(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	userList := models.User{Username: body.Username, Password: body.Password}
	d := models.DB.Create(&userList)
	if d.Error != nil {
		c.JSON(200, gin.H{
			"result": "用户已存在",
		})
	} else {
		c.JSON(200, gin.H{
			"result": "注册成功",
		})
	}
}

//登录
func (con ApiController) LoginUser(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	admin := models.User{}
	result := models.DB.Where("username = ?", body.Username).First(&admin)
	if result.Error != nil {
		c.JSON(200, gin.H{
			"result": "用户名不存在",
		})
	} else {
		if admin.Password == body.Password {

			if admin.PrivilegeLevel == 1 {
				type Data struct {
					Menu   []models.Menu `json:"menu"`
					Result string        `json:"result"`
				}
				var data = Data{
					Result: "登录成功",
				}

				models.DB.Find(&data.Menu)
				fmt.Println(data)
				c.JSON(200, data)
			} else {
				type Data struct {
					Menu   []models.Menu2 `json:"menu"`
					Result string         `json:"result"`
				}
				var data = Data{
					Result: "登录成功",
				}
				models.DB.Find(&data.Menu)
				c.JSON(200, data)
			}
		} else {
			c.JSON(200, gin.H{
				"result": "密码错误",
			})
		}
	}
}

//新增用户信息
func (con ApiController) AddUser(c *gin.Context) {
	var body struct {
		Name  string `json:"name"`
		Age   string `json:"age"`
		Sex   string `json:"sex"`
		Birth string `json:"birth"`
		Addr  string `json:"addr"`
	}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	var User = models.UserDetails{
		Name:  body.Name,
		Age:   body.Age,
		Sex:   body.Sex,
		Birth: body.Birth,
		Addr:  body.Addr,
	}
	d := models.DB.Create(&User)

	if d.Error != nil {
		c.JSON(200, gin.H{
			"result": "用户已存在",
		})

	} else {
		c.JSON(200, gin.H{
			"result": "添加成功",
		})
	}
}

//更新用户信息
func (con ApiController) EditUser(c *gin.Context) {
	var body struct {
		Name  string `json:"name"`
		Age   string `json:"age"`
		Sex   string `json:"sex"`
		Birth string `json:"birth"`
		Addr  string `json:"addr"`
	}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	var User = models.UserDetails{}
	d := models.DB.Model(&User).Where("name = ?", body.Name).Updates(body)
	if d.Error != nil {
		c.JSON(200, gin.H{
			"result": "更新失败",
		})
	} else {
		c.JSON(200, gin.H{
			"result": "更新成功",
		})
	}
}

//删除用户信息
func (con ApiController) DeleteUser(c *gin.Context) {
	var body struct {
		Name  string `json:"name"`
		Age   string `json:"age"`
		Sex   string `json:"sex"`
		Birth string `json:"birth"`
		Addr  string `json:"addr"`
	}
	if err := c.BindJSON(&body); err != nil {
		return
	}
	var User = models.UserDetails{}
	d := models.DB.Where("name=?", body.Name).Delete(&User)
	if d.Error != nil {
		c.JSON(200, gin.H{
			"result": "删除失败",
		})
	} else {
		c.JSON(200, gin.H{
			"result": "删除成功",
		})
	}
}
