package intpkg //same package name as source file

import (
    "testing" //import go package for testing related functionality
    )
var d2 = Card{ "D", 2 }
var s2 = Card{ "S", 2 }
var s3 = Card{ "S", 3 }
var s4 = Card{ "S", 4 }
var s5 = Card{ "S", 5 }
var s6 = Card{ "S", 6 }
var d7 = Card{ "D", 7 }
var s8 = Card{ "S", 8 }
var s9 = Card{ "S", 9 }
var st = Card{ "S", 10 }

var h1 = Hand{s5, s2, d7, s8, s9}
var h2 = Hand{s5, s2, d7, s8, st}
var handThatHasAFlush = Hand{s2, s3, s4, s5, s6}


func Test_shouldFindHighCard(t *testing.T) {
	
	hand := Hand{d2, s2, d7, s8, s9}
	if ( hand.HighCard() != s9) { 
		t.Error("wrong high card")
	}
	hand = Hand{d2, s2, d7, s9, s8}
	if ( hand.HighCard() != s9){
		t.Error("wrong high card")	
	}
}

func Test_shouldCompareTwoNormalHands(t *testing.T) {
	if( h1.Beats(h2)) {
		t.Error("h2 should beat h1");
	}
}

func Test_HandsAreATie(t *testing.T) {
	if ( h1.Tie(h2)) {
		t.Error("Comparing these hands are not a tie");
	}
}

func Test_hasFlush(t *testing.T) {
	if ( ! handThatHasAFlush.HasFlush() ) {
		t.Error("But, this hand does have a flush!")
	}
}


