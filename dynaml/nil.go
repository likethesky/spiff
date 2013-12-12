package dynaml

import (
	"github.com/cloudfoundry-incubator/spiff/yaml"
)

type NilExpr struct{}

func (e NilExpr) Evaluate(Binding) (yaml.Node, bool) {
	return nil, true
}
