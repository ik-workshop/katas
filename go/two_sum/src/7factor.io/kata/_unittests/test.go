package _unittests

import (
	"7factor.io/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Two Sum Kata", func() {
	Context("Handles arrays that are too small", func() {
		var genericReturn [2]int
		It("Handles an empty array", func() {
			var emptySlice []int
			result, err := kata.TwoSum(emptySlice, 0)
			Expect(err).ToNot(BeNil())
			Expect(result).To(Equal(genericReturn))
		})

		It("Handles an array of 1 index", func() {
			singleIndexArray := make([]int, 1)
			result, err := kata.TwoSum(singleIndexArray, 0)
			Expect(err).ToNot(BeNil())
			Expect(result).To(Equal(genericReturn))
		})
	})
})
