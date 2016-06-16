package arraymap

import "testing"
import "fmt"

// Test func Index(),  IndexS(), IndexI(), Indexes(), Value()
// with array with zero size.
func Test_Value_Zero_Slice(t *testing.T) {
	mapStrSome := []map[string]*someStruct{}
	mapIntSome := []map[int]*someStruct{}
	mapSome := []map[someStruct]*someStruct{}

	// func Index(aMap interface{}, k interface{}) int  test:
	if Index(&mapStrSome, "key") != -1 {
		t.Errorf("Get not -1 Index from empty slise.")
	}
	if Index(&mapIntSome, 1) != -1 {
		t.Errorf("Get not -1 Index from empty slise.")
	}
	if Index(&mapStrSome, true) != -1 {
		t.Errorf("Get not -1 Index from empty slise.")
	}
	if Index(&mapStrSome, 5.5) != -1 {
		t.Errorf("Get not -1 Index from empty slise.")
	}

	// func IndexS(aMap interface{}, k string) int  test:
	if IndexS(&mapStrSome, "key") != -1 {
		t.Errorf("Get not -1 IndexS from empty slise.")
	}
	if IndexS(&mapIntSome, "") != -1 {
		t.Errorf("Get not -1 IndexS from empty slise.")
	}

	// func IndexI(aMap interface{}, k int) int  test:
	if IndexI(&mapStrSome, 5) != -1 {
		t.Errorf("Get not -1 IndexI from empty slise.")
	}
	if IndexI(&mapStrSome, 0) != -1 {
		t.Errorf("Get not -1 IndexI from empty slise.")
	}

	// func Value(aMap interface{}, k interface{}) interface{} test:
	if Value(&mapSome, "key") != nil {
		t.Errorf("Get wrong value.")
	}
	if Value(&mapIntSome, 1) != nil {
		t.Errorf("Get wrong value.")
	}
	if Value(&mapStrSome, true) != nil {
		t.Errorf("Get wrong value.")
	}
	if Value(&mapStrSome, 5.5) != nil {
		t.Errorf("Get wrong value.")
	}

	// func Indexes(aMap interface{}, k interface{}) interface{} test:
	if len(Indexes(&mapSome, "key")) != 0 {
		t.Errorf("Get not emty slice.")
	}
	if len(Indexes(&mapSome, 2)) != 0 {
		t.Errorf("Get not emty slice.")
	}
	if len(Indexes(&mapSome, 2.3)) != 0 {
		t.Errorf("Get not emty slice.")
	}
	if len(Indexes(&mapSome, true)) != 0 {
		t.Errorf("Get not emty slice.")
	}
}

// func Index(),  IndexS(), IndexI(), Indexes(), Value():
func Test_Index(t *testing.T) {
	mapStrStr := []map[string]string{
		{"key1": "value 1",
			"key2": "value 2"},
		{"key3": "value 3",
			"key4": "value 4"},
	}

	mapIntStr := []map[int]string{
		{1: "value 1",
			2: "value 2"},
		{3: "value 3",
			4: "value 4.0"},
		{5: "value 5",
			4: "value 4.1"},
	}

	// func Index(aMap interface{}, k interface{}) int  test:
	if Index(&mapStrStr, "key") != -1 {
		t.Errorf("Get wrong Ingex.")
	}
	if Index(&mapStrStr, "key3") != 1 {
		t.Errorf("Get wrong Ingex.")
	}
	if Index(&mapIntStr, 7) != -1 {
		t.Errorf("Get wrong Ingex.")
	}
	if Index(&mapIntStr, 3) != 1 {
		t.Errorf("Get wrong Ingex.")
	}

	// func IndexS(aMap interface{}, k string) int  test:
	if IndexS(&mapStrStr, "key3") != 1 {
		t.Errorf("Get wrong IngexS.")
	}
	if IndexS(&mapStrStr, "") != -1 {
		t.Errorf("Get wrong IngexS.")
	}

	// func IndexI(aMap interface{}, k int) int  test:
	if IndexI(&mapIntStr, 7) != -1 {
		t.Errorf("Get wrong IngexI.")
	}
	if IndexI(&mapIntStr, 0) != -1 {
		t.Errorf("Get wrong IngexI.")
	}
	if IndexI(&mapIntStr, 5) != 2 {
		t.Errorf("Get wrong IngexI.")
	}

	//func Value(aMap interface{}, k interface{}) interface{} test:
	if Value(&mapIntStr, nil) != nil {
		t.Errorf("Get wrong value.")
	}
	if Value(&mapStrStr, "key") != nil {
		t.Errorf("Get wrong value.")
	}
	if Value(&mapIntStr, 2) == nil {
		t.Errorf("Get wrong value.")
	}
	if Value(&mapIntStr, 2.1) != nil {
		t.Errorf("Get wrong value.")
	}
	if Value(&mapIntStr, false) != nil {
		t.Errorf("Get wrong value.")
	}
	if Value(&mapStrStr, "key2") == nil {
		t.Errorf("Get wrong value.")
	}

	//func Indexes(aMap interface{}, k interface{}) interface{} test:
	if len(Indexes(&mapStrStr, "key")) != 0 {
		t.Errorf("Get wrong ingexes number.")
	}
	if indxs := Indexes(&mapStrStr, "key3"); len(indxs) != 1 {
		t.Errorf("Get wrong ingexes number.")
		if indxs[0] != 1 {
			t.Errorf("Get wrong ingexes.")
		}
	}
	if len(Indexes(&mapStrStr, 2.3)) != 0 {
		t.Errorf("Get wrong ingexes number.")
	}
	if len(Indexes(&mapStrStr, true)) != 0 {
		t.Errorf("Get wrong ingexes number.")
	}
	if len(Indexes(&mapIntStr, 0)) != 0 {
		t.Errorf("Get wrong ingexes number.")
	}
	if len(Indexes(&mapIntStr, 1)) != 1 {
		t.Errorf("Get wrong ingexes number.")
	}
	if indxs := Indexes(&mapIntStr, 4); len(indxs) != 2 {
		t.Errorf("Get wrong ingexes number.")
		if indxs[0] != 1 && indxs[1] != 2 {
			t.Errorf("Get wrong ingexes.")
		}
	}
}

// Test again.
func Test_Value(b *testing.T) {
	am := typeArrayMap{}
	for i := 0; i < 5; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		Insert(&am, strKey, &someStruct{
			name: fmt.Sprintf("Name %06d", i),
			age:  i,
		})
		Append(&am)
	}

	fmt.Println("\nfunc Index() out:")
	checkIndexes := []int{0, 1, 2, 3, 4, -1}
	gotValue := []bool{true, true, true, true, true, false}
	for i := 0; i < 6; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		if Index(&am, strKey) != checkIndexes[i] {
			b.Errorf("Got not correct index.")
		}
		if IndexS(&am, strKey) != checkIndexes[i] {
			b.Errorf("Got not correct index.")
		}
		if (Value(&am, strKey) == nil) == gotValue[i] {
			b.Errorf("Got not correct value.")
		}
	}
}
