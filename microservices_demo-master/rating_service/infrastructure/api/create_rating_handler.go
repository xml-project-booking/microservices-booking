package api

import (
	"term_service/application"
)

type CreateTermCommandHandler struct {
	termService *application.TermService
}

/*func NewCreateTermCommandHandler(userService *application.TermService, publisher saga.Publisher, subscriber saga.Subscriber) (*CreateTermCommandHandler, error) {
	o := &CreateTermCommandHandler{
		termService:       userService,
		replyPublisher:    publisher,
		commandSubscriber: subscriber,
	}
	err := o.commandSubscriber.Subscribe(o.handle)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (handler *CreateTermCommandHandler) handle(command *events.CreateOrderCommand) {
	id, err := primitive.ObjectIDFromHex(command.Order.Id)
	if err != nil {
		return
	}
	order := &domain.Term{Id: id}

	reply := events.CreateOrderReply{Order: command.Order}

	switch command.Type {
	case events.ShipOrder:
		err := handler.termService.Create(order)
		if err != nil {
			reply.Type = events.OrderShippingNotScheduled
			break
		}
		reply.Type = events.OrderShippingScheduled
	default:
		reply.Type = events.UnknownReply
	}

	if reply.Type != events.UnknownReply {
		_ = handler.replyPublisher.Publish(reply)
	}
}*/
