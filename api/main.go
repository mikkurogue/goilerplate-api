package api

import (
	"context"
	"goilerplate-api/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Start() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("no .env file exists...")
	}

	e.Use(middleware.Logger())

	// Default cors cfg
	e.Use(middleware.CORS())

	// Recover from panics if applicable
	e.Use(middleware.Recover())

	Routes(e)

	// Try to gracefully shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := e.Start(":1337"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func Routes(e *echo.Echo) {

	e.GET("/hello-world", handlers.HelloWorld)
	e.GET("/hello-world-json", handlers.HelloWorldJson)

	e.POST("/hello-world", handlers.PostHelloWorld)
	e.POST("/hello-world-multipart", handlers.PostHelloWorldMultipart)

	e.POST("/register", handlers.Register)
	e.POST("/login", handlers.Login)

	protected := e.Group("/protected")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(handlers.JWTClaims)
		},
		// Make sure this is the same env var
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	protected.Use(echojwt.WithConfig(config))

	protected.GET("/hello-protected-world", handlers.HelloProtectedWorld)
}
