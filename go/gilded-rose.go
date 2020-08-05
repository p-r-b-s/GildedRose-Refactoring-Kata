package main

import "strings"

type Item struct {
	name            string
	sellIn, quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if strings.HasPrefix(items[i].name, "Conjured") {
			update_conjured(items[i])
		} else {

			if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
				if items[i].quality > 0 {
					if items[i].name != "Sulfuras, Hand of Ragnaros" {
						items[i].quality = items[i].quality - 1
					}
				}
			} else {
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
					if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
						if items[i].sellIn < 11 {
							if items[i].quality < 50 {
								items[i].quality = items[i].quality + 1
							}
						}
						if items[i].sellIn < 6 {
							if items[i].quality < 50 {
								items[i].quality = items[i].quality + 1
							}
						}
					}
				}
			}

			if items[i].name != "Sulfuras, Hand of Ragnaros" {
				items[i].sellIn = items[i].sellIn - 1
			}

			if items[i].sellIn < 0 {
				if items[i].name != "Aged Brie" {
					if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
						if items[i].quality > 0 {
							if items[i].name != "Sulfuras, Hand of Ragnaros" {
								items[i].quality = items[i].quality - 1
							}
						}
					} else {
						items[i].quality = items[i].quality - items[i].quality
					}
				} else {
					if items[i].quality < 50 {
						items[i].quality = items[i].quality + 1
					}
				}
			}
		}
	}

}

func update_conjured(item *Item) {
	item.sellIn = item.sellIn - 1
	item.quality = item.quality - 2
}
