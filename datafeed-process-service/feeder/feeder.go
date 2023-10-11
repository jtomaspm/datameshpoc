package feeder

import (
	"sync"
)

type Feeder interface {
	Feed(wg *sync.WaitGroup)
}
