package html

import (
	"github.com/murlokswarm/app"
)

type node interface {
	app.DOMNode
	SetParent(node)
	Close()
	ConsumeChanges() []Change
}

func attrsEqual(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}

	for k, va := range a {
		if vb, ok := b[k]; !ok || va != vb {
			return false
		}
	}
	return true
}

// Change represents a change to perform in order to render a component within
// a control.
type Change struct {
	// The change type.
	Type string

	// A value that describes how to make the change.
	Value interface{}
}

const (
	createText = "createText"
	setText    = "setText"

	createElem   = "createElem"
	setAttrs     = "setAttrs"
	appendChild  = "appendChild"
	removeChild  = "removeChild"
	replaceChild = "replaceChild"

	setRoot    = "setRoot"
	deleteNode = "deleteNode"
)

type textValue struct {
	ID   string
	Text string `json:",omitempty"`
}

func createTextChange(id string) Change {
	return Change{
		Type: createText,
		Value: textValue{
			ID: id,
		},
	}
}

func setTextChange(id, text string) Change {
	return Change{
		Type: setText,
		Value: textValue{
			ID:   id,
			Text: text,
		},
	}
}

type elemValue struct {
	ID      string
	TagName string            `json:",omitempty"`
	Attrs   map[string]string `json:",omitempty"`
}

type childValue struct {
	ParentID string
	ChildID  string
	OldID    string `json:",omitempty"`
}

func createElemChange(n *elemNode) Change {
	return Change{
		Type: createElem,
		Value: elemValue{
			ID:      n.ID(),
			TagName: n.TagName(),
		},
	}
}

func setAttrsChange(id string, a map[string]string) Change {
	return Change{
		Type: setAttrs,
		Value: elemValue{
			ID:    id,
			Attrs: a,
		},
	}
}

func appendChildChange(parentID, childID string) Change {
	return Change{
		Type: appendChild,
		Value: childValue{
			ParentID: parentID,
			ChildID:  childID,
		},
	}
}

func removeChildChange(parentID, childID string) Change {
	return Change{
		Type: removeChild,
		Value: childValue{
			ParentID: parentID,
			ChildID:  childID,
		},
	}
}

func replaceChildChange(parentID, oldID, newID string) Change {
	return Change{
		Type: replaceChild,
		Value: childValue{
			ParentID: parentID,
			ChildID:  newID,
			OldID:    oldID,
		},
	}
}

func deleteNodeChange(id string) Change {
	return Change{
		Type:  deleteNode,
		Value: id,
	}
}
