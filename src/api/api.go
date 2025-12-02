package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/salmantaghooni/golang-car-web-api/api/middlewares"
	"github.com/salmantaghooni/golang-car-web-api/api/routers"
	"github.com/salmantaghooni/golang-car-web-api/api/validations"
	"github.com/salmantaghooni/golang-car-web-api/config"
	"github.com/salmantaghooni/golang-car-web-api/docs"
	"github.com/salmantaghooni/golang-car-web-api/pkg/logging"
	"github.com/salmantaghooni/golang-car-web-api/pkg/metrics"
)

var logger = logging.NewLogger(config.GetConfig())

func InitServer(cfg *config.Config) {
	r := gin.New()
	RegisterPrometheus()
	RegisterMiddlewares(r, cfg)
	RegisterRoutes(r, cfg)
	RegisterSwagger(r, cfg)
	RegisterValidator()
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}

func RegisterMiddlewares(r *gin.Engine, cfg *config.Config) {
	r.Use(middlewares.Prometheus())
	r.Use(middlewares.Cors(cfg))
	r.Use(middlewares.DefaultStracturedLogger(cfg))
	r.Use(middlewares.LimitByRequest())
	r.Use(gin.Logger(), gin.CustomRecovery(middlewares.ErrorHandler))
}

func RegisterValidator() {

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("mobile", validations.IranianMobileNumberValidate)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {

	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		countries := v1.Group("/countries", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		cities := v1.Group("/cities", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		// Property
		properties := v1.Group("/properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		propertyCategories := v1.Group("/property-categories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		//company
		companies := v1.Group("/companies", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		//base
		files := v1.Group("/files", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		colors := v1.Group("/colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		years := v1.Group("/years", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		//user
		users := v1.Group("/users")

		// Car
		carTypes := v1.Group("/car-types", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		gearboxes := v1.Group("/gearboxes", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelCommetns := v1.Group("/car-model-comments", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelProperties := v1.Group("/car-model-properties", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModels := v1.Group("/car-models", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelColors := v1.Group("/car-model-colors", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelYears := v1.Group("/car-model-years", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelPriceHistories := v1.Group("/car-model-price-histories", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))
		carModelImages := v1.Group("/car-model-images", middlewares.Authentication(cfg), middlewares.Authorization([]string{"admin"}))

		routers.CarType(carModels, cfg)
		routers.CarType(carTypes, cfg)
		routers.Gearbox(gearboxes, cfg)
		routers.CarModelProperty(carModelProperties, cfg)
		routers.CarModelComment(carModelCommetns, cfg)
		routers.CarModelColor(carModelColors, cfg)
		routers.CarModelYear(carModelYears, cfg)
		routers.CarModelPriceHistory(carModelPriceHistories, cfg)
		routers.CarModelImage(carModelImages, cfg)

		routers.User(users, cfg)
		routers.Country(countries, cfg)
		routers.City(cities, cfg)
		routers.File(files, cfg)
		routers.Company(companies, cfg)
		routers.Color(colors, cfg)
		routers.Year(years, cfg)

		// Property
		routers.Property(properties, cfg)
		routers.PropertyCategory(propertyCategories, cfg)
		r.Static("/static", "./uploads")
	}

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// v2 := api.Group("/v2")
	// {

	// }
}
func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.Port)
	docs.SwaggerInfo.Schemes = []string{"http"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

}

func RegisterPrometheus() {
	err := prometheus.Register(metrics.DbCall)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}

	err = prometheus.Register(metrics.HttpDuration)
	if err != nil {
		logger.Error(logging.Prometheus, logging.Startup, err.Error(), nil)
	}
}
