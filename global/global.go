package global

import (
	"sync"
)

const MAX_GOROUTINES = 5

var JobIdsMap sync.Map
