package main

import (
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/riadafridishibly/boring-todo/api"
	"github.com/riadafridishibly/boring-todo/frontend"
)

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func main() {
	todoApi, err := api.NewTodoAPI("test.db")
	if err != nil {
		log.Fatal(err)
	}
	routes := todoApi.Routes()

	distFS, err := fs.Sub(frontend.DistDir, "dist")
	if err != nil {
		panic(err)
	}
	FileServer(routes, "/", http.FS(distFS))

	http.ListenAndServe(":8989", routes)
}
