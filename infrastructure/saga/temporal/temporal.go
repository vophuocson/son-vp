package temporal

import (
	"context"
	"delivery-food/order/internal/core/port"
	"delivery-food/order/internal/core/port/workflow"

	"go.temporal.io/sdk/client"
	tw "go.temporal.io/sdk/workflow"
)

type temporalClient struct {
	client client.Client
}

type TemporalClient interface {
}

func NewTemporalClient(client client.Client) port.OrderOrchestratorTransaction {
	return &temporalClient{
		client: client,
	}
}

func (t *temporalClient) ExecuteWorkflowCreateOrder(workflowDefinition *workflow.WorkflowDefinition) error {

	options := client.StartWorkflowOptions{
		ID:        "create-order-workflow",
		TaskQueue: "task queue is the name of queue",
	}
	wEx, err := t.client.ExecuteWorkflow(context.Background(), options, t.DefinitionWorkflowCreateOrder, workflowDefinition)
	if err != nil {
		return err
	}
	wEx.Get(context.Background(), nil)
	return nil
}

func (t *temporalClient) DefinitionWorkflowCreateOrder(ctx tw.Context, workflowDefinition *workflow.WorkflowDefinition) error {
	for _, activity := range workflowDefinition.Steps {
		activity := activity
		err := tw.ExecuteActivity(ctx, activity.Execute()).Get(ctx, nil)
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				tw.ExecuteActivity(ctx, activity.Compensate()).Get(ctx, nil)
				// err = multierr.Append(err, errCompensation)
			}
		}()
	}
	return nil
}
