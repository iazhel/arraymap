package arraymap

import "testing"
import "fmt"

// DO not WORK

// Test Value on array with zero size.
func Test_Value_Zero_Slice(t *testing.T) {
	arrayMap := []map[string]*someStruct{}
	if Value(&arrayMap, "key") != nil {
		t.Errorf("Get not nil value from empty slise.")
	}
	if Index(&arrayMap, "key") != -1 {
		t.Errorf("Get not -1 index from empty slise.")
	}
}

// Test empty key using, and view output format.
func Test_Value(b *testing.T) {
	am := typeArrayMap{}
	intMap := []map[int]string{{1: " 11"}, {2: "22"}, {3: "33"}}
	for i := 0; i < 5; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		Insert(&am, strKey, &someStruct{
			name: fmt.Sprintf("Name %06d", i),
			age:  i,
		})
		Append(&am)
	}

	if Value(&am, "") != nil {
		b.Errorf("Get not nil value by empty key.")
	}
	if Value(&am, 33) != nil {
		b.Errorf("Get not nil value by integer.")
	}
	if Value(&am, nil) != nil {
		b.Errorf("Get not nil value by nil key.")
	}
	if Index(&am, "") != -1 {
		b.Errorf("Get not -1 index from empty key.")
	}
	if i := Index(&am, 33); i != -1 {
		b.Errorf("Get not -1 index from by integer key.")
	}
	if Index(&am, nil) != -1 {
		b.Errorf("Get not -1 index from nil key.")
	}
	if Index(&intMap, "") != -1 {
		b.Errorf("Get not -1 index from empty key.")
	}
	if i := Index(&intMap, 3); i != 2 {
		b.Errorf("Get not 2 index from by integer key.", i)
	}
	if Index(&intMap, nil) != -1 {
		b.Errorf("Get not -1 index from nil key.")
	}

	fmt.Println("\nfunc Index() out:")
	checkIndexes := []int{0, 1, 2, 3, 4, -1}
	for i := 0; i < 6; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		if Index(&am, strKey) != checkIndexes[i] {
			b.Errorf("Get not correct index.")
		}
		fmt.Printf("Key '%s'is founded in %d element of slice.\n", strKey, checkIndexes[i])
	}

	fmt.Println("\nfunc Value() out:")
	for i := -1; i < 6; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		val := Value(&am, strKey)
		fmt.Printf("Key '%s' corresponds '%#v'\n", strKey, val)
	}
}
