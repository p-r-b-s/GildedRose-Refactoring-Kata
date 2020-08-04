package main

import (
  "strings"
)

type Item struct {
	name            string
	sellIn, quality int
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
    switch {
      case strings.HasPrefix(item.name, "Aged Brie"):
        update_brie(item)
      case strings.HasPrefix(item.name, "Sulfuras"):
      //noop
      case strings.HasPrefix(item.name, "Backstage passes"):
        update_passes(item)
      case strings.HasPrefix(item.name, "Conjured"):
        update_conjured(item)
      default:
        update_item(item)
    }
  }
}		

func update_brie (item *Item) {
  item.sellIn = item.sellIn - 1
  if item.quality < 50 {
		item.quality = item.quality + 1
	}
}

func update_passes (item *Item) {
  item.sellIn = item.sellIn - 1
  switch {
    case item.sellIn > 10:
		  item.quality = item.quality + 1
    case item.sellIn > 5:
	  	item.quality = item.quality + 2
    case item.sellIn > 0:
	  	item.quality = item.quality + 3
    default:
      item.quality = 0
  }
  if item.quality > 50 {
		item.quality = 50
	}
}

func update_conjured (item *Item) {
  item.sellIn = item.sellIn - 1
  item.quality = item.quality - 2
}

func update_item (item *Item) {
  item.sellIn = item.sellIn - 1
  item.quality = item.quality - 1
}