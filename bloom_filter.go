package bloomfilter

import "golang.org/x/exp/constraints"

type BloomFilter[K constraints.Integer] uint64

func (f *BloomFilter[K]) Add(item K) {
	*f = *f | BloomFilter[K](item)
}

func (f *BloomFilter[K]) MayExist(item K) bool {
	return ((*f) & BloomFilter[K](item)) != 0
}
