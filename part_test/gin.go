package main 
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ping(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"msg":"ping pong"})
}

func user(c *gin.Context){
	name := c.Param("name")
	msg := "request name is "+name
	c.String(http.StatusOK,msg)
}

func form(c *gin.Context){
	name := c.PostForm("name")
	age := c.DefaultPostForm("age","25")
	c.JSON(http.StatusOK,gin.H{"name":name,"age":age})
}

func welcome(c *gin.Context){
	firstName := c.Query("first")
	lastName := c.DefaultQuery("last","guest")
	msg := "Hello, "+lastName +" " + firstName
	c.String(http.StatusOK,msg)
}

func main(){
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/user/:name", user)
	r.GET("/welcome", welcome)
	r.POST("/form", form)
	r.Run() // listen and serve on 0.0.0.0:8080
}