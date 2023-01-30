package channeld

import (
	"sync"

	"channeld.clewcat.com/channeld/pkg/common"
)

func GetNextId(m *map[uint32]interface{}, start uint32, min uint32, max uint32) (uint32, bool) {
	for i := min; i <= max; i++ {
		if _, exists := (*m)[start]; !exists {
			return start, true
		}
		if start < max {
			start++
		} else {
			start = min
		}
	}

	return 0, false
}

func GetNextIdSync(m *sync.Map, start common.ChannelId, min common.ChannelId, max common.ChannelId) (common.ChannelId, bool) {
	for i := min; i <= max; i++ {
		if _, exists := m.Load(start); !exists {
			return start, true
		}
		if start < max {
			start++
		} else {
			start = min
		}
	}

	return 0, false
}

func HashString(s string) uint32 {
	hash := uint32(17)
	for c := range s {
		hash = hash*31 + uint32(c)
	}
	return hash
}

func Pointer[K any](val K) *K {
	return &val
}
