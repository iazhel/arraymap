package arraymap

// Package empliments any function with array map types.
// the array map type example:
// var theArrayMap []map[type1]type2

// It have few functions. They find index of array element,
// where contains the map with searched key.

//  i:= arraymap.Index[&theArrayMap, theKey]
//
// IndexS() and IndexI() work only with string and int keys.
// They work fastly than Index().
// But Index() uses all key types.
// func Indexes() return int slice of all founded alements with
// searched key.

import (
	"fmt"
	"reflect"
)

// Finds first array index by any type key.
func Index(aMap interface{}, k interface{}) int {
	if checkTypes := VerifyTypes(aMap, k); checkTypes != "OK" {
		return -1
	}
	return indexes(aMap, k, 0)
}

// IndexS returns index of array element by string key.
func IndexS(aMap interface{}, k string) int {
	key := reflect.ValueOf(k)
	return index(aMap, key)
}

// IndexS returns index of array element by int key.
func IndexI(aMap interface{}, k interface{}) int {
	key := reflect.ValueOf(k)
	if key.Kind() != reflect.Int {
		return -1
	}
	return index(aMap, key)
}

// Finds all arrayMap index by any type key.
func Indexes(aMap interface{}, k interface{}) (positions []int) {
	if checkTypes := VerifyTypes(aMap, k); checkTypes != "OK" {
		return []int{}
	}
	for startInd, i := 0, 0; ; {
		if i = indexes(aMap, k, startInd); i < 0 {
			break
		}
		startInd = i + 1
		positions = append(positions, i)
	}
	return positions
}

// searching next index from start position.
func indexes(aMap interface{}, k interface{}, start int) int {
	key := reflect.ValueOf(k)
	ptr := reflect.ValueOf(aMap)
	slice := ptr.Elem()
	for i := start; i < slice.Len(); i++ {
		theMap := slice.Index(i)
		// look by key
		if val := theMap.MapIndex(key); val.Kind() != reflect.Invalid {
			return i
		}
	}
	return -1
}

// Return Value of element by key. Use reflect package for get need type.
func Value(aMap interface{}, k interface{}) interface{} {
	if checkTypes := VerifyTypes(aMap, k); checkTypes != "OK" {
		//		fmt.Println("Check types:", checkTypes)
		return nil
	}
	key := reflect.ValueOf(k)
	ptr := reflect.ValueOf(aMap)
	slice := ptr.Elem()
	for i := 0; i < slice.Len(); i++ {
		theMap := slice.Index(i)
		// look by key
		if val := theMap.MapIndex(key); val.Kind() != reflect.Invalid {
			return val.Interface()
		}
	}
	return nil
}

// Search first index by key.
func index(aMap interface{}, key reflect.Value) int {
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
						//return val.Interface()
						return i
					}
				}
			}
		}
	}
	return -1
	//return nil
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
	// check both key types.
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
