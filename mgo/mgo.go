package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Track struct {
	Artist string
	Title  string
}

type Playlist struct {
	Title  string
	Tracks []Track
}

func main() {

	// Open session with MongoDB.
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Connect with the playlists collection.
	c := session.DB("test").C("playlists")

	// Add a playlist to the collection.
	items := make([]Track, 4)
	items[0] = Track{"Airhead", "Pyramid Lake"}
	items[1] = Track{"DJ Butcher", "We Are The Music"}
	items[2] = Track{"Dragon", "Jeremy Kyle"}
	items[3] = Track{"Leo Zero", "It’s Time"}
	err = c.Insert(&Playlist{"The Smugglers Inn Podcast 03/08/12", items})
	if err != nil {
		panic(err)
	}

	// TODO: Add more playlists.

	// Search for playlist. TODO: Fuzzy search & multiple results.
	result := Playlist{}
	err = c.Find(bson.M{"title": "The Smugglers Inn Podcast 03/08/12"}).One(&result)
	if err != nil {
		panic(err)
	}

	// Print playlist.
	fmt.Println("Tracks:", result.Tracks)
}
