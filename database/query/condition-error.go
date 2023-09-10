package query

import (
	"github.com/khulnasoft-lab/gobaseline/database/accessor"
)

type errorCondition struct {
	err error
}

func newErrorCondition(err error) *errorCondition {
	return &errorCondition{
		err: err,
	}
}

func (c *errorCondition) complies(acc accessor.Accessor) bool {
	return false
}

func (c *errorCondition) check() error {
	return c.err
}

func (c *errorCondition) string() string {
	return "[ERROR]"
}
