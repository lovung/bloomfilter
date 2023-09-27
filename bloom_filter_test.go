package bloomfilter

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	var bf BloomFilter[int]

	// Add some items to the Bloom filter
	itemsToAdd := []int{1, 2, 3, 4, 5}
	for _, item := range itemsToAdd {
		bf.Add(item)
	}

	// Check if the added items may exist in the Bloom filter
	for _, item := range itemsToAdd {
		if !bf.MayExist(item) {
			t.Errorf("Item %d should exist in the Bloom filter", item)
		}
	}

	// Check if non-added items don't falsely appear to exist
	nonExistentItems := []int{0, 8}
	for _, item := range nonExistentItems {
		if bf.MayExist(item) {
			t.Errorf("Item %d falsely appears to exist in the Bloom filter", item)
		}
	}

	// Check if non-added items falsely appear to exist
	nonExistentItems = []int{6, 7}
	for _, item := range nonExistentItems {
		if !bf.MayExist(item) {
			t.Errorf("Item %d falsely appears to exist in the Bloom filter", item)
		}
	}
}
