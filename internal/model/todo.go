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
	"time"

	date "github.com/bykof/gostradamus"
	"github.com/chordflower/todoman/internal/utils"
	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
	"github.com/gofrs/uuid"
)

// Todo is the model for a todo/task
type Todo struct {
	baseModel
	Name         string        `json:"name"`          // The name of the todo
	Description  string        `json:"description"`   // A description for the todo
	Status       TodoStatus    `json:"status"`        // The status of the todo
	CompleteDate date.DateTime `json:"complete_date"` // An optional completion date of the todo
	StartDate    date.DateTime `json:"start_date"`    // An optional start date of the todo
	Priority     TodoPriority  `json:"priority"`      // The priority of the todo
	Notes        *dll.List     `json:"notes"`         // The notes that this todo contains
}

// NewTodo creates a new todo with the given name
func NewTodo(name string) *Todo {
	return &Todo{
		baseModel:   *newBaseModel(),
		Name:        name,
		Description: "",
		Status:      STATUS_NEW,
		Priority:    PRIORITY_NORMAL,
		Notes:       dll.New(),
	}
}

// AddNote adds the given note to this todo
func (t *Todo) AddNote(n *Note) {
	_, found := t.Notes.Find(func(index int, value any) bool {
		t, err := value.(*Note)
		if !err && t.ID == n.ID {
			return true
		}
		return false
	})
	if found == nil {
		t.Notes.Add(n)
	}
}

// RemoveNote removes the note with the given id from this todo
func (t *Todo) RemoveNote(id uuid.UUID) {
	position, _ := t.Notes.Find(func(index int, value any) bool {
		t, err := value.(*Note)
		if !err && t.ID == id {
			return true
		}
		return false
	})
	if position != -1 {
		t.Notes.Remove(position)
	}
}

// HasNote checks if the note with the given id belongs to this todo
func (t *Todo) HasNote(id uuid.UUID) bool {
	return t.Notes.Any(func(index int, value any) bool {
		t, err := value.(*Note)
		if !err && t.ID == id {
			return true
		}
		return false
	})
}

// String returns a string representation of the todo
func (t *Todo) String() string {
	return fmt.Sprintf(`{
      id: %d,
      creation_date: "%s",
      name: "%s",
      description: "%s",
      status: %d,
      complete_date: "%s",
      start_date: "%s",
      priority: %d,
      notes: (%s)
  }`, t.ID, t.CreationDate.Format(date.Iso8601TZ),
		t.Name, t.Description, t.Status, t.CompleteDate.Format(date.Iso8601TZ),
		t.StartDate.Format(date.Iso8601TZ), t.Priority, t.Notes)
}

// Validate checks if this task is valid
func (t *Todo) Validate() error {
	val := utils.NewValidator()
	val.IsNotEmpty(t.Name, "The name must not be empty")
	return val.AllValid()
}

// TodoStatus represents the status of a todo
type TodoStatus uint8

const (
	// STATUS_NEW represents the new status
	STATUS_NEW TodoStatus = iota
	// STATUS_STARTED represents the started status
	STATUS_STARTED
	// STATUS_PAUSED represents the paused status
	STATUS_PAUSED
	// STATUS_FINISHED represents the finished status
	STATUS_FINISHED
	// STATUS_DONE represents the done status
	STATUS_DONE
)

// TodoPriority represents the priority of a todo
type TodoPriority uint8

const (
	// PRIORITY_LOWEST represents a todo with the lowest priority
	PRIORITY_LOWEST TodoPriority = iota + 1
	// PRIORITY_LOWER represents a todo with a lower priority
	PRIORITY_LOWER
	// PRIORITY_LOW represents a todo with a low priority
	PRIORITY_LOW
	// PRIORITY_NORMAL represents a todo with normal priority (default)
	PRIORITY_NORMAL
	// PRIORITY_HIGH represents a todo with a high priority
	PRIORITY_HIGH
	// PRIORITY_HIGHER represents a todo with a higher priority
	PRIORITY_HIGHER
	// PRIORITY_HIGHEST represents a todo with the highest priority
	PRIORITY_HIGHEST
)

// AgileTodo represents a todo with some agile related fields
type AgileTodo struct {
	Todo
	Points            uint8         `json:"points"`             // The estimation points of this agile todo
	EstimatedDuration time.Duration `json:"estimated_duration"` // The estimated duration of this agile todo
	Effort            *dll.List     `json:"efforts"`            // The actual effort in duration of this agile todo
}

// NewAgileTodo creates a new agile todo
func NewAgileTodo(name string) *AgileTodo {
	return &AgileTodo{
		Todo:   *NewTodo(name),
		Points: 0,
		Effort: dll.New(),
	}
}

// AddEffort adds the given effort to this agile todo, if the sum of all efforts for a day are not more than 24h.
func (ag *AgileTodo) AddEffort(eff *Effort) bool {
	_, found := ag.Effort.Find(func(index int, value any) bool {
		t, err := value.(*Effort)
		if !err && t.ID == eff.ID {
			return true
		}
		return false
	})
	if found == nil {
		ef := found.(*Effort)

		// All efforts for the same year, month and day added together must not be more than 24 hours...
		t := time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC)
		ag.Effort.Select(findEffortAtSameTime(ef)).Map(func(index int, value any) any {
			return value.(*Effort).Duration
		}).Each(func(index int, value any) {
			t.Add(value.(time.Duration))
		})
		t.Add(eff.Duration)

		// If they aren't than we add the effort to the list
		if !t.After(time.Date(2022, time.January, 2, 0, 0, 0, 0, time.UTC)) {
			ag.Effort.Add(eff)
			return true
		} else { // Else we return false and do jack
			return false
		}
	} else {
		return false
	}
}

// RemoveEffort removes the effort with the given id
func (ag *AgileTodo) RemoveEffort(id uuid.UUID) {
	position, _ := ag.Effort.Find(func(index int, value any) bool {
		t, err := value.(*Effort)
		if !err && t.ID == id {
			return true
		}
		return false
	})
	if position != -1 {
		ag.Effort.Remove(position)
	}
}

// HasEffort checks if the effort with the given id is in this agile todo
func (ag *AgileTodo) HasEffort(id uuid.UUID) bool {
	return ag.Effort.Any(func(index int, value any) bool {
		t, err := value.(*Effort)
		if !err && t.ID == id {
			return true
		}
		return false
	})
}

// GetEffortsFor returns all of the efforts for the given year/month/day
func (ag *AgileTodo) GetEffortsFor(date date.DateTime) (ret []*Effort) {
	dateOnly := date.Copy().CeilDay()
	ret = make([]*Effort, 0)
	ag.Effort.Select(func(index int, value any) bool {
		eff, err := value.(*Effort)
		if !err && eff.Date.Copy().CeilDay().Time().Equal(dateOnly.Time()) {
			return true
		}
		return false
	}).Each(func(index int, value any) {
		ret = append(ret, value.(*Effort))
	})
	return
}

func findEffortAtSameTime(ef *Effort) func(index int, value any) bool {
	return func(index int, value any) bool {
		t, err := value.(*Effort)
		tmp := t.Date.Copy().CeilDay().Time()
		tmp2 := ef.Date.Copy().CeilDay().Time()
		if !err && tmp.Equal(tmp2) {
			return true
		}
		return false
	}
}

// String returns a string representation of this agile todo
func (ag *AgileTodo) String() string {
	return fmt.Sprintf(`{
      id: %d,
      creation_date: "%s",
      name: "%s",
      description: "%s",
      status: %d,
      complete_date: "%s",
      start_date: "%s",
      priority: %d,
      notes: (%s),
      points: %d,
      effort: (%s)
  }`, ag.ID, ag.CreationDate.Format(date.Iso8601TZ),
		ag.Name, ag.Description, ag.Status, ag.CompleteDate.Format(date.Iso8601TZ),
		ag.StartDate.Format(date.Iso8601TZ), ag.Priority, ag.Notes, ag.Points, ag.Effort)
}

// Validate checks if this agile todo is valid
func (ag *AgileTodo) Validate() error {
	val := utils.NewValidator()
	val.IsNotEmpty(ag.Name, "The name must not be empty")
	return val.AllValid()
}
