package router

import (
	// "context"
	"fmt"
	"net/http"

	"spot2yt/spotify"
	// "spot2yt/yt"

	// youtube "google.golang.org/api/youtube/v3"
)

func (router *Router) ConvertSong(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.Method, " request on /convert")

    r.ParseForm()
    songURL := r.FormValue("song-link-input")
    fmt.Println("songURL: ", songURL)

    songInfo,err := spotify.GetSongInfoFromURL(songURL, router.SpotifyClient)
    if err != nil {
        tempError := fmt.Errorf("Error getting song info from url %s: %s.", songURL, err)
        fmt.Println("tempError: ", tempError)
    }
    fmt.Println("songInfo: ", songInfo)

    searchName := string(songInfo.ArtistNames[0] + " - " + songInfo.Name)
    fmt.Println("searchName: ", searchName)
    // results,err := yt.SearchTitle(searchName)
    // if err != nil {
    //     panic(err)
    // }
    // fmt.Println("results: ", results)

    xs,err := router.YoutubeService.Search.List([]string{"snippet",}).Context(r.Context()).MaxResults(10).Q(searchName + "|music video").Do()
    if err != nil {
        panic(fmt.Errorf("Error searching: %s", err))
    }
    fmt.Println("xs: ", xs)
    for _,item := range xs.Items {
        fmt.Println("item: ", item.Id.VideoId)
        fmt.Println("snippet: ", item.Snippet)
    } 

    // searchListCall := youtube.SearchListCall{}
    //
    // searchListCall.Context(r.Context())
    // // searchListCall1 := searchListCall.Context(context.Background())
    // fmt.Println("searchListCall: ", &searchListCall)
    //
    // searchListCall1 := searchListCall.Q("lmao")
    // fmt.Println("searchListCall1: ", searchListCall1)
    //
    //
    // searchListResponse,err := searchListCall1.Do()
    //
    // if err != nil {
    //     panic(fmt.Errorf("Error running search list call: %s", err))
    // }
    //
    // fmt.Println("searchListResponse: ", searchListResponse)

}
