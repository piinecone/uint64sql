// Package uint64sql implements a sql-friendly uint64 for storing high bit set values
//
// Create a new uint64 like this:
//   n, err := uint64sql.New(1234567890)
package uint64sql

import (
	"database/sql/driver"
	"log"
	"strconv"
)

// Uint64Sql is a SQL-safe uint64
type Uint64Sql struct {
	value uint64
}

// New returns a new sql friendly uint64
func New(value uint64) Uint64Sql {
	return Uint64Sql{
		value: value,
	}
}

// String returns the string representation of the uint64
func (u Uint64Sql) String() string {
	return strconv.FormatUint(u.value, 10)
}

// Scan implements the sql.Scanner interface for database deserialization.
func (u *Uint64Sql) Scan(value interface{}) error {
	switch v := value.(type) {
	// if this is an integer type, try to instantiate it as a uint64
	case int, int8, int16, int32, int64:
		vint := strconv.FormatInt(v.(int64), 10)
		vuint64, err := strconv.ParseUint(vint, 10, 64)
		if err != nil {
			*u = New(uint64(vuint64))
			return err
		}
		*u = New(vuint64)
		return nil

	case uint, uint8, uint16, uint32, uint64:
		*u = New(v.(uint64))
		return nil

	// TODO handle sql bigint type
	case []uint, []uint8, []uint16, []uint32, []uint64:
		log.Println(v)
		log.Println("[]uint*")
		*u = New(v.(uint64))
		return nil

	default:
		vu, err := strconv.ParseUint(v.(string), 10, 64)
		*u = New(vu)
		return err
	}
}

// Value implements the driver.Valuer interface for database serialization.
func (u Uint64Sql) Value() (driver.Value, error) {
	return u.String(), nil
}
