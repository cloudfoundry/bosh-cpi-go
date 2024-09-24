package apiv1_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/cloudfoundry/bosh-cpi-go/apiv1"
)

var _ = Describe("AgentOptions", func() {
	var (
		opts AgentOptions
	)

	Describe("Validate", func() {
		BeforeEach(func() {
			opts = AgentOptions{
				Mbus: "fake-mbus",
				NTP:  []string{},
			}
		})

		It("does not return error if all fields are valid", func() {
			err := opts.Validate()
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns error if Mbus is empty", func() {
			opts.Mbus = ""

			err := opts.Validate()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Must provide non-empty Mbus"))
		})
	})
})
