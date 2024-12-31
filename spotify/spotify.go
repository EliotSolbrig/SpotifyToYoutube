package spotify

import (
    "fmt"
    "os"
    "net/http"
    "io"
    "strings"
    "encoding/json"

    sp "github.com/zmb3/spotify"
)

type SpotifyClient struct {
    Client *sp.Client
    RedirectURL string
}

type NewClientResponse struct {
    RedirectURL string `json"redirect_url"`
    Authenticator sp.Authenticator `json"authenticator"`
}

type AuthRequest struct {
    Client *sp.Client
    Success bool
}

func NewSpotifyClient() (*sp.Client, error) {
    fmt.Println("Getting/authorizing spotify client")

    // httpClient := &http.Client{}

    redirectURL := os.Getenv("HOSTNAME")
    if strings.Contains(redirectURL, "localhost") || strings.Contains(redirectURL, "127.0.0.1") {
        redirectURL += ":" + os.Getenv("PORT")
    }
    redirectURL += "/auth/spotify/get"

    res,err := http.Get(redirectURL)

    if err != nil {
        tempError := fmt.Errorf("Error making get request on /auth: %s", err)
        fmt.Println("tempError: ", tempError)
    }

    var resBody []byte

    // if res != nil {
        resBody,err = io.ReadAll(res.Body)
        if err != nil {
            return nil, fmt.Errorf("Error reading response body: %s", err) 
        }
    // }
    fmt.Println("resBody: ", string(resBody))

    authResponse := AuthRequest{}
    err = json.Unmarshal(resBody, &authResponse)

    // if err != nil {
    //     return NewClientResponse{}, fmt.Errorf("Error marshaling response while getting new spotify client: %s", err)
    // }


   return authResponse.Client, nil 

}
