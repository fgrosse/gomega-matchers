package matchers

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestRegisterExtendedTestingT(t *testing.T) {
	t.Skip()
	RegisterExtendedTestingT(t)
	Expect(false).To(BeTrue())
}
