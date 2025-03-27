package component

var registry = make(map[string]any)

func Register(name string, component any) {
	registry[name] = component
}
