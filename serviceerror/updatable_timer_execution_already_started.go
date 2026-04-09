package serviceerror

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/api/errordetails/v1"
)

type (
	// UpdatableTimerExecutionAlreadyStarted represents an updatable timer execution already started error.
	UpdatableTimerExecutionAlreadyStarted struct {
		Message        string
		StartRequestId string
		RunId          string
		st             *status.Status
	}
)

// NewUpdatableTimerExecutionAlreadyStarted returns new UpdatableTimerExecutionAlreadyStarted error.
func NewUpdatableTimerExecutionAlreadyStarted(message, startRequestId, runId string) error {
	return &UpdatableTimerExecutionAlreadyStarted{
		Message:        message,
		StartRequestId: startRequestId,
		RunId:          runId,
	}
}

// NewUpdatableTimerExecutionAlreadyStartedf returns new UpdatableTimerExecutionAlreadyStarted error with formatted message.
func NewUpdatableTimerExecutionAlreadyStartedf(startRequestId, runId, format string, args ...any) error {
	return &UpdatableTimerExecutionAlreadyStarted{
		Message:        fmt.Sprintf(format, args...),
		StartRequestId: startRequestId,
		RunId:          runId,
	}
}

// Error returns string message.
func (e *UpdatableTimerExecutionAlreadyStarted) Error() string {
	return e.Message
}

func (e *UpdatableTimerExecutionAlreadyStarted) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.AlreadyExists, e.Message)
	st, _ = st.WithDetails(
		&errordetails.UpdatableTimerExecutionAlreadyStartedFailure{
			StartRequestId: e.StartRequestId,
			RunId:          e.RunId,
		},
	)
	return st
}

func newUpdatableTimerExecutionAlreadyStarted(st *status.Status, errDetails *errordetails.UpdatableTimerExecutionAlreadyStartedFailure) error {
	return &UpdatableTimerExecutionAlreadyStarted{
		Message:        st.Message(),
		StartRequestId: errDetails.GetStartRequestId(),
		RunId:          errDetails.GetRunId(),
		st:             st,
	}
}
