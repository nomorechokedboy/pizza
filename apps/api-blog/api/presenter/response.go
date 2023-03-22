package presenter

import ()

type ResponseMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
