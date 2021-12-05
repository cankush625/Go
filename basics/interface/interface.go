package main

import "fmt"

type Song interface {
	addToFavorite() bool
	playSong() bool
	getSongName() string
}

type BollywoodSong struct {
	name   string
	length int
	singer string
	writer string
}

func (s *BollywoodSong) addToFavorite() bool {
	fmt.Println("Song " + s.name + " is added to favorite song list")
	return true
}

func (s BollywoodSong) playSong() bool {
	fmt.Println("Playing song " + s.name)
	return true
}

func (s BollywoodSong) getSongName() string {
	return s.name
}

func main() {
	var mySong Song

	bs := BollywoodSong{
		name:   "abs",
		length: 5,
		singer: "sfd",
		writer: "sdf",
	}

	mySong = &bs

	mySong.addToFavorite()
	mySong.playSong()
	fmt.Println(mySong.getSongName())

	// Find interface values
	// Under the hood, interface values can be thought of as a tuple of a value and a concrete type
	// Calling a method on an interface value executes the method of the same name on its underlying type
	fmt.Printf("(%v, %T)\n", mySong, mySong)
}
