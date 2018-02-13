package functions_for_govaluate

import (
	"testing"
	"github.com/jamillosantos/macchiato"
	"github.com/onsi/gomega"
	"github.com/onsi/ginkgo"
)

func TestFunctions(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	macchiato.RunSpecs(t, "Functions For Govaluate")
}
