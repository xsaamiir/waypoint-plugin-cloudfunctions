package cloudfunctionsutil

import (
	"context"
	"time"

	"google.golang.org/api/cloudfunctions/v1"
)

// WaitForOperation keeps polling long the operation until it finishes either
// successfully or with an error.
func WaitForOperation(
	ctx context.Context,
	service *cloudfunctions.Service,
	op *cloudfunctions.Operation,
) (*cloudfunctions.Operation, error) {
	var err error

	for !op.Done {
		opCall := service.Operations.Get(op.Name)
		opCall = opCall.Context(ctx)
		op, err = opCall.Do()
		if err != nil {
			return nil, err
		}

		time.Sleep(1 * time.Second)
	}

	return op, nil
}
