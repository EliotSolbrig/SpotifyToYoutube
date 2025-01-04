package router

import (
    "fmt"
    "context"
    "os"

    "spot2yt/spotify"

    "google.golang.org/api/option"
    youtube "google.golang.org/api/youtube/v3"
)

var basePasefiles []string = []string{
    "templates/base.html",
    "templates/components/header.html",
    "templates/components/footer.html",
}

type Router struct {
    SpotifyClient *spotify.SpotifyClient
    // YoutubeClient 
    YoutubeService *youtube.Service
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

    ytService,err := youtube.NewService(context.Background(), option.WithAPIKey(os.Getenv("YOUTUBE_API_KEY")))

    if err != nil {
        panic(fmt.Errorf("Error getting new youtube service: %s", err))
    }
    fmt.Println("ytService: ", ytService)

    return &Router{
        SpotifyClient: &spotify.SpotifyClient{
            Client: nil,
        },
        YoutubeService: ytService,
    }
}

func (router *Router) GetSpotifyAuthStatus() bool {
    var spotifyAuthStatus bool
    // var youtubeAuthStatus bool
    spotifyClient := router.SpotifyClient.Client
    fmt.Println("spotifyClient: ", spotifyClient)

    if spotifyClient == nil {
        spotifyAuthStatus = false
    } else {
        spotifyAuthStatus = true
    }

    return spotifyAuthStatus
}
