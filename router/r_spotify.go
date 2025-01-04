package router

import (
    "fmt"
    "net/http"
    "encoding/json"
    "os"
    "strings"

    "golang.org/x/oauth2"
    sp "github.com/zmb3/spotify"

    "spot2yt/spotify"
)

type AuthRequest struct {
    RedirectURL string `json"redirect_url"`
    Authenticator sp.Authenticator `json"authenticator"`
}

type AuthResponse struct {
    Client *sp.Client
    Success bool
    Token *oauth2.Token
}

func (router *Router) AuthSpotify(w http.ResponseWriter, r *http.Request){
    client,err := spotify.NewSpotifyClient()
    if err != nil {
        panic(err)
    }
    fmt.Println("client: ", client)

    // http.Redirect(w, r, "/", http.StatusSeeOther)

    w.Write([]byte("success"))

    http.Redirect(w, r, "/auth/spotify/get", http.StatusSeeOther)
}


func (router *Router) GetSpotifyClient(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.Method, " request on /auth")
    
    // clientInfo,err := spotify.NewSpotifyClient(w,r)
    // if err != nil {
    //     panic(fmt.Errorf("Error getting client info: %s", err))
    // }

    redirectURL := os.Getenv("HOSTNAME")
    if strings.Contains(redirectURL, "localhost") || strings.Contains(redirectURL, "127.0.0.1") {
        redirectURL += ":" + os.Getenv("PORT")
    }
    redirectURL += "/auth/spotify/get"
    // redirectURL += "/"
    auth := sp.NewAuthenticator(redirectURL, sp.ScopeUserReadPrivate)

    auth.SetAuthInfo(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET"))

    spClient := router.SpotifyClient
    fmt.Println("spClient: ", spClient)

    if spClient.Token != nil {
        newClient := auth.NewClient(spClient.Token)
        authResponse := AuthResponse{
            Client: &newClient,
            Token: spClient.Token,
        }

        fmt.Println("authResponse: ", authResponse.Token)

        response,err := json.Marshal(authResponse)
        if err != nil {
            panic(fmt.Errorf("Error marshaling auth response: %s", err))
        }
        fmt.Println("response: ", response)
        w.Write(response)
    }

    url := auth.AuthURL("weow")
    fmt.Println("url: ", url)

    http.Redirect(w, r, url, http.StatusSeeOther) 

    token, err := auth.Token("weow", r)
      if err != nil {
          tempError := fmt.Errorf("Error getting token: %s", err)
          fmt.Println("tempError: ", tempError)
            // http.Error(w, "Couldn't get token", http.StatusNotFound)
      }
      fmt.Println("token: ", token)
      // create a client using the specified token
      client := auth.NewClient(token)
      fmt.Println("client: ", client)

      authResponse := AuthResponse{
          Client: &client,
          Token: token,
      }

      fmt.Println("authResponse: ", authResponse.Token)

      response,err := json.Marshal(authResponse)
      if err != nil {
          panic(fmt.Errorf("Error marshaling auth response: %s", err))
      }
      fmt.Println("response: ", response)




    router.SpotifyClient.Client = &client
    router.SpotifyClient.Token = token


    // http.Redirect(w, r, "/", http.StatusSeeOther)
    w.Write(response)
}


func (router *Router) GetSpotifySongInfo(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.Method, " request on /getspotifysonginfo")

    spotifyAuthStatus := router.GetSpotifyAuthStatus()
    if spotifyAuthStatus == false {
        panic(fmt.Errorf("Spotify not authenticated"))
    }

    r.ParseForm()
    songURL := r.FormValue("song-info-div")
    fmt.Println("songURL: ", songURL)

    songInfo,err := spotify.GetSongInfoFromURL(songURL, router.SpotifyClient)
    if err != nil {
        panic(fmt.Errorf("Error getting song info from url: %s", err))
    }
    fmt.Println("songInfo: ", songInfo)
    searchName := songInfo.Name + " - " + songInfo.ArtistNames[0]
    fmt.Println("searchName: ", searchName)
}
