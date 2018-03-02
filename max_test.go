package functions_for_govaluate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jamillosantos/functions-for-govaluate"
)

var _ = Describe("max, min functions", func() {
	Describe("max", func() {
		It("should resolve a dual max", func() {
			c, err := functions_for_govaluate.Max(1, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(2.0))
		})

		It("should resolve a the max value for 5 values", func() {
			c, err := functions_for_govaluate.Max(4, 5, 3, 1, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(5.0))
		})

		It("should fail with a param type error (string supplied)", func() {
			c, err := functions_for_govaluate.Max(4, 5, 3, 1, 2, "non number")
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param type error (boolean supplied)", func() {
			c, err := functions_for_govaluate.Max(4, 5, 3, 1, 2, false)
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param count error", func() {
			c, err := functions_for_govaluate.Max()
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
		})
	})

	Describe("min", func() {
		It("should resolve a dual min", func() {
			c, err := functions_for_govaluate.Min(2, 1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should resolve a the min value for 5 values", func() {
			c, err := functions_for_govaluate.Min(4, 5, 3, 1, 2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(1.0))
		})

		It("should fail with a param type error (string supplied)", func() {
			c, err := functions_for_govaluate.Min(4, 5, 3, 1, 2, "non number")
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param type error (boolean supplied)", func() {
			c, err := functions_for_govaluate.Min(4, 5, 3, 1, 2, false)
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should fail with a param count error", func() {
			c, err := functions_for_govaluate.Min()
			Expect(err).NotTo(BeNil())
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
		})
	})
})
