package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Player struct {
	Team string `json:"team"`
	Home string `json:"home"`
	From string `json:"from"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Realizar una solicitud GET al backend para obtener la lista de jugadores
		resp, err := http.Get("http://localhost:8080/players") // Reemplaza con la URL de tu backend
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// Leer la respuesta del backend
		var players []Player
		err = json.NewDecoder(resp.Body).Decode(&players)
		if err != nil {
			log.Fatal(err)
		}

		// Realizar un seguimiento del equipo que llegó primero al home
		firstTeam := ""
		secondTeam := ""
		for _, player := range players {
			if firstTeam == "" {
				firstTeam = player.Team
			} else if secondTeam == "" {
				secondTeam = player.Team
				break
			}
		}

		// Definir la plantilla HTML
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Hula HOOP</title>
		</head>
		<body>
			<h1>Cobras VS Leones</h1>
			<ul>
				{{range .Players}}
				<li>
					<p><strong>Equipo:</strong> {{.Team}}</p>
					<p><strong>Casa:</strong> {{.Home}}</p>
					<p><strong>Desde:</strong> {{.From}}</p>
				</li>
				{{end}}
			</ul>
			{{if .FirstTeam}}<p><strong>El equipo ganador es:</strong> {{.FirstTeam}}</p>{{end}}
			{{if .SecondTeam}}<p><strong>El equipo perdedor es:</strong> {{.SecondTeam}}</p>{{end}}
		</body>
		</html>`

		// Crear la plantilla a partir del HTML
		tmpl := template.Must(template.New("playerList").Parse(html))

		// Crear un struct para pasar los datos a la plantilla
		data := struct {
			Players    []Player
			FirstTeam  string
			SecondTeam string
		}{
			Players:    players,
			FirstTeam:  firstTeam,
			SecondTeam: secondTeam,
		}

		// Renderizar la plantilla con los datos
		err = tmpl.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
