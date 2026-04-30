package chargebee

// Bool returns a pointer to the bool value passed.
func Bool(val bool) *bool {
	return &val
}

// BoolValue returns the value of the bool pointer passed or false if the pointer is nil.
func BoolValue(val *bool) bool {
	if val != nil {
		return *val
	}
	return false
}

// Int32 returns a pointer to the int32 value passed.
func Int32(val int32) *int32 {
	return &val
}

// Int32Value returns the value of the int32 pointer passed or 0 if the pointer is nil.
func Int32Value(val *int32) int32 {
	if val != nil {
		return *val
	}
	return 0
}

// Int64 returns a pointer to the int64 value passed.
func Int64(val int64) *int64 {
	return &val
}

// Int64Value returns the value of the int64 pointer passed or 0 if the pointer is nil.
func Int64Value(val *int64) int64 {
	if val != nil {
		return *val
	}
	return 0
}

// Float64 returns a pointer to the float64 value passed.
func Float64(val float64) *float64 {
	return &val
}

// Float64Value returns the value of the float64 pointer passed or 0 if the pointer is nil.
func Float64Value(val *float64) float64 {
	if val != nil {
		return *val
	}
	return 0
}
