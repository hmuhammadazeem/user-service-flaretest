package bloom


type Filter interface {
	Add(key string)
	MayExist(key string) bool
}

// BloomFilter: A naive implementation, experimenting with first time; below code is not working
// Todo: Estimate bits, number of hash functions
// Todo: Thread safe implementation
type BloomFilter struct {
	data    []byte
	hashers []Hasher
}

type Hasher func(key string) int64


// func (b *BloomFilter) Add(key string) {
// 	for _, hasher := range b.hashers {
// 		bit := hasher(key)
// 		b.data[bit] = 1
// 	}
// }

// func (b *BloomFilter) MayExist(key string) bool {
// 	for _, hasher := range b.hashers {
// 		bit := hasher(key)
// 		if b.data[bit] != 1 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func NewBloom() BloomFilter {

// }