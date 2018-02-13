package functions_for_govaluate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jamillosantos/functions-for-govaluate"
	"math"
)

var _ = Describe("cos functions", func() {
	Describe("cos", func() {
		It("should get correct value with a float param", func() {
			c, err := functions_for_govaluate.Cos(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cos(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := functions_for_govaluate.Cos(3)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cos(3)))
		})

		It("should return an param wrong type error", func() {
			_, err := functions_for_govaluate.Cos("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := functions_for_govaluate.Cos()
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := functions_for_govaluate.Cos(1, 2)
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("acos", func() {
		It("should get correct value with a float param", func() {
			c, err := functions_for_govaluate.Acos(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acos(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := functions_for_govaluate.Acos(1)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acos(1)))
		})

		It("should return an param wrong type error", func() {
			_, err := functions_for_govaluate.Acos("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := functions_for_govaluate.Acos()
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := functions_for_govaluate.Acos(1, 2)
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("cosh", func() {
		It("should get correct value with a float param", func() {
			c, err := functions_for_govaluate.Cosh(0.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cosh(0.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := functions_for_govaluate.Cosh(2)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Cosh(2)))
		})

		It("should return an param wrong type error", func() {
			_, err := functions_for_govaluate.Cosh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := functions_for_govaluate.Cosh()
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := functions_for_govaluate.Cosh(1, 2)
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})

	Describe("acosh", func() {
		It("should get correct value with a float param", func() {
			c, err := functions_for_govaluate.Acosh(3.1363122)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acosh(3.1363122)))
		})

		It("should get correct value with an integer param", func() {
			c, err := functions_for_govaluate.Acosh(6)
			Expect(err).To(BeNil())
			Expect(c).To(Equal(math.Acosh(6)))
		})

		It("should return an param wrong type error", func() {
			_, err := functions_for_govaluate.Acosh("invalid param value")
			Expect(err).NotTo(BeNil())
			Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
		})

		It("should get an wrong param count error (too few arguments)", func() {
			c, err := functions_for_govaluate.Acosh()
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})

		It("should get an wrong param count error (too many arguments)", func() {
			c, err := functions_for_govaluate.Acosh(1, 2)
			Expect(c).To(BeNil())
			Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
			Expect(err).NotTo(BeNil())
		})
	})
})
