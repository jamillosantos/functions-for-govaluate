package functions_for_govaluate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jamillosantos/functions-for-govaluate"
)

var _ = Describe("is_defined function", func() {
	It("should return that a nil value is not defined", func() {
		c, err := functions_for_govaluate.IsDefined(nil)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a non empty string is defined", func() {
		c, err := functions_for_govaluate.IsDefined("this is a non empty string")
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a non empty string is defined", func() {
		c, err := functions_for_govaluate.IsDefined("")
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a zero valued integer is not defined", func() {
		c, err := functions_for_govaluate.IsDefined(0)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a non zero valued integer is defined (positive)", func() {
		c, err := functions_for_govaluate.IsDefined(1)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a non zero valued integer is defined (negative)", func() {
		c, err := functions_for_govaluate.IsDefined(-1)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a zero valued float is not defined", func() {
		c, err := functions_for_govaluate.IsDefined(0.0)
		Expect(err).To(BeNil())
		Expect(c).To(BeFalse())
	})

	It("should return that a non zero valued float is defined (positive)", func() {
		c, err := functions_for_govaluate.IsDefined(0.000001)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return that a non zero valued float is defined (negative)", func() {
		c, err := functions_for_govaluate.IsDefined(-0.000001)
		Expect(err).To(BeNil())
		Expect(c).To(BeTrue())
	})

	It("should return an param wrong type error", func() {
		_, err := functions_for_govaluate.IsDefined(map[string]interface{}{
			"str": "value",
		})
		Expect(err).NotTo(BeNil())
		Expect(functions_for_govaluate.IsWrongParamType(err)).To(BeTrue())
	})

	It("should return a param count error (too few arguments)", func() {
		_, err := functions_for_govaluate.IsDefined()
		Expect(err).NotTo(BeNil())
		Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
	})

	It("should return a param count error (too many arguments)", func() {
		_, err := functions_for_govaluate.IsDefined(1, 2)
		Expect(err).NotTo(BeNil())
		Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
	})
})
