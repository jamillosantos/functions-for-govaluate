package functions_for_govaluate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jamillosantos/functions-for-govaluate"
	"math"
)

var _ = Describe("sin functions", func() {
	Describe("sin", func() {
		It("should get correct value with a float param", func() {
			c, err := functions_for_govaluate.Sin(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sin(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := functions_for_govaluate.Sin(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sin(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := functions_for_govaluate.Sin("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := functions_for_govaluate.Sin()
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := functions_for_govaluate.Sin(1, 2)
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("asin", func() {
		It("should get correct value with a float param", func() {
			c, err := functions_for_govaluate.Asin(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asin(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := functions_for_govaluate.Asin(1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asin(1)))
		})

		It("should return an param wrong type error", func() {
			_, err := functions_for_govaluate.Asin("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := functions_for_govaluate.Asin()
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := functions_for_govaluate.Asin(1, 2)
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("sinh", func() {
		It("should get correct value with a float param", func() {
			c, err := functions_for_govaluate.Sinh(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sinh(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := functions_for_govaluate.Sinh(2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Sinh(2)))
		})

		It("should return an param wrong type error", func() {
			_, err := functions_for_govaluate.Sinh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := functions_for_govaluate.Sinh()
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := functions_for_govaluate.Sinh(1, 2)
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("asinh", func() {
		It("should get correct value with a float param", func() {
			c, err := functions_for_govaluate.Asinh(3.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asinh(3.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := functions_for_govaluate.Asinh(6)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Asinh(6)))
		})

		It("should return an param wrong type error", func() {
			_, err := functions_for_govaluate.Asinh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := functions_for_govaluate.Asinh()
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := functions_for_govaluate.Asinh(1, 2)
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
