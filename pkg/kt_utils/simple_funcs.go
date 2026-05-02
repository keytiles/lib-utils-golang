package kt_utils

// Transforms anything into a pointer - useful for assigning values to pointer fields or structs / pointer vars.
//
// Deprecated: from Go 1.26 the `new(expr)` built in does exactly the same for expressions and where you already have the value in an existing variable
// `var v T` you can easily use `&v` anyways. So we keep this method for backward compatibility only.
func Ptr[T any](t T) *T { return &t }

// Helps to dereference a pointer safe way. If pointer is Nil the zero value is returned of the original type. Otherwise the dereferenced value.
func ValueFromPtr[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}
