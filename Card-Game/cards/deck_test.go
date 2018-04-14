package main

import "testing"


func TestNewDeck(t *testing.T){
	d:= newDeck()
	if len(d)!= 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
	 t.Errorf("Expected first card of Ace of Spades, but got %v",d[0])
	}

	if d[51] != "King of Clubs" {
	t.Errorf("Expected last card to be the King of Clubs %v", d[51])
	}

	if d[len(d)-1] != "King of Clubs" {
	t.Errorf("Expected last card to be King of CLubs %v", d[len(d)])
	}


}
