package main

import (
	"fmt"
	"github.com/xiaomingping/landlord/poker"
)



func main() {
	s := poker.CreateDeck()
	s.CountCards()
	fmt.Println(s)
}
