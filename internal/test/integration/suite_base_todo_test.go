package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/avisiedo/go-microservice-1/internal/api/header"
	"github.com/avisiedo/go-microservice-1/internal/api/http/public"
	"github.com/google/uuid"
)

// SuiteBaseTodos represents the base Suite to be used for smoke tests, this
// start the services before run the smoke tests.
// TODO the smoke tests cannot be executed in parallel yet, an alternative
// for them would be to use specific http and metrics service in one side,
// and to use a specific OrgID per test by using a generator for it which
// would provide data partition between the tests.
type SuiteBaseTodos struct {
	Suite
}

// SetupTest start the services and await until they are ready
// for being used.
func (s *SuiteBaseTodos) SetupTest() {
	s.Suite.SetupTest()
}

// TearDownTest Stop the services in an ordered way before every
// smoke test executed.
func (s *SuiteBaseTodos) TearDownTest() {
	s.Suite.TearDownTest()
}

func (s *SuiteBaseTodos) TodosCreateWithResponse(domain *public.ToDo) (*http.Response, error) {
	var headers http.Header = http.Header{}

	url := s.DefaultPublicBaseURL() + "/todos"
	headers.Add(header.HdrRequestID, "test_todos_create")
	return s.DoRequest(
		http.MethodPost,
		url,
		headers,
		domain,
	)
}

// RegisterIpaDomain is a helper function to register a domain with the API
// for a rhel-idm domain using the OrgID assigned to the unit test.
// Return the token response or error.
func (s *SuiteBaseTodos) TodosCreate(todo *public.ToDo) (*public.ToDo, error) {
	var (
		resp     *http.Response
		err      error
		data     []byte
		resource *public.ToDo = &public.ToDo{}
	)
	if resp, err = s.TodosCreateWithResponse(todo); err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Expected Status='%d' but got Status='%d'", http.StatusCreated, resp.StatusCode)
	}
	if data, err = io.ReadAll(resp.Body); err != nil {
		return nil, fmt.Errorf("failure when reading body: %w", err)
	}
	defer func() {
		errOld := err
		err = resp.Body.Close()
		if errOld != nil {
			err = errOld
		}
	}()

	if err = json.Unmarshal(data, resource); err != nil {
		return nil, fmt.Errorf("failure when unmarshalling the information: %s", err.Error())
	}
	return resource, nil
}

func (s *SuiteBaseTodos) TodosReadWithResponse(todoID uuid.UUID) (*http.Response, error) {
	headers := http.Header{}
	method := http.MethodGet
	url := s.DefaultPublicBaseURL() + "/todos/" + todoID.String()
	headers.Add(header.HdrRequestID, "test_todos_read")
	return s.DoRequest(
		method,
		url,
		headers,
		nil,
	)
}

// TodoRead is a helper function to read a domain with the API
// for a rhel-idm domain using the OrgID assigned to the unit test.
// Return the token response or error.
func (s *SuiteBaseTodos) TodosRead(todoID uuid.UUID) (*public.ToDo, error) {
	var (
		resp *http.Response
		err  error
		data []byte
	)
	if resp, err = s.TodosReadWithResponse(todoID); err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Expected Status='%d' but got Status='%d'", resp.StatusCode, http.StatusOK)
	}
	if data, err = io.ReadAll(resp.Body); err != nil {
		return nil, fmt.Errorf("failure when reading body: %w", err)
	}
	defer func() {
		errOld := err
		err = resp.Body.Close()
		if errOld != nil {
			err = errOld
		}
	}()

	resource := &public.ToDo{}
	if err = json.Unmarshal(data, resource); err != nil {
		return nil, fmt.Errorf("failure when unmarshalling the information: %w", err)
	}

	return resource, nil
}

func (s *SuiteBaseTodos) TodosDeleteWithResponse(todoID uuid.UUID) (*http.Response, error) {
	headers := http.Header{}
	method := http.MethodDelete
	url := s.DefaultPublicBaseURL() + "/todos/" + todoID.String()
	headers.Add(header.HdrRequestID, "test_todos_delete")
	return s.DoRequest(
		method,
		url,
		headers,
		nil,
	)
}

func (s *SuiteBaseTodos) TodosDelete(todoID uuid.UUID) error {
	var (
		resp *http.Response
		err  error
		data []byte
	)
	if resp, err = s.TodosDeleteWithResponse(todoID); err != nil {
		return err
	}
	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Expected Status='%d' but got Status='%d'", resp.StatusCode, http.StatusNoContent)
	}
	if data, err = io.ReadAll(resp.Body); err != nil {
		return fmt.Errorf("failure when reading body: %w", err)
	}
	defer func() {
		errOld := err
		err = resp.Body.Close()
		if errOld != nil {
			err = errOld
		}
	}()

	if len(data) > 0 {
		return fmt.Errorf("No content was expected")
	}
	return nil
}
