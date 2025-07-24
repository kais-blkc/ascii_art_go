package event

import (
	"sync"

	"github.com/google/uuid"
)

type EventData map[string]any

type EventHandler func(EventData)

type handerWrapper struct {
	id      string
	handler EventHandler
}

type EventListener struct {
	mu       sync.RWMutex
	handlers map[string]map[string]handerWrapper
}

func NewEventListener() *EventListener {
	return &EventListener{
		handlers: make(map[string]map[string]handerWrapper),
	}
}

func (el *EventListener) Subscribe(event string, handler EventHandler) string {
	el.mu.Lock()
	defer el.mu.Unlock()

	id := uuid.New().String()
	wrapper := handerWrapper{
		id:      id,
		handler: handler,
	}

	_, ok := el.handlers[event]
	if !ok {
		el.handlers[event] = make(map[string]handerWrapper)
	}

	el.handlers[event][id] = wrapper
	return id
}

func (el *EventListener) Unsubscribe(event string, id string) bool {
	el.mu.Lock()
	defer el.mu.Unlock()

	_, ok := el.handlers[event]
	if ok {
		_, handlerExists := el.handlers[event][id]
		if handlerExists {
			delete(el.handlers[event], id)
			return true
		}
	}

	return false
}

func (el *EventListener) UnsubscribeAll(event string) {
	el.mu.Lock()
	defer el.mu.Unlock()

	delete(el.handlers, event)
}

func (el *EventListener) Emit(event string, data EventData) {
	el.mu.RLock()
	defer el.mu.RUnlock()

	handlers, ok := el.handlers[event]
	if ok {
		for _, wrapper := range handlers {
			go wrapper.handler(data)
		}
	}
}
