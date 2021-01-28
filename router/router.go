package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	middlewares "github.com/yy-c00/hotel-demo/middleware"
	"github.com/yy-c00/hotel-demo/model"
	"os"
)

//New returns a model.Router instance
func New() model.Router {
	return model.Router{
		RoomAccessRouter:  NewRoomAccessRouter(),
		UserAccessRouter: NewUserAccessRouter(),
		SalesRouter: NewSalesRouter(),
		StorageRouter: NewStorageRouter(),
		HistoryRouter: NewHistoryRouter(),
	}
}

//SetMiddlewares set a default middlewares
func SetMiddlewares(e *echo.Echo) {
	e.Use(middleware.Recover())

	//Default config (insecure)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
	}))

	logger, err := os.OpenFile("logs.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err == nil {
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: `{"host": "${host}", "remote_ip": "${remote_ip}", "method": "${method}", "uri": "${uri}", "code": ${status}, "error": "${error}", "bytes_in": ${bytes_in}, "bytes_out": ${bytes_out}, "Cookies":"${header:Cookie}"}` + "\n",
			Output: logger,
		}))
	}
}

//SetRoutes sets routes to *echo.Echo based on model.Router instance
func SetRoutes(e *echo.Echo, r model.Router) {

	e.POST("user/new", r.UserAccessRouter.CreateNewUser, middlewares.ToRoot())
	e.POST("user/validate", r.UserAccessRouter.ValidateUser)
	e.GET("user/available", r.UserAccessRouter.AvailableUsers, middlewares.ToRoot())
	e.GET("user/logged", r.UserAccessRouter.UserIsLogged, middlewares.ToUser())


	e.POST("room/new", r.RoomAccessRouter.CreateNewRoom, middlewares.ToRoot())
	e.POST("room/validate", r.RoomAccessRouter.ValidateRoom)
	e.GET("room/available", r.RoomAccessRouter.AvailableRooms)
	e.GET("room/logged", r.RoomAccessRouter.RoomIsLogged, middlewares.ToRoom())


	e.POST("sales/new", r.SalesRouter.NewSale, middlewares.ToAnyOne)
	e.PUT("sales/confirm", r.SalesRouter.ConfirmSale, middlewares.ToUser())

	e.DELETE("sales/cancel", r.SalesRouter.CancelSale, middlewares.ToUser())

	e.GET("sales/get", r.SalesRouter.GetSaleById, middlewares.ToUser())
	e.GET("sales/all", r.SalesRouter.GetAllPendingSales, middlewares.ToUser())

	e.GET("sales/channel", r.SalesRouter.Channel, middlewares.ToAnyOne)


	product := e.Group("storage/product")

	product.POST("", r.StorageRouter.AddProduct, middlewares.ToRoot())
	product.PUT("", r.StorageRouter.SetProduct, middlewares.ToRoot())
	product.DELETE("", r.StorageRouter.DeleteProduct, middlewares.ToRoot())
	product.GET("", r.StorageRouter.GetProductByID, middlewares.ToUser())

	product.GET("/search", r.StorageRouter.Search, middlewares.ToUser())

	category := e.Group("storage/category")

	category.GET("/all", r.StorageRouter.GetAll, middlewares.ToAnyOne)
	category.POST("/all", r.StorageRouter.AddCategories, middlewares.ToRoot())
	category.GET("", r.StorageRouter.GetCategoryByID, middlewares.ToAnyOne)
	category.POST("", r.StorageRouter.AddCategory, middlewares.ToRoot())
	category.DELETE("", r.StorageRouter.DeleteCategory, middlewares.ToRoot())

	e.GET("storage/channel", r.StorageRouter.Channel, middlewares.ToAnyOne)
	e.POST("storage/log", r.StorageRouter.Log, middlewares.ToRoot())


	history := e.Group("history", middlewares.ToRoot())
	history.GET("/all", r.HistoryRouter.GetLogsByRange)
	history.GET("", r.HistoryRouter.GetSaleLogById)
	history.GET("/search", r.HistoryRouter.SearchByRange)
}