package integration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/avisiedo/go-microservice-1/internal/api/header"
	"github.com/avisiedo/go-microservice-1/internal/api/http/public"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// SuiteTodosRead is the suite to validate the smoke test when read domain endpoint at GET /api/todo/v1/todos/:uuid
type SuiteTodosRead struct {
	SuiteBaseOneTodo
}

func (s *SuiteTodosRead) SetupTest() {
	s.SuiteBaseOneTodo.SetupTest()
}

func (s *SuiteTodosRead) TearDownTest() {
	s.SuiteBaseOneTodo.TearDownTest()
}

func (s *SuiteTodosRead) TestReadTodos() {
	t := s.T()
	url := fmt.Sprintf("%s/%s/%s", s.DefaultPublicBaseURL(), "todos", *s.Todos.TodoId)

	// Prepare the tests
	testCases := []TestCase{
		{
			Name: "TestTodosRead",
			Given: TestCaseGiven{
				Method: http.MethodGet,
				URL:    url,
				Header: http.Header{
					header.HdrRequestID: {"test_todos_read"},
				},
			},
			Expected: TestCaseExpect{
				StatusCode: http.StatusOK,
				Header: http.Header{
					header.HdrRequestID: {"test_todos_read"},
				},
				BodyFunc: WrapBodyFuncTodoResponse(t, func(t *testing.T, body *public.ToDo) {
					require.NotNil(t, body)
					assert.Equal(t, s.Todos.TodoId, body.TodoId)
					assert.Equal(t, s.Todos.Title, body.Title)
					assert.Equal(t, s.Todos.Description, body.Description)
					if s.Todos.DueDate != nil {
						require.NotNil(t, body.DueDate)
						assert.Equal(t, *s.Todos.DueDate, *body.DueDate)
					} else {
						assert.Nil(t, body.DueDate)
					}
				}),
			},
		},
	}

	// Execute the test cases
	s.RunTestCases(testCases)
}
