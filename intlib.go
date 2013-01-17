package intpkg

import "sort"

type Card struct {
	Suit string
	Rank int
}

type Hand []Card

func (h Hand) Len() int {
	return len(h)
} 

func (h Hand) Less(i, j int) bool {
	return h[i].Rank > h[j].Rank
}

func (h Hand) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hand) HighCard() Card {
	sort.Sort(h)
	return h[0]
}

func (h Hand) Beats( other Hand ) bool {
  return other.HighCard().Rank < h.HighCard().Rank
}

func (h Hand) Tie( other Hand ) bool {
	return other.HighCard().Rank == h.HighCard().Rank
}

func (h Hand) HasFlush() bool {
	return true
}
