package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

type Player struct {
    ID    int    `json:"id"`
    Team  string `json:"team"`
    Home  string `json:"home"`
    From  string `json:"from"`
}

var players = make(map[int]Player)
var nextID = 1

func handleJSON(w http.ResponseWriter, r *http.Request) {
    var player Player
    err := json.NewDecoder(r.Body).Decode(&player)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    player.ID = nextID
    players[player.ID] = player
    nextID++

    fmt.Println("ID:", player.ID)
    fmt.Println("Team:", player.Team)
    fmt.Println("Home:", player.Home)
    fmt.Println("From:", player.From)

    response := map[string]string{"message": "JSON received and player saved successfully"}
    json.NewEncoder(w).Encode(response)
}

func getPlayers(w http.ResponseWriter, r *http.Request) {
    var playerList []Player
    for _, player := range players {
        playerList = append(playerList, player)
    }

    json.NewEncoder(w).Encode(playerList)
}

func main() {
    http.HandleFunc("/json", handleJSON)
    http.HandleFunc("/players", getPlayers)

    log.Fatal(http.ListenAndServe(":8080", nil))
}