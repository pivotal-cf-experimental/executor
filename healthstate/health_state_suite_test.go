package healthstate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHealthState(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HealthState Suite")
}
