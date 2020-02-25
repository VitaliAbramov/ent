// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/pet"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Buffer holds the value of the "buffer" field.
	Buffer []byte `json:"buffer,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// NewName holds the value of the "new_name" field.
	NewName string `json:"new_name,omitempty"`
	// Blob holds the value of the "blob" field.
	Blob []byte `json:"blob,omitempty"`
	// State holds the value of the "state" field.
	State user.State `json:"state,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges     UserEdges `json:"edges"`
	user_pets *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Car holds the value of the car edge.
	Car []*Car
	// Pets holds the value of the pets edge.
	Pets *Pet
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// CarOrErr returns the Car value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CarOrErr() ([]*Car, error) {
	if e.loadedTypes[0] {
		return e.Car, nil
	}
	return nil, &NotLoadedError{edge: "car"}
}

// PetsOrErr returns the Pets value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) PetsOrErr() (*Pet, error) {
	if e.loadedTypes[1] {
		if e.Pets == nil {
			// The edge pets was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pet.Label}
		}
		return e.Pets, nil
	}
	return nil, &NotLoadedError{edge: "pets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // age
		&sql.NullString{}, // name
		&sql.NullString{}, // nickname
		&sql.NullString{}, // phone
		&[]byte{},         // buffer
		&sql.NullString{}, // title
		&sql.NullString{}, // new_name
		&[]byte{},         // blob
		&sql.NullString{}, // state
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // user_pets
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[0])
	} else if value.Valid {
		u.Age = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field nickname", values[2])
	} else if value.Valid {
		u.Nickname = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field phone", values[3])
	} else if value.Valid {
		u.Phone = value.String
	}
	if value, ok := values[4].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field buffer", values[4])
	} else if value != nil {
		u.Buffer = *value
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field title", values[5])
	} else if value.Valid {
		u.Title = value.String
	}
	if value, ok := values[6].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field new_name", values[6])
	} else if value.Valid {
		u.NewName = value.String
	}
	if value, ok := values[7].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field blob", values[7])
	} else if value != nil {
		u.Blob = *value
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[8])
	} else if value.Valid {
		u.State = user.State(value.String)
	}
	values = values[9:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field user_pets", value)
		} else if value.Valid {
			u.user_pets = new(int)
			*u.user_pets = int(value.Int64)
		}
	}
	return nil
}

// QueryCar queries the car edge of the User.
func (u *User) QueryCar() *CarQuery {
	return (&UserClient{config: u.config}).QueryCar(u)
}

// QueryPets queries the pets edge of the User.
func (u *User) QueryPets() *PetQuery {
	return (&UserClient{config: u.config}).QueryPets(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("entv2: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", buffer=")
	builder.WriteString(fmt.Sprintf("%v", u.Buffer))
	builder.WriteString(", title=")
	builder.WriteString(u.Title)
	builder.WriteString(", new_name=")
	builder.WriteString(u.NewName)
	builder.WriteString(", blob=")
	builder.WriteString(fmt.Sprintf("%v", u.Blob))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", u.State))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
