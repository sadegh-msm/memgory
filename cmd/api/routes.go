package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo) *echo.Echo {
	// use logger middleware to log events on server
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_unix}, " +
			"uri=${uri}, " +
			"status=${status}, " +
			"error=${error}, " +
			"host={host}, " +
			"id=${id}, " +
			"latency_human=${latency_human}\n",
	}))
	// use cors middleware to stop not allowed requests
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	// use body limit middleware to stop requests that want to store much data in database
	e.Use(middleware.BodyLimit("10K"))
	// use rate limiter to stop high loads
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(50)))

	e.POST("/set", Set)
	e.GET("/get", Get)
	e.DELETE("/del", Del)
	e.GET("/keys", Keys)
	e.GET("/list", List)
	e.GET("/save", Save)
	e.POST("/load", Load)

	return e
}
