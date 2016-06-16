package arraymap

type typeMap map[string]interface{}
type typeArrayMap []typeMap

//FindByKey is usual find method realisation.
func FindByKey(self *typeArrayMap, key string) (result []interface{}, exs bool) {
	for _, m := range *self {
		if val, prs := m[key]; prs == true {
			result = append(result, val)
			exs = true
		}
	}
	return result, exs
}

// Delete removes first element in arraymap,
// which contains searched key.
func Delete(self *typeArrayMap, key string) {
	for i, m := range *self {
		if _, prs := m[key]; prs == true {
			delete((*self)[i], key)
		}
	}

}

// Insert adds new element into last map in the array.
func Insert(self *typeArrayMap, key string, val interface{}) {
	l := len(*self) - 1
	if l == -1 {
		Append(self)
		l = 0
	}
	(*self)[l][key] = val
}

// Append adds one empty map to the array.
func Append(self *typeArrayMap) {
	*self = append(*self, typeMap{})
}
