package functions_for_govaluate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jamillosantos/functions-for-govaluate"
	"math"
)

var _ = Describe("sqrt function", func() {
	It("should get correct value with a float param", func() {
		c, err := functions_for_govaluate.Sqrt(3.1363122)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(math.Sqrt(3.1363122)))
	})

	It("should get correct value with an integer param", func() {
		c, err := functions_for_govaluate.Sqrt(6)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(math.Sqrt(6)))
	})

	It("should return an param wrong type error", func() {
		_, err := functions_for_govaluate.Sqrt("invalid param value")
		Expect(err).NotTo(BeNil())
		Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
	})

	It("should get an wrong param count error (too few arguments)", func() {
		c, err := functions_for_govaluate.Sqrt()
		Expect(c).To(BeNil())
		Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
		Expect(err).NotTo(BeNil())
	})

	It("should get an wrong param count error (too many arguments)", func() {
		c, err := functions_for_govaluate.Sqrt(1, 2)
		Expect(c).To(BeNil())
		Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
		Expect(err).NotTo(BeNil())
	})
})
