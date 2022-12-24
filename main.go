package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/getlantern/systray"
	"github.com/go-chi/chi/v5"
	"github.com/riadafridishibly/boring-todo/api"
	"github.com/riadafridishibly/boring-todo/frontend"
	"github.com/skratchdot/open-golang/open"
)

const SITE_URL = "http://localhost:8989/"

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

func handleOpen(m *systray.MenuItem) {
	go func() {
		for range m.ClickedCh {
			open.Run(SITE_URL)
		}
	}()
}

func handleQuit(m *systray.MenuItem) {
	go func() {
		for range m.ClickedCh {
			systray.Quit()
		}
	}()
}

var server *http.Server

func newServer() *http.Server {
	srv := &http.Server{}
	srv.Addr = ":8989"

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

	srv.Handler = routes

	return srv
}

func onReady() {
	systray.SetTitle("TODO")
	systray.SetTooltip("Boring todos")
	o := systray.AddMenuItem("Open", "Open boring todo application")
	handleOpen(o)

	q := systray.AddMenuItem("Quit", "Quit the boring application")
	handleQuit(q)

	server = newServer()
	go func() {
		err := server.ListenAndServe()
		if err == http.ErrServerClosed {
			fmt.Println("Server shutdown successful")
		} else {
			log.Println(err)
		}
	}()
}

func onExit() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}

func main() {
	systray.Run(onReady, onExit)
}
