package main

import (
    "fmt"
    "net/http"
    "os"

	"github.com/gorilla/mux"
     "github.com/joho/godotenv"

    "spot2yt/router"
)

var basePasefiles []string = []string{
    "templates/base.html",
    "templates/components/header.html",
    "templates/components/footer.html",
}

const staticDir string = "/static/"

func main(){
    fmt.Println("Spotify to youtube converter")

    godotenv.Load()

    r := mux.NewRouter()
    router := router.NewRouter()

    fmt.Println("router: ", router)


    staticHandler := http.StripPrefix(staticDir, http.FileServer(http.Dir("static/")))
    r.PathPrefix(staticDir).Handler(staticHandler)

    if router.SpotifyClient.Client == nil {
        tempError := fmt.Errorf("Spotify client is null")
        fmt.Println("tempError: ", tempError)
    }

    r.HandleFunc("/", router.HomePage)
    r.HandleFunc("/auth/spotify", router.AuthSpotify)
    r.HandleFunc("/auth/spotify/get", router.GetSpotifyClient)

    port := os.Getenv("PORT")

    fmt.Println("Listening on port ", port)
    panic(http.ListenAndServe(":" + port, r))

}

