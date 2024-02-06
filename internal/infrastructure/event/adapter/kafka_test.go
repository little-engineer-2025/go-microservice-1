package adapter

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/avisiedo/go-microservice-1/internal/api/event"
	"github.com/avisiedo/go-microservice-1/internal/api/header"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewKafka(t *testing.T) {
	result := NewKafkaHeaders()
	assert.IsType(t, KafkaAdapter{}, result)
}

func TestGetEchoHeader(t *testing.T) {
	type TestCaseGiven struct {
		Ctx     echo.Context
		Key     string
		Default []string
	}
	type TestCase struct {
		Name     string
		Given    TestCaseGiven
		Expected []string
	}
	testCases := []TestCase{
		{
			Name: "Return default value",
			Given: TestCaseGiven{
				Ctx: echo.New().NewContext(
					&http.Request{
						Header: map[string][]string{},
					},
					&echo.Response{},
				),
				Key:     "My-Key",
				Default: []string{"my-default-value"},
			},
			Expected: []string{"my-default-value"},
		},
		{
			Name: "Return the value of the key",
			Given: TestCaseGiven{
				Ctx: echo.New().NewContext(
					&http.Request{
						Header: http.Header{
							"My-Key": {
								"no-my-default-key",
							},
						},
					},
					&echo.Response{},
				),
				Key:     "My-Key",
				Default: []string{"my-default-value"},
			},
			Expected: []string{"no-my-default-key"},
		},
	}

	for _, testCase := range testCases {
		t.Log(testCase.Name)
		result := getEchoHeader(testCase.Given.Ctx, testCase.Given.Key, testCase.Given.Default)
		assert.Equal(t, testCase.Expected, result)
	}
}

func TestFromEchoContextRequestIdHeader(t *testing.T) {
	ctx := echo.New().NewContext(
		&http.Request{
			Header: map[string][]string{
				// Generated with: ./scripts/header.sh 999999
				string(header.HdrRequestID): {"eyJpZGVudGl0eSI6eyJ0eXBlIjoiQXNzb2NpYXRlIiwiYWNjb3VudF9udW1iZXIiOiIiLCJpbnRlcm5hbCI6eyJvcmdfaWQiOiI5OTk5OTkifX19Cg=="},
			},
		},
		&echo.Response{},
	)
	event := event.TopicTodoCreated
	result, err := NewKafkaHeaders().FromEchoContext(ctx, event)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(result))

	// Check that the header request-id was generated
	boolResult := func() bool {
		for _, h := range result {
			if h.Key == string(header.HdrRequestID) {
				if string(h.Value) != "" {
					return true
				}
			}
		}
		return false
	}()
	assert.True(t, boolResult)
}

func TestFromEchoContext(t *testing.T) {
	type TestCaseGiven struct {
		Ctx   echo.Context
		Event string
	}
	type TestCaseExpected struct {
		Headers []kafka.Header
		err     error
	}
	type TestCase struct {
		Name     string
		Given    TestCaseGiven
		Expected TestCaseExpected
	}
	testCases := []TestCase{
		{
			Name: "Error when ctx is nil",
			Given: TestCaseGiven{
				Ctx:   nil,
				Event: "",
			},
			Expected: TestCaseExpected{
				err: fmt.Errorf("ctx cannot be nil"),
			},
		},
		{
			Name: "Error when Event is empty string",
			Given: TestCaseGiven{
				Ctx: echo.New().NewContext(
					&http.Request{
						Header: map[string][]string{},
					},
					&echo.Response{},
				),
				Event: "",
			},
			Expected: TestCaseExpected{
				err: fmt.Errorf("event cannot be an empty string"),
			},
		},
		{
			Name: "Success transformation",
			Given: TestCaseGiven{
				Ctx: echo.New().NewContext(
					&http.Request{
						Header: map[string][]string{
							string(header.HdrRequestID): {"XBlIjoiQXNzb2NpYXRlIiwiYWNjb3V"},
						},
					},
					&echo.Response{},
				),
				Event: event.TopicTodoCreated,
			},
			Expected: TestCaseExpected{
				err: nil,
				Headers: []kafka.Header{
					{
						Key:   string("Topic"),
						Value: []byte(event.TopicTodoCreated),
					},
					{
						Key:   string(header.HdrRequestID),
						Value: []byte("XBlIjoiQXNzb2NpYXRlIiwiYWNjb3V"),
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Log(testCase.Name)
		result, err := NewKafkaHeaders().FromEchoContext(testCase.Given.Ctx, testCase.Given.Event)
		if testCase.Expected.err != nil {
			require.Error(t, err)
			assert.Equal(t, testCase.Expected.err.Error(), err.Error())
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.Expected.Headers, result)
		}
	}
}
