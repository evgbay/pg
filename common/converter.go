package common

func ToValue[T any](t *T) T {
	if t == nil {
		var def_ T
		return def_
	}
	return *t
}
