package kt_utils

// Transforms anything into a pointer - useful for assigning values to pointer fields os structs / pointer vars
func Ptr[T any](t T) *T { return &t }
