package main

import (
	"fmt"
	// "os"
	// "github.com/joho/godotenv"  // install go get github.com/joho/godotenv
	"io"
	"log"
	"net/http"
	"regexp"
)

type Handler func(*Context)

type Route struct {
	Pattern *regexp.Regexp
	Handler Handler
}

type App struct {
	Routes       []Route
	DefaultRoute Handler
}

func NewApp() *App {
	app := &App{
		DefaultRoute: func(ctx *Context) {
			ctx.Text(http.StatusNotFound, "Not found")
		},
	}

	return app
}

func (a *App) Handle(pattern string, handler Handler) {
	re := regexp.MustCompile(pattern)
	route := Route{Pattern: re, Handler: handler}

	a.Routes = append(a.Routes, route)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := &Context{Request: r, ResponseWriter: w}

	for _, rt := range a.Routes {
		if matches := rt.Pattern.FindStringSubmatch(ctx.URL.Path); len(matches) > 0 {
			if len(matches) > 1 {
				ctx.Params = matches[1:]
			}

			rt.Handler(ctx)
			return
		}
	}

	a.DefaultRoute(ctx)
}

type Context struct {
	http.ResponseWriter
	*http.Request
	Params []string
}

func (c *Context) Text(code int, body string) {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain")
	c.WriteHeader(code)

	io.WriteString(c.ResponseWriter, fmt.Sprintf("%s\n", body))
}

func main() {
	// reference => https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438
	app := NewApp()

// 	err := godotenv.Load(".env")

//   if err != nil {
//     log.Fatalf("Error loading .env file")
//   }
//   httpPort := ""
//   httpPort := os.Getenv("SERVER_PORT")
// 	if httpPort != nil {
		httpPort := ":9000"
// 	}

	app.Handle(`.`, func(ctx *Context) {
const httpPort = 3000;
http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist")))
		
	})

	log.Printf("Running on port %s\n", httpPort)
	err := http.ListenAndServe(httpPort, app)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}