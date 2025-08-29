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
		fmt.Println(workDir)
		staticDir = filepath.Join(workDir, "internal/web/static")
		fmt.Println(staticDir)
	}
	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Adapt(func(w http.ResponseWriter, r *http.Request) (Node, error) {
		return SunsetPage(), nil
	}))

	http.ListenAndServe(":8080", nil)
}

func Switch(children ...Node) Node {
	return Input(Type("checkbox"), Role("switch"), Group(children))
}

const themeSwitchID = "theme-switch"

func SunsetPage() Node {
	// Load Eastern timezone
	easternTZ, err := time.LoadLocation("America/New_York")
	if err != nil {
		// Fallback to UTC if timezone loading fails
		easternTZ = time.UTC
	}

	return Page(HTML5Props{
		Title:       "Sunset",
		Description: "Watch the sunset together. Become and opacarophile.",
		Body: []Node{
			Main(
				Section(
					lucide.Sunset(
						Class("sunset-icon"),
					),
				),
				Section(
					H1(
						Text(fmt.Sprintf("If you are reading this (on %s), you are invited to watch the sunset at:", time.Now().In(easternTZ).Format("Monday, January 2, 2006"))),
					),
					H3(
						Text("Little Island, Manhattan"),
					),
				),
			),
			Footer(
				Label(
					Div(
						Class("theme-switch-container"),
						lucide.Sun(),
						Switch(ID(themeSwitchID)),
						lucide.Moon(),
					),
				),
				Script(Raw(`
					document.addEventListener("DOMContentLoaded", ()=>{
						// Get the theme switch element
						const themeSwitch = document.getElementById('`+themeSwitchID+`');
						
						// Check if user prefers dark mode and set switch state accordingly
						const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
						themeSwitch.checked = prefersDark;
						
						// Listen for theme switch changes
						themeSwitch.addEventListener('change', (e) => {
							if (e.target.checked) {
								document.documentElement.setAttribute('data-theme', 'dark');
							} else {
								document.documentElement.setAttribute('data-theme', 'light');
							}
						});
						
						// Listen for system theme changes
						window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
							if (e.matches) {
								themeSwitch.checked = true;
							} else {
								themeSwitch.checked = false;
							}
						});
					});
				`),
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
			Link(Rel("stylesheet"), Href("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.classless.min.css")),
			Link(Rel("stylesheet"), Href("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.colors.min.css")),
			// Link(Rel("stylesheet"), Href("https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.fluid.classless.min.css")),
			// Link(Rel("stylesheet"), Href("https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css")),
			Link(Rel("stylesheet"), Href("/static/pico.css")),
			Link(Rel("stylesheet"), Href("/static/index.css")),
			Script(
				Src("https://cdn.jsdelivr.net/gh/bigskysoftware/fixi@0.9.2/fixi.js"),
				CrossOrigin("anonymous"),
				Integrity("sha256-0957yKwrGW4niRASx0/UxJxBY/xBhYK63vDCnTF7hH4="),
				Async(),
			),
			Group(props.Head),
		},

		Body: []Node{
			Group(props.Body),
		},
	})
}
