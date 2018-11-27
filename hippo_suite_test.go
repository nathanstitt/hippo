package hippo

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDis(t *testing.T) {
	BeforeEach(func() {
		TestingEnvironment.DBConnectionUrl = "postgres://localhost/hippo_dev?sslmode=disable"
	})
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hippo Test Suite")
}
