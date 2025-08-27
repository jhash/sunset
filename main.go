package main

import (
	"net/http"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/http"

	b "github.com/willoma/bulma-gomponents"
)

func main() {
	http.HandleFunc("/", Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return SunsetPage(), nil
	}))

	http.ListenAndServe(":8080", nil)
}

func SunsetPage() Node {
	return Page(HTML5Props{
		Title: "Sunset",
		Description: "Sunset",
		Body: []Node{
			b.Hero(
				b.HeroHead(
					b.Container(
						b.Title(
							Text("If you are reading this, you are invited to watch the sunset at:"),
						),
					),
				),
				b.HeroFoot(
					b.Container(
						b.Title(Text("Pier 6, Brooklyn")),
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
			Group(props.Head),
		},

		Body: Group(props.Body),
	})
}