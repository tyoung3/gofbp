package core

//var _ Conn = (*InArrayPort)(nil)

type OutArrayPort struct {
	network GenNet

	//portName  string
	name  string // full name
	array []*OutPort
	//closed    bool
	connected bool
}

func (o *OutArrayPort) send(p *Process, pkt *Packet) bool { panic("send on array port") }

func (o *OutArrayPort) GetArrayItem(i int) *OutPort {
	if i >= len(o.array) {
		return nil
	}
	return o.array[i]
}

func (o *OutArrayPort) SetArrayItem(o2 *OutPort, i int) {
	if i >= len(o.array) {
		// add to .array to fit c2
		increaseBy := make([]*OutPort, i-len(o.array)+1)
		o.array = append(o.array, increaseBy...)
	}
	o.array[i] = o2
}

func (o *OutArrayPort) ArrayLength() int {
	return len(o.array)
}

func (o *OutArrayPort) Close() {
	for _, v := range o.array {
		v.Close()
	}
}

func (o *OutArrayPort) IsConnected() bool {
	//return o.connected
	return true
}
