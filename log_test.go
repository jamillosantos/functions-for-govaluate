package functions_for_govaluate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jamillosantos/functions-for-govaluate"
	"math"
)

var _ = Describe("log functions", func() {
	Describe("log", func() {
		It("should get correct value with a float param", func() {
			c, err := functions_for_govaluate.Log(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := functions_for_govaluate.Log(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := functions_for_govaluate.Log("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := functions_for_govaluate.Log()
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := functions_for_govaluate.Log(1, 2)
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("log10", func() {
		It("should get correct value with a float param", func() {
			c, err := functions_for_govaluate.Log10(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log10(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := functions_for_govaluate.Log10(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Log10(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := functions_for_govaluate.Log10("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := functions_for_govaluate.Log10()
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := functions_for_govaluate.Log10(1, 2)
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
