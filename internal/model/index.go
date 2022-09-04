package model

import (
	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
	"github.com/gofrs/uuid"
)

// Item represents an index item
type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// NewItem creates a new item
func NewItem(id uuid.UUID, name string) *Item {
	return &Item{
		ID:   id,
		Name: name,
	}
}

// Index represents an item index
type Index struct {
	Items *dll.List `json:"items"`
}

// NewIndex creates a new index
func NewIndex() *Index {
	return &Index{
		Items: dll.New(),
	}
}

// AddItem adds the given item to this index
func (i *Index) AddItem(item *Item) {
	_, found := i.Items.Find(func(index int, value any) bool {
		t, err := value.(*Item)
		if !err && t.ID == item.ID {
			return true
		}
		return false
	})
	if found == nil {
		i.Items.Add(item)
	}
}

// RemoveItem removes the item with the given id from this index
func (i *Index) RemoveItem(id uuid.UUID) {
	position, _ := i.Items.Find(func(index int, value any) bool {
		t, err := value.(*Item)
		if !err && t.ID == id {
			return true
		}
		return false
	})
	if position != -1 {
		i.Items.Remove(position)
	}
}

// HasItem checks if the item with the given id belongs to this index
func (i *Index) HasItem(id uuid.UUID) bool {
	return i.Items.Any(func(index int, value any) bool {
		t, err := value.(*Item)
		if !err && t.ID == id {
			return true
		}
		return false
	})
}
