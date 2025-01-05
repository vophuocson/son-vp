package dto

import (
	"gopkg.in/guregu/null.v4"
)

type Message struct {
	Partition null.String
	Topic     null.String
	Key       null.String
	Body      any
	Header    any
}
