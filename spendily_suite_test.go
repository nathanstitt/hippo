package hippo_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSpendily(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spendily Suite")
}
