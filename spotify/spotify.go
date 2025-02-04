package spotify

import (
    "fmt"
    "os"
    "net/http"
    "io"
    "strings"
    "encoding/json"

    "golang.org/x/oauth2"
    sp "github.com/zmb3/spotify"

)

type SpotifyClient struct {
    Client *sp.Client
    RedirectURL string
    Token *oauth2.Token
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
    // redirectURL += "/auth/spotify/get"
    redirectURL += "/"

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

    fmt.Println("authResponse: ", authResponse)

    // if err != nil {
    //     return NewClientResponse{}, fmt.Errorf("Error marshaling response while getting new spotify client: %s", err)
    // }


   return authResponse.Client, nil 

}

type SongInfo struct {
    Name string
    ArtistNames []string
    AlbumName string
}

func GetSongInfoFromURL(songURI string, spotifyClient *SpotifyClient) (SongInfo, error){
    // spotifyAuthStatus := router.NewRouter().GetSpotifyAuthStatus() 
    withoutParams := strings.Split(songURI,"?")[0]
    fmt.Println("withoutParams: ", withoutParams)

    songInfo := SongInfo{}

    trackID := strings.Split(withoutParams, "track/")[1]
    fmt.Println("trackID: ", trackID)
    trackInfo, err := spotifyClient.Client.GetTrack(sp.ID(trackID))
    if err != nil {
        return SongInfo{}, err
    }
    songInfo.Name = trackInfo.Name
    fmt.Println("trackInfo: ", trackInfo)

    fmt.Println("artists: ", trackInfo.Artists[0])

    artistNames := []string{}
    for _,artist := range trackInfo.Artists {
        fmt.Println("artist: ", artist)
        artistNames = append(artistNames, artist.Name)
    }

    songInfo.ArtistNames = artistNames

    songInfo.AlbumName = trackInfo.Album.Name

    fmt.Println("songInfo: ", songInfo)

    return songInfo, nil
}
