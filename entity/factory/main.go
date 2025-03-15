package factory

import (
	"reflect"

	"github.com/kettek/reborp/acts"
)

type Action = acts.Action

// Entity is an interface that defines an entity that can be updated.
type Entity interface {
	Update() []Action
	Components() []any
	SetComponent(comp any)
}

var registry = make(map[string]func(...any) Entity)

// RegisterEntityFunc registers a factory function for a given entity name.
func RegisterEntityFunc(name string, factory func(...any) Entity) {
	registry[name] = factory
}

// RegisterEntity registers an entity type for a given name and uses reflection to create a new instance.
func RegisterEntity(name string, entity Entity) {
	elem := reflect.TypeOf(entity).Elem()
	registry[name] = func(...any) Entity {
		return reflect.New(elem).Interface().(Entity)
	}
}

// CreateEntity creates a new entity by name. Returns nil if the entity does not exist.
func CreateEntity(name string, components ...any) Entity {
	if factory, ok := registry[name]; ok {
		return factory(components...)
	}
	return nil
}
