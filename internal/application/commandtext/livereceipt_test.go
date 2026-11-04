package commandtext_test

import (
	"testing"

	"github.com/IgnisDevNE/SuiteWard/internal/application/commandtext"
)

// A proposal number of zero is not one that Reference writes, so it is not read as a reference.
func TestParseReferenceRefusesAProposalOfZero(t *testing.T) {
	if _, _, ok := commandtext.ParseReference("P0-R1"); ok {
		t.Error("P0-R1 was read as a reference")
	}
}
