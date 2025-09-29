package api

import (
	"fmt"

	middleware "github.com/Mehrshadshaabani/Web-API-Clean/api/midlewares"
	"github.com/Mehrshadshaabani/Web-API-Clean/api/routers"
	validation "github.com/Mehrshadshaabani/Web-API-Clean/api/validations"
	"github.com/Mehrshadshaabani/Web-API-Clean/config"
	"github.com/Mehrshadshaabani/Web-API-Clean/docs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(cfg *config.Config) {
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		val.RegisterValidation("mobile", validation.IranianMobileNumberValidator, true)
		val.RegisterValidation("password", validation.PasswordValidator, true)
	}
	r.Use(middleware.Cors(cfg))
	r.Use(gin.Logger(), gin.Recovery(), middleware.LimitByRequest())
	ConfigSwagger(r, *cfg)
	api := r.Group("/api")
	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")

		routers.Health(health)
		routers.Testrouter(test_router)
	}

	v2 := api.Group("/v2")
	{
		health := v2.Group("/health")
		routers.Health(health)
	}

	r.Run(":5005")
}

func ConfigSwagger(r *gin.Engine, cfg config.Config) {
	docs.SwaggerInfo.Title = "Golang Clean Api"
	docs.SwaggerInfo.Description = "This is a sample server for a Golang Clean API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", cfg.Server.InternalPort)
	docs.SwaggerInfo.BasePath = "/api"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
