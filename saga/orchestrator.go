package saga

import "context"

type Orchestrator struct {
	Steps []Step
}

func NewOrchestrator() Orchestrator {
	return Orchestrator{
		Steps: []Step{},
	}
}

func (o *Orchestrator) AddStep(step Step) {
	o.Steps = append(o.Steps, step)
}

func (o *Orchestrator) Run(ctx context.Context) error {
	failedIndex := -1

	for i, step := range o.Steps {
		err := step.Invoke(ctx)
		if err != nil {
			failedIndex = i
			break
		}
	}

	for i := failedIndex; i >= 0; i-- {
		err := o.Steps[i].Compensate(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
