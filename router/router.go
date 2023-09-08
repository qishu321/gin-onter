package router

import (
	"gin-onter/api/lottery"
	"gin-onter/conf"
	"gin-onter/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(conf.AppMode)

	r := gin.Default()
	r.Use(middleware.Cors())

	s := &http.Server{
		Addr:           conf.HttpPort,
		Handler:        r,
		ReadTimeout:    conf.ReadTimeout,
		WriteTimeout:   conf.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	//r.Use(middleware.Jaeger())
	r.Use(middleware.Logger())
	r.StaticFS("/static",http.Dir(conf.StaticFS))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})
	api_lottery := r.Group("api/lottery")
	{
		api_lottery.POST("Addlottery",lottery.Addlottery)
		api_lottery.POST("Editlottery",lottery.Editlottery)
		api_lottery.POST("Getlottery",lottery.Getlottery)
		api_lottery.POST("Dellottery",lottery.Dellottery)
		api_lottery.POST("Lotterymain",lottery.Lotterymain)

		api_lottery.POST("Addprize",lottery.Addprize)
		api_lottery.POST("EditPrize",lottery.EditPrize)
		api_lottery.POST("GetPrize",lottery.GetPrize)
		api_lottery.POST("DelPrize",lottery.DelPrize)



	}
	s.ListenAndServe()

}