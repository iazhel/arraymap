package arraymap

type typeMap map[string]interface{}
type typeArrayMap []typeMap

func FindByKey(self *typeArrayMap, key string) (result []interface{}, exs bool) {
	for _, m := range *self {
		if val, prs := m[key]; prs == true {
			result = append(result, val)
			exs = true
		}
	}
	return result, exs
}

func Delete(self *typeArrayMap, key string) {
	for i, m := range *self {
		if _, prs := m[key]; prs == true {
			delete((*self)[i], key)
		}
	}

}

func Insert(self *typeArrayMap, key string, val interface{}) {
	l := len(*self) - 1
	if l == -1 {
		Append(self)
		l = 0
	}
	(*self)[l][key] = val
}

func Append(self *typeArrayMap) {
	*self = append(*self, typeMap{})
}
