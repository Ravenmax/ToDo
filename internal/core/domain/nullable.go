package domain

/*
Nullable
-fields not provided {Set=false}
-fields provided: value{Set=true, Value != null}
-fileds proveded: null {Set=True, Value = null}
*/
type Nullable[T any] struct {
	Value *T
	Set   bool
}
