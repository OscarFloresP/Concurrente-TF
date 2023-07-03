package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"net"
	"time"
	"net/http"
	"bytes"
	"log"
	
)

type Player struct {
	Team string `json:"team"`
	Home string `json:"home"`
	From string `json:"from"`
}

type Message struct {
	Cmd        string `json:"cmd"`
	Hostname   string `json:"hostname"`
	Contestant Player `json:"player"`
}

type Info struct {
	team       string
	hostname   string
	prev       string
	next       string
	challenger *Player
}

func listen(hostname string, chInfo chan Info) {
	if ln, err := net.Listen("tcp", hostname); err == nil {
		defer ln.Close()
		fmt.Println("Listening...")
		for {
			if cn, err := ln.Accept(); err == nil {
				go handle(cn, chInfo)
			}
		}
	}
}

func handle(cn net.Conn, chInfo chan Info) {
	defer cn.Close()
	fmt.Printf("Connection accepted from %s\n", cn.RemoteAddr())
	msg := &Message{}
	dec := json.NewDecoder(cn)
	if err := dec.Decode(msg); err == nil {
		//fmt.Println(msg)
		switch msg.Cmd {
		case "jump":
			// Crear una solicitud POST al backend con el JSON
			url := "http://localhost:8080/json" // Reemplaza con la URL de tu backend
			contestantJSON, err := json.Marshal(msg.Contestant)
			if err != nil {
    			log.Fatal(err)
			}
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(contestantJSON))

			//req, err := http.NewRequest("POST", url, bytes.NewBuffer(dec.Decode(msg.Contestant)))
			if err != nil {
				log.Fatal(err)
			}
			req.Header.Set("Content-Type", "application/json")
		
			// Realizar la solicitud al backend
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
		
			// Leer la respuesta del backend
			var response map[string]string
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				log.Fatal(err)
			}
			info := <-chInfo
			fmt.Println(info, msg)
			enc := json.NewEncoder(cn)
			if err := enc.Encode(Message{Cmd: "ok"}); err != nil {
				fmt.Printf("Can't encode OK REPLY\n%s\n", err)
			}
			player := msg.Contestant
			if info.challenger != nil {
				var loser Player
				if rand.Intn(100) >= 50 {
					loser = player
					player = *info.challenger
				} else {
					loser = *info.challenger
				}
				send(loser.Home, Message{Cmd: "send new", Hostname: info.hostname},
					func(cn net.Conn) {})
			}
			if info.next == "" || info.prev == "" {
				fmt.Printf("Ganaron los del equipo %s\n", player.Team)
				return
			}
			var remote string
			if player.From == info.prev {
				remote = info.next
			} else {
				remote = info.prev
			}
			player.From = info.hostname
			needToFreeInfo := true
			send(remote, Message{"jump", info.hostname, player}, func(cn2 net.Conn) {
				duration := time.Second * 3
				if err := cn2.SetReadDeadline(time.Now().Add(duration)); err != nil {
					fmt.Printf("SetReadDeadline failed:\n%s\n", err)
					panic("OMG!")
				}

				dec := json.NewDecoder(cn2)
				msg2 := &Message{}
				if err := dec.Decode(msg2); err == nil {
					fmt.Println("Se supone que recibimos OK")
				} else {
					if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
						if msg.Hostname == info.next {
							info.challenger = &msg.Contestant
							fmt.Println("liberando ando info func send jump")
							needToFreeInfo = false
							chInfo <- info
						}
						fmt.Printf("read timeout:\n%s\n", err)
					} else {
						fmt.Printf("read error:\n%s\n", err)
					}
				}
			})
			fmt.Println("liberando ando info end of jump")
			if needToFreeInfo {
				chInfo <- info
			}
		case "send new":
			info := <-chInfo
			var remote string
			player := Player{Home: info.hostname, From: info.hostname}
			if info.prev == "" {
				remote = info.next
				player.Team = "Cobras"
			} else {
				remote = info.prev
				player.Team = "Leones"
			}
			fmt.Printf("Sending new player from %s\n", info.team)
			needToFreeInfo := true
			send(remote, Message{"jump", info.hostname, player}, func(cn2 net.Conn) {
				duration := time.Second
				if err := cn2.SetReadDeadline(time.Now().Add(duration)); err != nil {
					fmt.Printf("SetReadDeadline failed:\n%s\n", err)
					panic("OMG!")
				}

				dec := json.NewDecoder(cn2)
				msg2 := &Message{}
				if err := dec.Decode(msg2); err == nil {
					fmt.Println("Se supone que recibimos OK")
				} else {
					if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
						if msg.Hostname == info.next {
							info.challenger = &msg.Contestant
							fmt.Println("liberando ando info func send send new")
							needToFreeInfo = false
							chInfo <- info
						}
						fmt.Printf("read timeout:\n%s\n", err)
					} else {
						fmt.Printf("read error:\n%s\n", err)
					}
				}
			})
			fmt.Println("liberando ando info end of send new")
			if needToFreeInfo {
				chInfo <- info
			}
		}
	} else {
		fmt.Printf("Couldn't decode: %s\n", err)
	}
}

func send(remote string, msg Message, f func(cn net.Conn)) {
	if cn, err := net.Dial("tcp", remote); err == nil {
		defer cn.Close()
		enc := json.NewEncoder(cn)
		if err := enc.Encode(msg); err == nil {
			f(cn)
		} else {
			fmt.Printf("Couldn't enconde %s\n", err)
		}
	} else {
		fmt.Printf("Failed to send: %s\n", err)
	}
}

func main() {
	special := flag.Bool("s", false, "The special flag for testing stuff.")
	hostname := flag.String("h", "", "IP/Hostname:port to listen on.")
	prevRemote := flag.String("p", "", "Previous node. If empty, home team 1")
	nextRemote := flag.String("n", "", "Next node. If empty, home team 2")
	flag.Parse()

	if *special { // TODO assuming that -n and -p are set, check
		go send(*nextRemote, Message{Cmd: "send new"}, func(cn net.Conn) {})
		go send(*prevRemote, Message{Cmd: "send new"}, func(cn net.Conn) {})
		time.Sleep(time.Second)
		return
	}

	if *hostname == "" || (*prevRemote == "" && *nextRemote == "") {
		flag.PrintDefaults()
		return
	}

	var team string
	if *prevRemote == "" {
		team = "Team 1: Cobras"
	} else {
		team = "Team 2: Leones"
	}

	chInfo := make(chan Info, 1)
	chInfo <- Info{team, *hostname, *prevRemote, *nextRemote, nil}
	listen(*hostname, chInfo)
}
