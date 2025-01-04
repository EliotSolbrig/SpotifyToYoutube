package router

import (
	// "context"
	"fmt"
	"net/http"
    // "encoding/json"
    "html/template"

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
        thumbnails := item.Snippet.Thumbnails
        fmt.Println("thumbnails: ", thumbnails)
    } 

    firstResult := xs.Items[0]

    convertedInfo := map[string]any{
        "ID": firstResult.Id.VideoId,
        "Title": firstResult.Snippet.Title,
        "Thumbnail": firstResult.Snippet.Thumbnails.High.Url,
        "Url": "https://www.youtube.com/watch?v=" + firstResult.Id.VideoId,
        "EmbedUrl": "https://www.youtube.com/embed/" + firstResult.Id.VideoId,
    }
    fmt.Println("convertedInfo: ", convertedInfo)

    templates := []string{"templates/components/songinfo.html",}
    tmpl := template.Must(template.ParseFiles(templates...))
    err = tmpl.ExecuteTemplate(w, "songinfo", convertedInfo)

    if err != nil {
        panic(fmt.Errorf("Error executing songinfo template: %s", err))
    }

    // response,err := json.Marshal(convertedInfo)
    // if err != nil {
    //     panic(fmt.Errorf("Error marshaling response: %s", err))
    // }
    //
    // w.Write(response)
}
