package temporal

import (
	"errors"
	"sync"

	"go.temporal.io/sdk/client"
)

var (
	instanceOrchestrator client.Client
	onceOrchestrator     sync.Once
)

func NewTemporalClientInstance() (client.Client, error) {
	onceOrchestrator.Do(func() {
		c, err := client.Dial(client.Options{})
		if err != nil {
			return
		}
		instanceOrchestrator = c
	})
	if instanceOrchestrator == nil {
		return nil, errors.New("temporal client instanceOrchestrator is not initialized")
	}
	return instanceOrchestrator, nil
}
