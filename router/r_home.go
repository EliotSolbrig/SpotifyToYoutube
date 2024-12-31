package router

import (
    "fmt"
    "net/http"
    "html/template"
)

func (router *Router) HomePage(w http.ResponseWriter, r *http.Request){

    var spotifyAuthStatus bool
    // var youtubeAuthStatus bool
    spotifyClient := router.SpotifyClient.Client
    fmt.Println("spotifyClient: ", spotifyClient)

    if spotifyClient == nil {
        spotifyAuthStatus = false
    } else {
        spotifyAuthStatus = true
    }

    templates := append([]string{"templates/pages/homepage.html",}, basePasefiles...)

    tmpl := template.Must(template.ParseFiles(templates...))

    err := tmpl.ExecuteTemplate(w, "base", map[string]any{
        "Data": map[string]any{
            "SpotifyAuthStatus": spotifyAuthStatus,
        },
    })

    if err != nil {
        panic(fmt.Errorf("Error executing home page template: %s", err))
    }
}
