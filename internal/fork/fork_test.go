package fork

import (
	"testing"
)

func TestNewForkCreatesForkWithID(t *testing.T) {
	id := 1
	fork := NewFork(id)
	if fork.ID != id {
		t.Errorf("expected fork ID to be %d, got %d", id, fork.ID)
	}
}

func TestForkLockUnlock(t *testing.T) {
	fork := NewFork(1)
	fork.Lock()

	if fork.TryLock() {
		t.Errorf("expected fork to be locked")
	}
	fork.Unlock()
	if !fork.TryLock() {
		t.Errorf("expected fork to be unlocked")
	}
	fork.Unlock()
}
