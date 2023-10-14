package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// go does not have classes (objects), but just make a new type
// create a new type of 'deck'
// which is a slice of strings

type deck []string

func (d deck) print() { // the receiver syntax says any variable (d) of type 'deck' can call the print function
	//reason cards.print() works, cards gets passed into the function
	// 'd' can be considered like 'this' keyword or 'self'

	for i, card := range d {
		fmt.Println(i, card)
	}

}

// creates and returns a deck
func newDeck() deck { // no need for receiver to call function

	cards := deck{} // lets create two list of card suits and numbers to dynamically create a deck of cards

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	cardValues := []string{"Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King", "Ace"}

	for _, suit := range cardSuits { // underscores tell compiler, we know there is a value here but we do not want to use it
		for _, val := range cardValues {

			cards = append(cards, val+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) { //parenthesis for multiple return values

	return d[:handSize], d[handSize:]

}

func (d deck) toString() string { //receiver and not parameter to call it as cards.toString() better syntax

	return strings.Join([]string(d), ",") //stringify cards to eventually get a []byte to save to hard drive

}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666) // type cast d.string to byte slice and 666 grants read/ write permissions to all r (4) + w(2)

}

func newDeckFromFile(filename string) deck {

	byteslice, err := os.ReadFile(filename)

	if err != nil {

		fmt.Println("Error:", err)

		os.Exit(1)
	}

	stringyDeck := string(byteslice)

	newDeck := strings.Split(stringyDeck, ",")

	return deck(newDeck)
}

func (d deck) shuffleDeck() {

	// for i := range d {

	// 	newIndex := rand.Intn((len(d) - 1))

	// 	d[i], d[newIndex] = d[newIndex], d[i]
	// }

	//random generator returns same sequence ea iteration because golang uses same 'seed' of data

	//use type Rand to specifiy 'seed' or 'source' for true random values

	source := rand.NewSource(time.Now().UnixNano()) //to get an int64 param, use time package and Now() function to get unixnano so seed value becomes unqiue time to get a unique sequence

	r := rand.New(source)

	for i := range d {

		newIndex := r.Intn(len(d) - 1)

		d[i], d[newIndex] = d[newIndex], d[i]

	}
}

func (d deck) getIndices() [][]string {

	var ccc [][]string

	var cc []string

	var c []string

	for i, card := range d {

		cc = append(cc, fmt.Sprint(i))
		c = append(c, card)

	}

	ccc = append(ccc, cc, c)

	return ccc
}

func (d deck) getLength() int {

	return len(d)
}
