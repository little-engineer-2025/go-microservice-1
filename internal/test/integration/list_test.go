package integration

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	// TODO Add here your test suites
	suite.Run(t, new(SuiteTodosCreate))
	suite.Run(t, new(SuiteTodosRead))
	// suite.Run(t, new(SuiteTodosDelete))
	// suite.Run(t, new(SuiteTodosUpdate))
	// suite.Run(t, new(SuiteTodosList))
}
