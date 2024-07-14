package saga

import "context"

type Step interface {
	Invoke(ctx context.Context) error
	Compensate(ctx context.Context) error
}
