package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	memory "github.com/rogamba/hexagonal/adapters/memory"
	api "github.com/rogamba/hexagonal/api"
	config "github.com/rogamba/hexagonal/config"
	services "github.com/rogamba/hexagonal/domain/services"
)

// Define the adapters to be used (for example type of storage)

type App interface {
	Run() error
	RunAsync() chan error
}

// Define adapters to be used in the application
type app struct {
	config      config.AppConfig
	router      *chi.Mux
	userService services.UserService
}

// Inject the dependencies to the app
func NewApp() App {
	config, _ := config.ParseAppConfig()

	// Here we choose which persistence layer we want to use
	userMemoryStorage, _ := memory.NewMemoryUserRepository()
	tweetMemoryStorage, _ := memory.NewMemoryTweetRepository()

	// Services
	userService := services.NewUserService(
		userMemoryStorage,
		tweetMemoryStorage,
	)

	app := &app{
		config:      config,
		userService: userService,
	}

	// Third-party routing
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Bind router
	app.router = r
	setRoutes(app)

	return app
}

func (a *app) RunAsync() chan error {
	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port : " + a.config.Port)
		errs <- a.Run()
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
	return errs
}

func (a *app) Run() error {
	fmt.Println("Running server...")
	return http.ListenAndServe(":"+a.config.Port, a.router)
}

func setRoutes(a *app) error {
	handler := api.NewHandler(a.userService)
	a.router.Get("/{id}", handler.GetUser)
	return nil
}
