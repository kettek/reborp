package component

// Name is a component that holds a name.
type Name struct {
	name string
}

// Name returns the name.
func (n *Name) Name() string {
	return n.name
}

// SetName sets the name.
func (n *Name) SetName(name string) {
	n.name = name
}
