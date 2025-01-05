package port

import (
	order "delivery-food/order/internal/core/domain"
	"delivery-food/order/internal/core/port/workflow"

	"github.com/google/uuid"
)

type OrderRepository interface {
	Create(*order.Order) error
	GetByID(uuid.UUID) (*order.Order, error)
}

type OrderProducer interface {
	VerifyConsumer(*order.Order) error
	CreateTicket(o *order.Order) error
	CompensateTicket(o *order.Order) error
	AuthenticateCard(o *order.Order) error
	ApproveTicketCreation(o *order.Order) error
	ApproveOrderCreation(o *order.Order) error
}

type OrderOrchestratorTransaction interface {
	ExecuteWorkflowCreateOrder(workflowDefinition *workflow.WorkflowDefinition) error
}
