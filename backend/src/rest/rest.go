package rest

import "github.com/gin-gonic/gin"

// RunAPI is a function that starts the rest api's
func RunAPI(address string) error {

	h, err := NewHandler("mysql", "root:0911@tcp(127.0.0.1:3306)/gomusic?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	return RunAPIWithHandler(address, h)
}

// RunAPIWithHandler is a function that starts the api engine
func RunAPIWithHandler(addres string, h HandlerInterface) error {

	r := gin.Default()

	r.GET("/products", h.GetProducts)
	r.GET("/promos", h.GetPromos)

	userGroup := r.Group("/users")
	{
		userGroup.POST("/signin", h.SignIn)
		userGroup.POST("", h.AddUser)
		userGroup.POST("/charge", h.Charge)
	}

	userGroup = r.Group("/user")
	{
		userGroup.POST("/:id/signout", h.SignOut)
		userGroup.GET("/:id/orders", h.GetOrders)
	}

	return r.Run(addres)
	// return r.RunTLS(addres, "cert.pem", "key.pem")

}
