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

	date "github.com/bykof/gostradamus"
	"github.com/chordflower/todoman/internal/utils"
)

type Note struct {
	baseModel
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

// NewNote creates a new note with the given name and author
func NewNote(name, author string) *Note {
	return &Note{
		baseModel:   *newBaseModel(),
		Name:        name,
		Description: "",
		Author:      author,
	}
}

// String returns a string representation of this note
func (n *Note) String() string {
	return fmt.Sprintf(`{
    id: %d,
    creation_date: %s,
    name: %s,
    description: %s,
    author: %s
  }`, n.ID, n.CreationDate.Format(date.Iso8601TZ),
		n.Name, n.Description, n.Author)
}

// Validate validates if this note is valid
func (n *Note) Validate() error {
	val := utils.NewValidator()
	val.IsNotEmpty(n.Author, "The author must not be empty")
	val.IsNotEmpty(n.Name, "The name must not be empty")
	return val.AllValid()
}
