package integration

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/avisiedo/go-microservice-1/internal/api/header"
	"github.com/avisiedo/go-microservice-1/internal/api/http/public"
	builder_api "github.com/avisiedo/go-microservice-1/internal/test/builder/api/http"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// SuiteTodosCreate is the suite to validate the smoke test when read domain endpoint at POST /api/todo/v1/todos
type SuiteTodosCreate struct {
	SuiteBaseTodos
}

func (s *SuiteTodosCreate) SetupTest() {
	s.SuiteBaseTodos.SetupTest()
}

func (s *SuiteTodosCreate) TearDownTest() {
	s.SuiteBaseTodos.TearDownTest()
}

func (s *SuiteTodosCreate) TestTodosCreate() {
	t := s.T()
	url := fmt.Sprintf("%s/%s", s.DefaultPublicBaseURL(), "todos")
	resource := builder_api.NewToDo().
		WithTitle("Test title").
		WithDescription("Test description").
		Build()

	// Prepare the tests
	testCases := []TestCase{
		{
			Name: "TestTodosCreate",
			Given: TestCaseGiven{
				Method: http.MethodPost,
				URL:    url,
				Header: http.Header{
					header.HdrRequestID: {"test_todos_create"},
				},
				Body: resource,
			},
			Expected: TestCaseExpect{
				StatusCode: http.StatusCreated,
				Header: http.Header{
					header.HdrRequestID: {"test_todos_create"},
				},
				BodyFunc: WrapBodyFuncTodoResponse(t, func(t *testing.T, body *public.ToDo) {
					require.NotNil(t, body)
					assert.Equal(t, resource.Title, body.Title)
					assert.Equal(t, resource.Description, body.Description)
					if resource.DueDate != nil {
						require.NotNil(t, body.DueDate)
						assert.Equal(t, *resource.DueDate, *body.DueDate)
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
