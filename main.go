package main

import (
	"context"
	"flag"
	"fmt"
	"hcn/api"
	"hcn/config"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/csrf"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/itrepablik/itrlog"
	"github.com/itrepablik/sakto"
)

//CurrentLocalTime bla bla...
var CurrentLocalTime = sakto.GetCurDT(time.Now(), "America/New_York")

func main() {
	os.Setenv("TZ", config.SiteTimeZone) // Set the local timezone globally
	fmt.Println("Starting the web servers at ", CurrentLocalTime)
	var dir string
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&dir, "dir", "static", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()

	router := mux.NewRouter()
	csrfMiddleware := csrf.Protect(
		[]byte(config.SecretKeyCORS),
		csrf.TrustedOrigins([]string{config.SiteDomainName}),
	)

	// This is related to the CORS config to allow all origins []string{"*"} or specify only allowed IP or hostname.
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT"}),
		handlers.AllowedOrigins([]string{config.SiteDomainName}),
	)

	router.Use(cors)
	router.Use(csrfMiddleware)
	router.Use(loggingMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))

	// This will serve the files under http://localhost:3000/static/<filename>
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	// Initialize the APIs here
	api.MainRouters(router) // URLs for the main app.

	srv := &http.Server{
		Addr: "localhost:4000",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		msg := `Web server started at `
		fmt.Println(msg, CurrentLocalTime)
		itrlog.Info("Web server started at ", CurrentLocalTime)
		if err := srv.ListenAndServe(); err != nil {
			itrlog.Error("Hola ", err)
		}
	}() // Note the parentheses - must call the function.

	// BUFFERED CHANNELS = QUEUES
	c := make(chan os.Signal, 1) // Queue with a capacity of 1.

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	fmt.Println("Shutdown web server at " + CurrentLocalTime.String())
	itrlog.Warn("Server has been shutdown at ", CurrentLocalTime.String())
	os.Exit(0)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, router *http.Request) {
		// Do stuff here
		req := "IP:" + sakto.GetIP(router) + ":" + router.RequestURI + ":" + CurrentLocalTime.String()
		fmt.Println(req)
		itrlog.Info(req)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, router)
	})
}
