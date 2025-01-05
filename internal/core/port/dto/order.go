package dto

import "github.com/google/uuid"

type ConfirmCreateOrder struct {
	OrderID           uuid.UUID
	ChannelNamesReply map[string]bool
}

type ReplyOrderCreation struct {
	OrderID          uuid.UUID
	ServiceNameReply string
	Success          bool
	NativeError      int
}
