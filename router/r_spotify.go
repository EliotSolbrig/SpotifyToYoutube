package router

import (
    "fmt"
    "net/http"
    "encoding/json"
    "os"
    "strings"

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
}

func (router *Router) AuthSpotify(w http.ResponseWriter, r *http.Request){
    client,err := spotify.NewSpotifyClient()
    if err != nil {
        panic(err)
    }
    fmt.Println("client: ", client)

    // http.Redirect(w, r, "/", http.StatusSeeOther)

    // w.Write([]byte("success"))
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
    redirectURL += "/"
    auth := sp.NewAuthenticator(redirectURL, sp.ScopeUserReadPrivate)

    auth.SetAuthInfo(os.Getenv("SPOTIFY_CLIENT_ID"), os.Getenv("SPOTIFY_CLIENT_SECRET"))

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
      }

      response,err := json.Marshal(authResponse)
      if err != nil {
          panic(fmt.Errorf("Error marshaling auth response: %s"))
      }




    router.SpotifyClient.Client = &client


    w.Write([]byte(response))
}
