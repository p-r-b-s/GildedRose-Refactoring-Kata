package main

import "testing"

func Test_Conjured(t *testing.T) {
	var items = []*Item{
		&Item{"Conjured fluff", 1, 2},
	}

	UpdateQuality(items)

	if items[0].name != "Conjured fluff" {
		t.Errorf("Name: Expected %s but got %s ", "fixme", items[0].name)
	}
	if items[0].sellIn != 0 {
		t.Errorf("sellIn: Expected %d but got %d ", 0, items[0].sellIn)
	}
	if items[0].quality != 0 {
		t.Errorf("quality: Expected %d but got %d ", 0, items[0].quality)
	}
}
