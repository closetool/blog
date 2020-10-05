package collectionsutils

func IsSliceEmpty(slice []interface{}) bool {
	if slice == nil || len(slice) == 0 {
		return false
	}
	return true
}
