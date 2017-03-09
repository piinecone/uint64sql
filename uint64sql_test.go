package uint64sql

import (
	"database/sql/driver"
	"testing"
)

var testTable = map[uint64]string{
	3141592653589793: "3141592653589793",
}

func TestNew(t *testing.T) {
	for f, s := range testTable {
		u := New(f)
		if u.String() != s {
			t.Errorf(
				"expected %s, got %s (%s)",
				s,
				u.String(),
				u.String(),
			)
		}
	}
}

// func TestUint64Sql_Overflow(t *testing.T) {
// 	if !didPanic(func() { New(math.MaxInt32) }) {
// 		t.Fatalf("should have gotten an overflow panic")
// 	}
// }

func TestUint64Sql_Scan(t *testing.T) {
	// test the Scan method that implements the sql.Scanner interface
	a := Uint64Sql{}
	dbvalue := uint64(1234567890)
	expected := New(dbvalue)

	err := a.Scan(dbvalue)
	if err != nil {
		// Scan failed... no need to test result value
		t.Errorf("a.Scan(1234567890) failed with message: %s", err)

	} else {
		// Scan succeeded... test resulting values
		if a.value != expected.value {
			t.Errorf("%s does not equal to %s", a, expected)
		}
	}

	dbInt64 := int64(7241575154197211182)
	expected = New(7241575154197211182)

	err = a.Scan(dbInt64)
	if err != nil {
		// Scan failed... no need to test result value
		t.Errorf("a.Scan(7241575154197211182) failed with message: %s", err)

	} else {
		// Scan succeeded... test resulting values
		if a.value != expected.value {
			t.Errorf("%s does not equal to %s", a, expected)
		}
	}

	dbStr := string("12")
	expected = New(12)

	err = a.Scan(dbStr)
	if err != nil {
		// Scan failed... no need to test result value
		t.Errorf("a.Scan(12) failed with message: %s", err)

	} else {
		// Scan succeeded... test resulting values
		if a.value != expected.value {
			t.Errorf("%s does not equal to %s", a, expected)
		}
	}
}

func TestUint64Sql_Value(t *testing.T) {
	// Make sure this does implement the database/sql's driver.Valuer interface
	var u Uint64Sql
	if _, ok := interface{}(u).(driver.Valuer); !ok {
		t.Error("Uint64Sql does not implement driver.Valuer")
	}

	// check that normal case is handled appropriately
	a := New(1234567890)
	expected := "1234567890"
	value, err := a.Value()
	if err != nil {
		t.Errorf("Uint64Sql(1234567890).Value() failed with message: %s", err)
	} else if value.(string) != expected {
		t.Errorf("%s does not equal to %s", a, expected)
	}
}

func didPanic(f func()) bool {
	ret := false
	func() {

		defer func() {
			if message := recover(); message != nil {
				ret = true
			}
		}()

		// call the target function
		f()

	}()

	return ret
}
