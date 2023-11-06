//Crie uma struct chamada Playlist com os campos "nome" e "músicas".
//O campo "músicas" deve ser um slice de structs, cada uma representando
//uma música com os campos "título", "artista" e "duração". Escreva uma
//função que receba uma Playlist como parâmetro e imprima o nome da playlist,
//o tempo total da playlist e as informações de cada música.

package main

import (
	"fmt"
	"time"
)

type Musica struct {
	Título  string
	Artista string
	Duração time.Duration
}
type Playlist struct {
	Nome    string
	Músicas []Musica
}

func imprimirPlaylist(pl Playlist) {
	fmt.Printf("Nome da Playlist: %s\n", pl.Nome)
	fmt.Println("Músicas:")
	for i, música := range pl.Músicas {
		fmt.Printf("Música %d:\n", i+1)
		fmt.Printf("Título: %s\n", música.Título)
		fmt.Printf("Artista: %s\n", música.Artista)
		fmt.Printf("Duração: %s\n", música.Duração)
	}
}

func main() {
	playlist := Playlist{
		Nome: "Minha Playlist",
		Músicas: []Musica{
			{"Música 1", "Artista 1", 4*time.Minute + 30*time.Second},
			{"Música 2", "Artista 2", 3*time.Minute + 45*time.Second},
		},
	}

	imprimirPlaylist(playlist)
}
