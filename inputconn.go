package gofbp

type inputCommon interface {
	IsDrained() bool
	resetForNextExecution()
	IsEmpty() bool
	IsClosed() bool
	Close()
	receive(*Process) *Packet
	//Name() string
	//SetName(string)
}

type InputConn interface {
	inputCommon
	//receive(p *Process) *Packet
}

type InputArrayConn interface {
	inputCommon

	GetArrayItem(i int) *InPort
	SetArrayItem(c *InPort, i int)
	ArrayLength() int
}
