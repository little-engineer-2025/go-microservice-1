package integration

import (
	"github.com/avisiedo/go-microservice-1/internal/api/http/public"
	builder_api "github.com/avisiedo/go-microservice-1/internal/test/builder/api/http"
	"github.com/avisiedo/go-microservice-1/internal/test/builder/helper"
	"github.com/google/uuid"
)

// SuiteBaseOneTodo is the suite to validate the smoke test that require one registered resource
type SuiteBaseOneTodo struct {
	SuiteBaseTodos
	Todos public.ToDo
}

func (s *SuiteBaseOneTodo) SetupTest() {
	s.SuiteBaseTodos.SetupTest()
	s.Todos = public.ToDo{}
	uuidRandom := &uuid.UUID{}
	*uuidRandom = uuid.New()
	oneTodo, err := s.TodosCreate(
		builder_api.
			NewToDo().
			WithID(uuidRandom).
			WithTitle(helper.GenRandTitle()).
			WithDescription(helper.GenRandDescription()).
			Build(),
	)
	if err != nil {
		s.FailNow("error creating todo", err.Error())
	}
	s.Todos = *oneTodo
}

func (s *SuiteBaseOneTodo) TearDownTest() {
	s.Todos = public.ToDo{}
	s.SuiteBaseTodos.TearDownTest()
}
