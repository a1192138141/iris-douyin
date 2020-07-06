package socket

import (
	"ims/lib"
)

type Realize struct {
	*WsMessage
	*Ws
}

func NewRealize(message *WsMessage,  ws *Ws) *Realize {
	return  &Realize{message,ws}
}


func (this *Realize) Test()  {
	this.Ws.Conn.To(this.Ws.Conn.ID()).EmitMessage(lib.SuccessWsResponseMsg("success"))
}