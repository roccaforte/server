package api

import (
	"net/http"
	"time"

	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

// New returns the main Roccaforte Server router.
func New() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CloseNotify)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Roccaforte Server"))
	})

	r.Mount("/login", loginRouter())
	r.Mount("/users", userRouter())
	r.Mount("/secrets", secretRouter())

	return r
}
