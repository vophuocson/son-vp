package workflow

import "sync"

type Command func() error
type Compensation func() error

type Step struct {
	Command        Command
	CompensateFunc Compensation
}

type Activity struct {
	Steps         []*Step
	Compensations []Compensation
}

func (s *Activity) AddStep(step *Step) {
	s.Steps = append(s.Steps, step)
}

func (a *Activity) Execute() error {
	if len(a.Steps) > 0 {
		var wg sync.WaitGroup
		wg.Add(len(a.Steps))
		for _, step := range a.Steps {
			step := step
			go func() {
				defer wg.Done()
				err := step.Command()
				if err != nil {
					return
				}
				a.Compensations = append(a.Compensations, step.CompensateFunc)
			}()
		}
	}
	return nil
}

func (a *Activity) Compensate() error {
	if len(a.Compensations) > 0 {
		var wg sync.WaitGroup
		wg.Add(len(a.Compensations))
		for _, f := range a.Compensations {
			f := f
			go func() {
				defer wg.Done()
				err := f()
				if err != nil {
					if a.Compensations != nil {
						a.Compensations = append(a.Compensations)
					}
					wg.Done()
					return
				}
				wg.Done()
			}()
		}
		return nil
	}
	return nil
}

type WorkflowDefinition struct {
	Steps []*Activity
}
