package functions_for_govaluate_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/jamillosantos/functions-for-govaluate"
)

var _ = Describe("ifnull function", func() {
	It("should return the first param", func() {
		c, err := functions_for_govaluate.IfNull(1, 2)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(1))
	})

	It("should return the second param", func() {
		c, err := functions_for_govaluate.IfNull(nil, 2)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(2))
	})

	It("should return the thrid param", func() {
		c, err := functions_for_govaluate.IfNull(nil, nil, 3)
		Expect(err).To(BeNil())
		Expect(c).To(Equal(3))
	})

	It("should return nil", func() {
		c, err := functions_for_govaluate.IfNull(nil, nil, nil, nil)
		Expect(err).To(BeNil())
		Expect(c).To(BeNil())
	})

	It("should return a param count error (too few params)", func() {
		c, err := functions_for_govaluate.IfNull()
		Expect(err).NotTo(BeNil())
		Expect(functions_for_govaluate.IsWrongParamsCount(err)).To(BeTrue())
		Expect(c).To(BeNil())
	})
})
