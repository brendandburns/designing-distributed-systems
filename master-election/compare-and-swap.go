package election

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}
var store = map[string]string{}

func compareAndSwap(key, nextValue, currentValue string) (bool, error) {
	lock.Lock()
	defer lock.Unlock()
	if _, found := store[key]; !found {
		if len(currentValue) == 0 {
			store[key] = nextValue
			return true, nil
		}
		return false, fmt.Errorf("Expected value %s for key %s, but found empty", currentValue, key)
	}
	if store[key] == currentValue {
		store[key] = nextValue
		return true, nil
	}
	return false, nil
}
