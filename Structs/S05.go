//Usando as mesmas structs do exercício anterior, escreva uma função que receba um título
//de música como parâmetro e retorne uma lista das Playlists que possuem esse título.

package main

import (
	"fmt"
	"time"
)

type Musica2 struct {
	Título   string
	Artista  string
	Duração  time.Duration
}

type Playlist2 struct {
	Nome    string
	Músicas []Musica
}

func buscarPlaylistsPorTítulo(título string, playlists []Playlist) []Playlist {
	var playlistsEncontradas []Playlist

	for _, playlist := range playlists {
		for _, música := range playlist.Músicas {
			if música.Título == título {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}

	return playlistsEncontradas
}

func main() {
	playlist1 := Playlist{
		Nome: "Minha Playlist 1",
		Músicas: []Musica{
			{ "Música 1", "Artista 1", 4 * time.Minute + 30 * time.Second },
			{ "Música 2", "Artista 2", 3 * time.Minute + 45 * time.Second },
		},
	}

	playlist2 := Playlist{
		Nome: "Minha Playlist 2",
		Músicas: []Musica{
			{ "Música 2", "Artista 2", 3 * time.Minute + 45 * time.Second },
			{ "Música 3", "Artista 3", 5 * time.Minute + 15 * time.Second },
		},
	}

	playlists := []Playlist{playlist1, playlist2}

	títuloDesejado := "Música 2"
	playlistsEncontradas := buscarPlaylistsPorTítulo(títuloDesejado, playlists)

	fmt.Printf("Playlists que contêm a música com título '%s':\n", títuloDesejado)
	for _, playlist := range playlistsEncontradas {
		fmt.Printf("- %s\n", playlist.Nome)
	}
}
