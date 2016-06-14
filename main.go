package arraymap

import "fmt"
import "reflect"

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

func VerifyTypes(aMap interface{}, k interface{}) string {
	// verify on pointer
	ptr := reflect.ValueOf(aMap)
	if ptr.Kind() != reflect.Ptr {
		return "Argument is not pointer."
	}
	//verify on slice
	slice := ptr.Elem()
	if slice.Kind() != reflect.Slice {
		return "Argument is not slice."
	}
	//verify slice size
	if slice.Len() == 0 {
		return "Slise size is zero."
	}
	//verify on map
	theMap := slice.Index(0)
	if theMap.Kind() != reflect.Map {
		return "Argument is not map."
	}

	// check map key types.
	for j := 0; j < slice.Len(); j++ {
		keys := theMap.MapKeys()
		if len(keys) > 0 {
			key := keys[0]
			if t1, t2 := key.Kind(), reflect.ValueOf(k).Kind(); t1 != t2 {
				return fmt.Sprintf("Keys types are different, find by key '%v', have: '%v'.", t1, t2)
			} else {
				return "OK"
			}
		}
	}
	return "Maps are empty."
}

// Find arrayMap element by key. It uses reflect package.
func Value(aMap interface{}, k interface{}) interface{} {
	// Check types.
	if checkTypes := VerifyTypes(aMap, k); checkTypes != "OK" {
		fmt.Println("Check types:", checkTypes)
		return nil
	}
	key := reflect.ValueOf(k)
	ptr := reflect.ValueOf(aMap)
	slice := ptr.Elem()
	for i := 0; i < slice.Len(); i++ {
		theMap := slice.Index(i)
		//verify on key presents
		if val := theMap.MapIndex(key); val.Kind() != reflect.Invalid {
			return val.Elem().Interface()
		}
	}
	return nil
}

// Find first arrayMap index by map key. It uses reflect package.
func Index(aMap interface{}, k interface{}) int {
	if checkTypes := VerifyTypes(aMap, k); checkTypes != "OK" {
		fmt.Println("Check types:", checkTypes)
		return -1
	}
	key := reflect.ValueOf(k)
	ptr := reflect.ValueOf(aMap)
	slice := ptr.Elem()
	for i := 0; i < slice.Len(); i++ {
		theMap := slice.Index(i)
		//verify on key presents
		if val := theMap.MapIndex(key); val.Kind() != reflect.Invalid {
			return i
		}
	}
	return -1
}

// Index with if in every step.
func Value_ifs(aMap interface{}, k interface{}) interface{} {
	fmt.Println("Verify Types:", VerifyTypes(aMap, k))
	key := reflect.ValueOf(k)
	// verify on pointer
	if ptr := reflect.ValueOf(aMap); ptr.Kind() == reflect.Ptr {
		//verify on slice
		if slice := ptr.Elem(); slice.Kind() == reflect.Slice {
			for i := 0; i < slice.Len(); i++ {
				//verify on map
				if theMap := slice.Index(i); theMap.Kind() == reflect.Map {
					//verify on key presents
					if val := theMap.MapIndex(key); val.Kind() == reflect.Invalid {
						continue
					} else {
						return val.Interface()
					}
				}
			}
		}
	}
	return nil
}

type typeTarget struct {
	name string
	ip   string
	port int
}

func main() {

	am := typeArrayMap{}
	n := 1000
	Append(&am)
	for i := 0; i < n; i++ {
		target := typeTarget{
			name: "the name",
			ip:   "192.168.0.10",
			port: 199,
		}
		strK := fmt.Sprintf("Target %06d", i)
		Insert(&am, strK, &target)
		Append(&am)
	}
	target, exs := FindByKey(&am, "Target 000123")
	fmt.Println(exs)
	fmt.Printf("%#v", target)
	fmt.Println(target)
}
