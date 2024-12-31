package router

import (
    // "fmt"

    "spot2yt/spotify"
)

var basePasefiles []string = []string{
    "templates/base.html",
    "templates/components/header.html",
    "templates/components/footer.html",
}

type Router struct {
    SpotifyClient *spotify.SpotifyClient
    // YoutubeClient 
}

// func NewRouter(service service.IService) *Router {
func NewRouter() *Router {

    // spClient,err := spotify.NewSpotifyClient()
    // if err != nil {
    //     tempError := fmt.Errorf("Error getting new spotify client: %s", err)
    //     fmt.Println("tempError: ", tempError)
    //
    // }
    // spClient := spotify.SpotifyClient{}

    return &Router{
        SpotifyClient: &spotify.SpotifyClient{
            Client: nil,
        },
    }
}
