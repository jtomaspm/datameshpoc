package worker

import "sync"

type Worker interface {
	Run(wg *sync.WaitGroup)
}
