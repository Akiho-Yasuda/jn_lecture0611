package main

import (
    "log"
    "net/http"
    "encoding/json"
)


type Intro struct {
    Name        string  `json:"name"`
    KG          string  `json:"kg"`
    LoginName   string  `json:"login_name"`
    Birthday    string  `json:"birthday"`
    Hobby       string  `json:"hobby"`
}

func Root(w http.ResponseWriter, r *http.Request) {
    data := Intro{
        Name: "Akiho-Yasuda",
        KG: "HAISYS",
        LoginName: "suyaa",
        Birthday: "2000.02.07",
        Hobby: "PA",
    }

    json.NewEncoder(w).Encode(data)
}

func main() {
    http.HandleFunc("/", Root)

    log.Fatal(http.ListenAndServe(":8000", nil))
}
