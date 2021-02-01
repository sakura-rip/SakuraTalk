package utils

func IsStrInSlice(baseSlice []string, object string) bool {
	for _, con := range baseSlice {
		if con == object {
			return true
		}
	}
	return false
}
