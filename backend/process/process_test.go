package process

import (
	"strings"
	"testing"
)

func TestProcesses(t *testing.T) {
	// As long there is a
	p := getProcessList()

	if len(p) <= 0 {
		t.Fatal("should have processes")
	}

	// Find at least one process
	found := false
	for _, p1 := range p {
		// a running application must include go
		if strings.Contains(p1.Command, "go") {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("should have Go")
	}
}
