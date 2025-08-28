package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
	. "maragu.dev/gomponents/http"

	lucide "github.com/eduardolat/gomponents-lucide"
)

func main() {
	godotenv.Load()

	staticDir := os.Getenv("STATIC_PATH")
	if staticDir == "" {
		workDir, _ := os.Getwd()
		staticDir = filepath.Join(workDir, "web/static")
	}
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return SunsetPage(), nil
	}))

	http.ListenAndServe(":8080", nil)
}

func SunsetPage() Node {
	return Page(HTML5Props{
		Title:       "Sunset",
		Description: "Sunset",
		Body: []Node{
			Div(
				Class("container is-flex is-flex-direction-column is-align-items-center is-flex-grow-1"),
				Div(
					Class("block pt-6"),
					lucide.Sunset(
						Style("width: 150px; height: 150px; color:rgb(33, 33, 33);"),
					),
				),
				Div(
					Class("hero is-medium is-warning is-outlined m-5"),
					Div(
						Class("hero-body"),
						H1(
							Class("hero-title is-size-1 has-text-weight-bold block"),
							Text(fmt.Sprintf("If you are reading this (on %s), you are invited to watch the sunset at:", time.Now().Format("Monday, January 2, 2006"))),
						),
						H3(
							Class("hero-subtitle is-size-2"),
							Text("Astoria Park War Memorial, Queens"),
						),
					),
				),
			),
		},
	})
}

func Page(props HTML5Props) Node {
	return HTML5(HTML5Props{
		Title:       props.Title,
		Description: props.Description,

		Head: []Node{
			Link(Rel("stylesheet"), Href("https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css")),
			Link(Rel("stylesheet"), Href("/static/index.css")),
			Group(props.Head),
		},

		Body: []Node{
			Group(props.Body),
		},
	})
}
