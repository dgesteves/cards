package main

func main() {
	d := newDeck()
	d.shuffle()
	d.deal(2)
	err := d.saveToFile("my_cards")
	if err != nil {
		return
	}
	d.print()
	d2 := newDeckFromFile("my_cards")
	d2.print()
}
