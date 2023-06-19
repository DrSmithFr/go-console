package main

import (
	"encoding/json"
	"github.com/DrSmithFr/go-console"
	"github.com/DrSmithFr/go-console/table"
)

func getMyJSONBytes() []byte {
	data := struct {
		// json tags are optionally but if set they are being used for the headers on `PrintJSON`.
		Firstname string `json:"first name"`
		Lastname  string `json:"last name"`
	}{"Georgios", "Callas"}
	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}

	return b
}

func main() {
	cmd := go_console.NewScript().Build()

	jsonData := getMyJSONBytes()

	tab := table.
		NewTable().
		ParseJSON(jsonData)

	render := table.
		NewRender(cmd.Output).
		SetContent(tab)

	render.Render()
}
