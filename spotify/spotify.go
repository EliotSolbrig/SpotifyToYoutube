package spotify

import (
    "fmt"
    "os"
    "net/http"

    sp "github.com/zmb3/spotify"
)

type SpotifyClient struct {
    Client *sp,Client.
}

func NewSpotifyClient(){
    redirectURL := os.Getenv("HOSTNAME") " + "/auth"
    auth = sp.NewAuthenticator(redirectURL, , spotify.ScopeUserReadPrivate)

    auth.SetAuthInfo(os.Getenv(("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET")))
}
