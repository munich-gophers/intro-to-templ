package main

import (
	// "context"
	"context"
	"fmt"
	"net/http"

	// "os"

	page "munich-gophers-intro-to-templ/internal/templates/page"

	"github.com/a-h/templ"
)

// func main() {
// 	component := page.Index()
// 	component.Render(context.Background(), os.Stdout)
// }

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	component := templ.Handler(page.About())
	withMiddleware := Middleware(component)

	// Static files
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// Routes
	http.Handle("/", templ.Handler(page.Index()))
	http.Handle("/about", withMiddleware)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
