package hhcm

import (
	"testing"
  "reflect"
)

func TestInsert(t *testing.T) {
	h := NewHHCM(10, 10, 2)

  h.Insert("red")
  if !reflect.DeepEqual(h.heavyHitters, map[string]uint64{"red": 1}) {
    t.Errorf("Insert incorrect")
  }

  h.Insert("red")
  h.Insert("red")
  if !reflect.DeepEqual(h.heavyHitters, map[string]uint64{"red": 3}) {
    t.Errorf("Insert incorrect")
  }

  h.Insert("blue")
  if !reflect.DeepEqual(h.heavyHitters, map[string]uint64{"red": 3, "blue": 1}) {
    t.Errorf("Insert incorrect")
  }

  h.Insert("blue")
  if !reflect.DeepEqual(h.heavyHitters, map[string]uint64{"red": 3, "blue": 2}) {
    t.Errorf("Insert incorrect")
  }

  h.Insert("yellow")
  if !reflect.DeepEqual(h.heavyHitters, map[string]uint64{"red": 3, "blue": 2}) {
    t.Errorf("Insert incorrect")
  }

  h.Insert("yellow")
  if !reflect.DeepEqual(h.heavyHitters, map[string]uint64{"red": 3, "blue": 2}) {
    t.Errorf("Insert incorrect")
  }

  h.Insert("yellow")
  if !reflect.DeepEqual(h.heavyHitters, map[string]uint64{"red": 3, "yellow": 3}) {
    t.Errorf("Insert incorrect")
  }
}

/* Test stub- should create HHCM instances each with various values for
 * heavyHitters map and for numHitters and verify that the Report function
 * for each HHCM prints the top N most frequent paths, in descending order of
 * frequency, where N is numHitters
 *
 */
func Report(t *testing.T) {
}
