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
	"github.com/gofrs/uuid"
)

// Effort represents the effort a agile todo takes to finish
type Effort struct {
	ID          uuid.UUID     `json:"id"`          // An unique ID for the effort
	Date        date.DateTime `json:"date"`        // The date of the effort
	Duration    time.Duration `json:"duration"`    // The work duration
	Description string        `json:"description"` // The work description
}

// NewEffort creates a new effort object
func NewEffort(date date.DateTime, duration time.Duration) *Effort {
	id, _ := uuid.NewV1()
	return &Effort{
		ID:          id,
		Date:        date,
		Duration:    duration,
		Description: "",
	}
}

// String returns a string representation of this effort object
func (e *Effort) String() string {
	return fmt.Sprintf(`{
    id: %d,
    date: %s,
    duration: %s
    description: %s
  }`, e.ID, e.Date.Format(date.Iso8601TZ), e.Duration.String(), e.Description)
}

// Validate checks if this effort is valid
func (e *Effort) Validate() error {
	val := utils.NewValidator()
	val.IsDateDefined(e.Date, "The effort date is not defined")
	val.IsPositive(e.Duration.Nanoseconds(), "The duration must not be zero")
	return val.AllValid()
}
