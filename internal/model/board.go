// Copyright 2022 carddamom
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"fmt"
	"image/color"

	date "github.com/bykof/gostradamus"
	"github.com/chordflower/todoman/internal/utils"
	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
	"github.com/gofrs/uuid"
)

// Board represents the model of a board aka container of taks
type Board struct {
	baseModel
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Colour      color.RGBA `json:"colour"`
	Todos       dll.List   `json:"-"`
}

// NewBoard creates a new board with the given values
func NewBoard(name string, colour color.RGBA) (board *Board) {
	board = &Board{
		baseModel: *newBaseModel(),
		Name:      name,
		Colour:    colour,
	}
	return
}

// NewBoard2 creates a new board with the given values
func NewBoard2(name string, colour string) (board *Board) {
	board = &Board{
		baseModel: *newBaseModel(),
		Name:      name,
		Colour:    color.RGBA{},
	}
	fmt.Sscanf(colour, "(%d,%d,%d,%d)", board.Colour.R, board.Colour.G, board.Colour.B, board.Colour.A)
	return
}

// AddTodo adds a new todo to this board
func (b *Board) AddTodo(t *Todo) {
	_, found := b.Todos.Find(func(index int, value any) bool {
		t, err := value.(*Todo)
		if !err && t.ID == t.ID {
			return true
		}
		return false
	})
	if found == nil {
		b.Todos.Add(t)
	}
}

// RemoveTodo removes the todo with the given id from this board
func (b *Board) RemoveTodo(id uuid.UUID) {
	position, _ := b.Todos.Find(func(index int, value any) bool {
		t, err := value.(*Todo)
		if !err && t.ID == id {
			return true
		}
		return false
	})
	if position != -1 {
		b.Todos.Remove(position)
	}
}

// HasTodo checks if this board has a todo with the given id
func (b *Board) HasTodo(id uuid.UUID) bool {
	return b.Todos.Any(func(index int, value any) bool {
		t, err := value.(*Todo)
		if !err && t.ID == id {
			return true
		}
		return false
	})
}

// Validate checks if this board is valid
func (b *Board) Validate() error {
	val := utils.NewValidator()
	val.IsNotEmpty(b.Name, "The board name must not be empty")
	return val.AllValid()
}

// ColourToString returns a string representation of the current board colour
func (b *Board) ColourToString() string {
	return fmt.Sprintf("(%d,%d,%d,%d)", b.Colour.R, b.Colour.G, b.Colour.B, b.Colour.A)
}

// String converts a board to string format
func (b *Board) String() string {
	return fmt.Sprintf(`{
    id: "%s",
    creation_date: "%s",
    name: "%s",
    description: "%s",
    colour: (%d,%d,%d,%d)
  }`, b.ID, b.CreationDate.Format(date.Iso8601TZ), b.Name, b.Description, b.Colour.R, b.Colour.G, b.Colour.B, b.Colour.A)
}
