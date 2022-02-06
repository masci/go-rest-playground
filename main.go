package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	s "github.com/go-rest-playground/storage"
)

var dbFile = flag.String("use-db", "", "Path to the SQLite database file")
var storage s.Storage

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano()) // this is needed by the storage package to create IDs

	// init the storage system according to user's preferences
	if *dbFile != "" {
		storage = s.NewSqliteStorage(*dbFile)
		fmt.Println("Using SQLite database", *dbFile)
	} else {
		storage = s.NewVolatileStorage()
		fmt.Println("Using in-memory storage, all data will be lost on exit")
	}
	defer storage.Close()

	// create routes
	r := chi.NewRouter()

	// healthcheck for containerized deployments
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// classes
	r.Route("/classes", func(r chi.Router) {
		r.Get("/", ListClasses)
		r.Post("/", CreateClass)
		r.Route("/{classID}", func(r chi.Router) {
			r.Use(ClassCtx)
			r.Get("/", GetClass)
			r.Put("/", UpdateClass)
			r.Delete("/", DeleteClass)
		})
	})

	// bookings
	r.Route("/bookings", func(r chi.Router) {
		r.Get("/", ListBookings)
		r.Post("/", CreateBooking)
		r.Route("/{bookingID}", func(r chi.Router) {
			r.Use(BookingCtx)
			r.Get("/", GetBooking)
			r.Put("/", UpdateBooking)
			r.Delete("/", DeleteBooking)
		})
	})

	// fire up the web server
	http.ListenAndServe(":3333", r)
}

// ClassCtx loads and injects a Class object into the request.
// In case the Class cannot be found, it returns a 404
func ClassCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		class, err := storage.GetClass(chi.URLParam(r, "classID"))
		if err != nil {
			render.Render(w, r, ErrNotFound(err))
			return
		}

		ctx := context.WithValue(r.Context(), "class", class)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// BookingCtx loads and injects a Class object into the request.
// In case the Class cannot be found, it returns a 404
func BookingCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "bookingID"))
		if err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}

		booking, err := storage.GetBooking(id)
		if err != nil {
			render.Render(w, r, ErrNotFound(err))
			return
		}

		ctx := context.WithValue(r.Context(), "booking", booking)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
