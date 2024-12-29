package main

import (
    "fmt"
    "net/http"
    "html/template"

	"github.com/gorilla/mux"

    "main/spotify"
    "main/router"
)

const staticDir string = "/static/"

func main(){
    fmt.Println("Spotify to youtube converter")

    r := mux.NewRouter()
    router := router.NewRouter()

    spotifyClient := spotify.NewSpotifyClient()

    staticHandler := http.StripPrefix(staticDir, http.FileServer(http.Dir("static/")))
    r.PathPrefix(staticDir).Handler(staticHandler)

    if spotifyClient == nil {
        panic(fmt.Errorf("Spotify client is null"))
    }

    r.HandleFunc("/"

}

func HomePage(w http.ResponseWriter, r *http.Request){
    templates := append([]string{"templates/pages/homepage.html"}, basePasefiles...)

    tmpl := template.Must(template.ParseFiles(templates...))

    err := tmpl.ExecuteTemplate(w, "base", map[string]any{
        "Data": map[string]any{
            "Test": "hi",
        },
    })
}
