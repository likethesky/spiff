package dynaml

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-incubator/spiff/yaml"
)

var _ = Describe("autos", func() {
	Context("when the path is resource_pools.*.size", func() {
		expr := AutoExpr{[]string{"resource_pools", "some_pool", "size"}}

		It("sums up the instances of the jobs in the pool", func() {
			jobs := []yaml.Node{
				map[string]yaml.Node{
					"name":          "some_job",
					"resource_pool": "some_pool",
					"instances":     3,
				},
				map[string]yaml.Node{
					"name":          "some_other_job",
					"resource_pool": "some_pool",
					"instances":     5,
				},
				map[string]yaml.Node{
					"name":          "some_other_job",
					"resource_pool": "some_other_pool",
					"instances":     5,
				},
			}

			binding := FakeBinding{
				FoundFromRoot: map[string]yaml.Node{
					"jobs": jobs,
				},
			}

			Expect(expr).To(EvaluateAs(8, binding))
		})

		Context("when one of the jobs has non-numeric instances", func() {
			It("returns nil", func() {
				jobs := []yaml.Node{
					map[string]yaml.Node{
						"name":          "some_job",
						"resource_pool": "some_pool",
						"instances":     3,
					},
					map[string]yaml.Node{
						"name":          "some_other_job",
						"resource_pool": "some_pool",
						"instances":     &IntegerExpr{4},
					},
					map[string]yaml.Node{
						"name":          "some_other_job",
						"resource_pool": "some_other_pool",
						"instances":     5,
					},
				}

				binding := FakeBinding{
					FoundFromRoot: map[string]yaml.Node{
						"jobs": jobs,
					},
				}

				Expect(expr).To(FailToEvaluate(binding))
			})
		})
	})
})
