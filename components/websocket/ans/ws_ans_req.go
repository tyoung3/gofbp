package ans

import (
	"github.com/jpaulm/gofbp/core"
)

// WSAnsReq struct contains input and output gofbp.core socket connections
type WSAnsReq struct {
	ipt core.InputConn
	out core.OutputConn
}

// Setup Function defines in and out ports
func (wsansreq *WSAnsReq) Setup(p *core.Process) {
	wsansreq.ipt = p.OpenInPort("IN")
	wsansreq.out = p.OpenOutPortOptional("OUT")
}

// MustRun Function is defined elsewhere
func (WSAnsReq) MustRun() {}

// Execute Function runs a WSAnsReq gofbp component instance process
func (wsansreq *WSAnsReq) Execute(p *core.Process) {

	for {
		pkt := p.Receive(wsansreq.ipt) // open bracket
		if pkt == nil {
			break
		}
		p.Send(wsansreq.out, pkt)

		pkt = p.Receive(wsansreq.ipt) // connection
		p.Send(wsansreq.out, pkt)

		pkt = p.Receive(wsansreq.ipt) //"namelist"
		p.Discard(pkt)

		pkt = p.Create("line1")
		p.Send(wsansreq.out, pkt)
		pkt = p.Create("line2")
		p.Send(wsansreq.out, pkt)
		pkt = p.Create("line3")
		p.Send(wsansreq.out, pkt)

		pkt = p.Receive(wsansreq.ipt) // close bracket
		p.Send(wsansreq.out, pkt)
	}

}
