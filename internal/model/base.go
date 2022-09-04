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
	"github.com/gofrs/uuid"
)

// baseModel contains some shared fields for all models
type baseModel struct {
	ID           uuid.UUID     `json:"id"`            // An unique ID for the model
	CreationDate date.DateTime `json:"creation_date"` // The creation date of the model
}

func newBaseModel() *baseModel {
	id, _ := uuid.NewV1()
	return &baseModel{
		ID:           id,
		CreationDate: date.Now(),
	}
}

type mmodel interface {
	fmt.Stringer
	Validate() error
}
