package vm

// VM Main VM type to issue commands against.
type VM struct {
	Name string
	Type string
	ID   string
	IP   string
}

func (vm VM) listVMs() VM {
	myVM := VM{"MyName", "MyType", "MyID", "MyIP"}
	return myVM
}
