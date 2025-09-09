package routers

import (
	handlers "github.com/Mehrshadshaabani/Web-API-Clean/api/Handlers"
	"github.com/gin-gonic/gin"
)

func Testrouter(r *gin.RouterGroup) {
	h := handlers.Newtesthandler()
	r.GET("/", h.HandleTestRequest)
	r.GET("/user/:id", h.GetuserbyId)
	r.POST("/user", h.Createuser)
	r.GET("/users", h.GetAllUsers)
	r.POST("/getheader", h.Headerbinder)
}
