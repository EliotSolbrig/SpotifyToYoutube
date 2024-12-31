package router

import (
    "fmt"
    "net/http"
)

func (router *Router) ConvertSong(w http.ResponseWriter, r *http.Request){
    fmt.Println(r.Method, " request on /convert")

    r.ParseForm()
    songURL := r.FormValue("song-link-input")
    fmt.Println("songURL: ", songURL)
}
