package delay_lock

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	mutex := New(time.Duration(3) * time.Second)
	if mutex == nil {
		t.Errorf("it should not be nil")
	}
}
