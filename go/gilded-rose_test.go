package main

import "testing"

type Test_Case struct {
	name                              string
  item                              Item
	expected_sellIn, expected_quality int
}

func Test_cases (t *testing.T) {
  var cases = []*Test_Case{
    &Test_Case{"normal", Item{"Normal Item", 2, 2},1,1},
    &Test_Case{"brie", Item{"Aged Brie", 2, 2},1,3},
    &Test_Case{"brie q limit", Item{"Aged Brie", 2, 50},1,50},
    &Test_Case{"sulfuras", Item{"Sulfuras, Hand of Ragnaros", 2, 20},2,20},
    &Test_Case{"sulfuras other", Item{"Sulfuras, Foot of Bifrost Giant", 2, 20},2,20},
    &Test_Case{"backstage",  Item{"Backstage passes to a TAFKAL80ETC concert", 12, 20},11,21},
    &Test_Case{"backstage other",  Item{"Backstage passes to a BIG concert", 12, 20},11,21},
    &Test_Case{"backstage eleven",  Item{"Backstage passes to a TAFKAL80ETC concert", 11, 20},10,21},
    &Test_Case{"backstage ten",  Item{"Backstage passes to a TAFKAL80ETC concert", 10, 20},9,22},
    &Test_Case{"backstage six",  Item{"Backstage passes to a TAFKAL80ETC concert", 6, 20},5,22},
    &Test_Case{"backstage five",  Item{"Backstage passes to a TAFKAL80ETC concert", 5, 20},4,23},
    &Test_Case{"backstage one",  Item{"Backstage passes to a TAFKAL80ETC concert", 1, 20},0,23},
    &Test_Case{"backstage zero",  Item{"Backstage passes to a TAFKAL80ETC concert", 0, 20},-1,0},
    &Test_Case{"backstage q limit",  Item{"Backstage passes to a TAFKAL80ETC concert", 11, 50},10,50},
    &Test_Case{"backstage ten q limit",  Item{"Backstage passes to a TAFKAL80ETC concert", 10, 50},9,50},
    &Test_Case{"backstage five q limit",  Item{"Backstage passes to a TAFKAL80ETC concert", 5, 50},4,50},
    &Test_Case{"conjured",  Item{"Conjured cakes", 5, 10},4,8},

  }

  for _, tc := range cases {
    Helper(t, tc)
  }

}

func Helper(t *testing.T, test_case *Test_Case) {
	var items = []*Item{
    &test_case.item,
	}
	UpdateQuality(items)
	if items[0].name != test_case.item.name {
    t.Errorf("TestName: %s name: Expected %s but got %s ", test_case.name, test_case.item.name, items[0].name)
  }
  if items[0].sellIn != test_case.expected_sellIn {
		t.Errorf("TestName: %s sellIn: Expected %d but got %d ", test_case.name, test_case.expected_sellIn, items[0].sellIn)
	}
	if items[0].quality != test_case.expected_quality {
		t.Errorf("TestName: %s quality: Expected %d but got %d ", test_case.name, test_case.expected_quality, items[0].quality)
	}
}