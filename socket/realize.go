package socket

import (
	"fmt"
)

type Realize struct {
	*WsMessage
	*Ws
}

func NewRealize(message *WsMessage,  ws *Ws) *Realize {
	return  &Realize{message,ws}
}


func (this *Realize) Test()  {
	fmt.Print(this)
}