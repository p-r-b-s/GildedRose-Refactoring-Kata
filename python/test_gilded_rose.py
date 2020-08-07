# -*- coding: utf-8 -*-
import unittest

from gilded_rose import Item, GildedRose

class GildedRoseTest(unittest.TestCase):

  backstage_pass_array_quality = [(30,1)]
  #You are here. We're looking at a tuple that has the SellIn and Q values

    def test_update_normal_item(self):
        items = [Item("foo", 1, 1)]
        gilded_rose = GildedRose(items)
        gilded_rose.update_quality()
        self.assertEquals("foo", items[0].name)
        self.assertEquals(0, items[0].sell_in)
        self.assertEquals(0, items[0].quality)

    def test_quality_never_negative(self):
        items = [Item("foo", 1, 0)]
        gilded_rose = GildedRose(items)
        gilded_rose.update_quality()
        self.assertEquals("foo", items[0].name)
        self.assertEquals(0, items[0].sell_in)
        self.assertEquals(0, items[0].quality)

    def test_aged_brie_increases_in_quality(self):
        items = [Item("Aged Brie", 1, 0)]
        gilded_rose = GildedRose(items)
        gilded_rose.update_quality()
        self.assertEquals("Aged Brie", items[0].name)
        self.assertEquals(0, items[0].sell_in)
        self.assertEquals(1, items[0].quality)

    def test_quality_max_limit(self):
        items = [Item("Aged Brie", 1, 50)]
        gilded_rose = GildedRose(items)
        gilded_rose.update_quality()
        self.assertEquals("Aged Brie", items[0].name)
        self.assertEquals(0, items[0].sell_in)
        self.assertEquals(50, items[0].quality)

    # def test_quality_max_limit(self):
    #     items = [Item("foo", 1, 52)]
    #     gilded_rose = GildedRose(items)
    #     gilded_rose.update_quality()
    #     self.assertEquals("foo", items[0].name)
    #     self.assertEquals(0, items[0].sell_in)
    #     self.assertEquals(50, items[0].quality)

if __name__ == '__main__':
    unittest.main()
