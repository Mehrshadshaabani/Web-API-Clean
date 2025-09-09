package routers

import (
	handlers "github.com/Mehrshadshaabani/Web-API-Clean/api/Handlers"
	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {
	handler := &handlers.HealthHandler{}
	r.POST("/", handler.Health)
}
