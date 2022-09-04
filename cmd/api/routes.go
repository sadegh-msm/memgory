package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"simple-redis/cmd/databeses"
)

func NewRouter(e *echo.Echo) *echo.Echo {
	db := databeses.NewStorage()
	cmd := NewCmd(db)

	// use logger middleware to log events on server
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "time=${time_rfc3339}, " +
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
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(10)))

	e.POST("/set", cmd.Set)
	e.GET("/get", cmd.Get) // get value from key by query parameters {"name" for database name} {"key" for data key}
	e.DELETE("/del", cmd.Del)
	e.POST("/use", cmd.UseDB)
	e.POST("/keyregex", cmd.KeyRegex)
	e.GET("/listdata", cmd.ListData) // get list of values from database name by query parameters {"name" for database name}
	e.GET("/listdb", cmd.ListDBs)    // get list of databases from storage name by query parameters {"name" for storage name}
	e.POST("/save", cmd.Save)
	e.POST("/load", cmd.Load)

	return e
}
