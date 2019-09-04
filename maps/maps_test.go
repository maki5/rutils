package maps

import (
	"reflect"
	"testing"
)

func TestCompact(t *testing.T) {
	type testData struct {
		m        Map
		response Map
	}

	examples := map[string]testData{
		"empty map":             testData{m: Map{}, response: Map{}},
		"map with values":       testData{m: Map{"key1": "val1", "key2": "val2"}, response: Map{"key1": "val1", "key2": "val2"}},
		"map with empty values": testData{m: Map{"key1": "val1", "key2": "val2", "key3": nil}, response: Map{"key1": "val1", "key2": "val2"}},
	}

	for k, v := range examples {
		initialMap := v.m
		v.m.Compact()

		if !reflect.DeepEqual(v.m, v.response) {
			t.Errorf("test [%v] failed on method Compact with params(initialMap: %v), expected to be %v got %v",
				k, initialMap, v.response, v.m)
		}
	}
}

func TestKeys(t *testing.T) {
	type testData struct {
		m        Map
		response []interface{}
	}

	examples := map[string]testData{
		"empty map":       testData{m: Map{}, response: []interface{}{}},
		"map with values": testData{m: Map{"key1": "val1", "key2": "val2"}, response: []interface{}{"key1", "key2"}},
	}

	for k, v := range examples {
		initialMap := v.m
		keys := v.m.Keys()

		if !reflect.DeepEqual(*keys, v.response) {
			t.Errorf("test [%v] failed on method Keys with params(initialMap: %v), expected to be %v got %v",
				k, initialMap, v.response, *keys)
		}
	}
}

func TestValues(t *testing.T) {
	type testData struct {
		m        Map
		response []interface{}
	}

	examples := map[string]testData{
		"empty map":       testData{m: Map{}, response: []interface{}{}},
		"map with values": testData{m: Map{"key1": "val1", "key2": "val2"}, response: []interface{}{"val1", "val2"}},
	}

	for k, v := range examples {
		initialMap := v.m
		vals := v.m.Values()

		if !reflect.DeepEqual(*vals, v.response) {
			t.Errorf("test [%v] failed on method Values with params(initialMap: %v), expected to be %v got %v",
				k, initialMap, v.response, *vals)
		}
	}
}

func TestFetchValues(t *testing.T) {
	type testData struct {
		m        Map
		keys     []interface{}
		response []interface{}
	}

	examples := map[string]testData{
		"empty map with key selector": testData{m: Map{}, keys: []interface{}{"key1"}, response: []interface{}{}},
		"empty map":                   testData{m: Map{}, keys: []interface{}{}, response: []interface{}{}},
		"fetch one value":             testData{m: Map{"key1": "val1", "key2": "val2"}, keys: []interface{}{"key1"}, response: []interface{}{"val1"}},
		"fetch multiple values": testData{m: Map{"key1": "val1", "key2": "val2", "key3": "val3"}, keys: []interface{}{"key1", "key2"},
			response: []interface{}{"val1", "val2"}},
	}

	for k, v := range examples {
		initialMap := v.m
		vals := v.m.FetchValues(v.keys)

		if !reflect.DeepEqual(*vals, v.response) {
			t.Errorf("test [%v] failed on method FetchValues with params(initialMap: %v), expected to be %v got %v",
				k, initialMap, v.response, *vals)
		}
	}
}

func TestMerge(t *testing.T) {
	type testData struct {
		m        Map
		otherMap Map
		response Map
	}

	examples := map[string]testData{
		"empty map":     testData{m: Map{}, otherMap: Map{}, response: Map{}},
		"one empty map": testData{m: Map{}, otherMap: Map{"key1": "val1", "key2": "val2"}, response: Map{"key1": "val1", "key2": "val2"}},
		"both map with values": testData{m: Map{"key1": "val1", "key2": "val2"}, otherMap: Map{"key3": "val3"},
			response: Map{"key1": "val1", "key2": "val2", "key3": "val3"}},
		"maps with same values": testData{m: Map{"key1": "val1", "key2": "val2", "key3": "val3"}, otherMap: Map{"key3": "val32", "key4": "val4"},
			response: Map{"key1": "val1", "key2": "val2", "key3": "val32", "key4": "val4"}},
	}

	for k, v := range examples {
		initialMap := v.m
		v.m.Merge(v.otherMap)

		if !v.m.Equal(v.response) {
			t.Errorf("test [%v] failed on method Values with params(initialMap: %v), expected to be %v got %v",
				k, initialMap, v.response, v.m)
		}
	}
}
