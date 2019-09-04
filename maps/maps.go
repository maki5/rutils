package maps

// Map type alias for map[interface{}]interface{}
type Map map[interface{}]interface{}

// Compact removes keys with nil values
func (m *Map) Compact() {
	for k, v := range *m {
		if v == nil {
			delete(*m, k)
		}
	}
}

// Keys returns map keys
func (m *Map) Keys() *[]interface{} {
	keys := make([]interface{}, 0, 0)

	for k := range *m {
		keys = append(keys, k)
	}

	return &keys
}

// Values returns map values
func (m *Map) Values() *[]interface{} {
	values := make([]interface{}, 0, 0)

	for _, v := range *m {
		values = append(values, v)
	}

	return &values
}

// FetchValues returns array containig the values asociated with the given keys
func (m *Map) FetchValues(keys []interface{}) *[]interface{} {
	values := make([]interface{}, 0, 0)
	newMap := *m

	for _, k := range keys {
		if newMap[k] != nil {
			values = append(values, newMap[k])
		}
	}

	return &values
}

// Equal compare two maps
func (m *Map) Equal(mapToCompare Map) bool {
	if len(*m) != len(mapToCompare) {
		return false
	}

	newMap := *m

	for k := range *m {
		if newMap[k] != mapToCompare[k] {
			return false
		}
	}
	return true
}

// Merge merges two initial map with given
func (m *Map) Merge(otherMap Map) {
	newMap := *m

	for k, v := range otherMap {
		newMap[k] = v
	}

	*m = newMap
}
