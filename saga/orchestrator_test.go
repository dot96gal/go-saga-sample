package saga

import (
	"context"
	"errors"
	"testing"

	"github.com/dot96gal/go-saga-sample/mock"
	"go.uber.org/mock/gomock"
)

func TestOrchestrator_Run(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	testcases := []struct {
		name     string
		steps    []Step
		expected error
	}{
		{
			name:     "no step",
			steps:    []Step{},
			expected: nil,
		},
		{
			name: "success 1st invoke, 1st compensate never called",
			steps: func() []Step {
				s1 := mock.NewMockStep(ctrl)
				gomock.InOrder(
					s1.EXPECT().Invoke(ctx).Return(nil),
				)
				// 1st/2nd compensate never called
				s1.EXPECT().Compensate(ctx).Return(nil).Times(0)

				steps := []Step{}
				steps = append(steps, s1)

				return steps
			}(),
			expected: nil,
		},
		{
			name: "failure 1st invoke, success 1st compensate",
			steps: func() []Step {
				s1 := mock.NewMockStep(ctrl)
				gomock.InOrder(
					s1.EXPECT().Invoke(ctx).Return(errors.New("invoke failure")),
					s1.EXPECT().Compensate(ctx).Return(nil),
				)

				steps := []Step{}
				steps = append(steps, s1)

				return steps
			}(),
			expected: nil,
		},
		{
			name: "success 1st/2nd invoke, 2nd/1st compensate never called",
			steps: func() []Step {
				s1 := mock.NewMockStep(ctrl)
				s2 := mock.NewMockStep(ctrl)
				gomock.InOrder(
					s1.EXPECT().Invoke(ctx).Return(nil),
					s2.EXPECT().Invoke(ctx).Return(nil),
				)
				// 2nd/1st compensate never called
				s2.EXPECT().Compensate(ctx).Return(nil).Times(0)
				s1.EXPECT().Compensate(ctx).Return(nil).Times(0)

				steps := []Step{}
				steps = append(steps, s1)
				steps = append(steps, s2)

				return steps
			}(),
			expected: nil,
		},
		{
			name: "failure 1st invoke, success 1st compensate, 2nd step never called",
			steps: func() []Step {
				s1 := mock.NewMockStep(ctrl)
				s2 := mock.NewMockStep(ctrl)
				gomock.InOrder(
					s1.EXPECT().Invoke(ctx).Return(errors.New("invoke failure")),
					s1.EXPECT().Compensate(ctx).Return(nil),
				)
				// 2nd step never called
				s2.EXPECT().Invoke(ctx).Return(nil).Times(0)
				s2.EXPECT().Compensate(ctx).Return(nil).Times(0)

				steps := []Step{}
				steps = append(steps, s1)
				steps = append(steps, s2)

				return steps
			}(),
			expected: nil,
		},
		{
			name: "success 1st invoke, failure 2nd invoke, success 2nd/1st compensate",
			steps: func() []Step {
				s1 := mock.NewMockStep(ctrl)
				s2 := mock.NewMockStep(ctrl)
				gomock.InOrder(
					s1.EXPECT().Invoke(ctx).Return(nil),
					s2.EXPECT().Invoke(ctx).Return(errors.New("invoke failure")),
					s2.EXPECT().Compensate(ctx).Return(nil),
					s1.EXPECT().Compensate(ctx).Return(nil),
				)

				steps := []Step{}
				steps = append(steps, s1)
				steps = append(steps, s2)

				return steps
			}(),
			expected: nil,
		},
		{
			name: "success 1st invoke, success 2nd invoke, success 3rd invoke, 3rd/2nd/1st compensate never called",
			steps: func() []Step {
				s1 := mock.NewMockStep(ctrl)
				s2 := mock.NewMockStep(ctrl)
				s3 := mock.NewMockStep(ctrl)
				gomock.InOrder(
					s1.EXPECT().Invoke(ctx).Return(nil),
					s2.EXPECT().Invoke(ctx).Return(nil),
					s3.EXPECT().Invoke(ctx).Return(nil),
				)
				// 3rd/2nd/1st compensate never called
				s3.EXPECT().Compensate(ctx).Return(nil).Times(0)
				s2.EXPECT().Compensate(ctx).Return(nil).Times(0)
				s1.EXPECT().Compensate(ctx).Return(nil).Times(0)

				steps := []Step{}
				steps = append(steps, s1)
				steps = append(steps, s2)
				steps = append(steps, s3)

				return steps
			}(),
			expected: nil,
		},
		{
			name: "failure 1st invoke, success 1st compensate, 3rd/2nd step never called",
			steps: func() []Step {
				s1 := mock.NewMockStep(ctrl)
				s2 := mock.NewMockStep(ctrl)
				s3 := mock.NewMockStep(ctrl)
				gomock.InOrder(
					s1.EXPECT().Invoke(ctx).Return(errors.New("invoke failure")),
					s1.EXPECT().Compensate(ctx).Return(nil),
				)
				// 3rd/2nd step never called
				s3.EXPECT().Invoke(ctx).Return(nil).Times(0)
				s3.EXPECT().Compensate(ctx).Return(nil).Times(0)
				s2.EXPECT().Invoke(ctx).Return(nil).Times(0)
				s2.EXPECT().Compensate(ctx).Return(nil).Times(0)

				steps := []Step{}
				steps = append(steps, s1)
				steps = append(steps, s2)
				steps = append(steps, s3)

				return steps
			}(),
			expected: nil,
		},
		{
			name: "success 1st invoke, failure 2nd invoke, success 2nd/1st compensate, 3rd step never called",
			steps: func() []Step {
				s1 := mock.NewMockStep(ctrl)
				s2 := mock.NewMockStep(ctrl)
				s3 := mock.NewMockStep(ctrl)
				gomock.InOrder(
					s1.EXPECT().Invoke(ctx).Return(nil),
					s2.EXPECT().Invoke(ctx).Return(errors.New("invoke failure")),
					s2.EXPECT().Compensate(ctx).Return(nil),
					s1.EXPECT().Compensate(ctx).Return(nil),
				)
				// 3rd step never called
				s3.EXPECT().Invoke(ctx).Return(nil).Times(0)
				s3.EXPECT().Compensate(ctx).Return(nil).Times(0)

				steps := []Step{}
				steps = append(steps, s1)
				steps = append(steps, s2)
				steps = append(steps, s3)

				return steps
			}(),
			expected: nil,
		},
		{
			name: "success 1st/2nd invoke, failure 3rd invoke, success 3rd/2nd/1st compensate",
			steps: func() []Step {
				s1 := mock.NewMockStep(ctrl)
				s2 := mock.NewMockStep(ctrl)
				s3 := mock.NewMockStep(ctrl)
				gomock.InOrder(
					s1.EXPECT().Invoke(ctx).Return(nil),
					s2.EXPECT().Invoke(ctx).Return(nil),
					s3.EXPECT().Invoke(ctx).Return(errors.New("invoke failure")),
					s3.EXPECT().Compensate(ctx).Return(nil),
					s2.EXPECT().Compensate(ctx).Return(nil),
					s1.EXPECT().Compensate(ctx).Return(nil),
				)

				steps := []Step{}
				steps = append(steps, s1)
				steps = append(steps, s2)
				steps = append(steps, s3)

				return steps
			}(),
			expected: nil,
		},
	}

	for _, tt := range testcases {
		orchestrator := NewOrchestrator()
		for _, step := range tt.steps {
			orchestrator.AddStep(step)
		}

		actual := orchestrator.Run(ctx)
		if !errors.Is(actual, tt.expected) {
			t.Errorf("got=%v, want=%v", actual, tt.expected)
		}
	}
}
