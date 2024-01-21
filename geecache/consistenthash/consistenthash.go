package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32

type Map struct {
	// the hashing function that can be customized by the user
	hash Hash
	// the number of virtual nodes
	replicas int
	// the consistent hash ring
	keys []int
	// hashMap mapping the virtual nodes with real nodes
	hashMap map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some keys to the hash
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for ii := 0; ii < m.replicas; ii++ {
			hash := int(m.hash([]byte(strconv.Itoa(ii) + key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key
func (m *Map) Get(key string) string {
	if len(m.keys) == 0 {
		return ""
	}
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(ii int) bool {
		return m.keys[ii] >= hash
	})
	return m.hashMap[m.keys[idx%len(m.keys)]]
}
