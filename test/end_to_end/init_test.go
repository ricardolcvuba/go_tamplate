package integration

import (
	"go_template/test/shared"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitE2E(t *testing.T) {
	_, _, err := shared.SetupTest()
	if err != nil {
		t.Fatalf("Failed to set up test: %v", err)
	}

	assert.Equal(t, true, true)
}
