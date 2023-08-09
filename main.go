package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Name    string
	Pokemon []Pokemon
}

type Pokemon struct {
	EntryNo int
	Species PokemonSpecies
}

type PokemonSpecies struct {
	Name string
}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.Species.Name)
	}

}
