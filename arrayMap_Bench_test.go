package arraymap

import "fmt"
import "testing"

type someStruct struct {
	name string
	age  int
}

// Find Elements in array map.
func Benchmark_Find_5(b *testing.B)           { benchmark_Find(5, b) }
func Benchmark_Find_Reflect_5(b *testing.B)   { benchmark_Find_Reflect(5, b) }
func Benchmark_Find_50(b *testing.B)          { benchmark_Find(50, b) }
func Benchmark_Find_Reflect_50(b *testing.B)  { benchmark_Find_Reflect(50, b) }
func Benchmark_Find_500(b *testing.B)         { benchmark_Find(500, b) }
func Benchmark_Find_Reflect_500(b *testing.B) { benchmark_Find_Reflect(500, b) }

// Delete Elements in array map.
func Benchmark_Del_5(b *testing.B) { benchmark_Del(5, b) }

//func Benchmark_Del_50(b *testing.B)  { benchmark_Del(50, b) }
//func Benchmark_Del_500(b *testing.B) { benchmark_Del(500, b) }

// Make arrayMap on 5/50/500 elements.
// Each map element is in new array element.
// It should benchmark and check Insert/Find/Delete methods.
//func Benchmark_Insert_Find_Del_Append_5(b *testing.B) { test_IFD(5, "append", b) }
//func Benchmark_Insert_Find_Del_Append_50(b *testing.B)  { test_IFD(50, "append", b) }
//func Benchmark_Insert_Find_Del_Append_500(b *testing.B) { test_IFD(500, "append", b) }

// Make arrayMap on 5/50/500 elements.
// All map elements ares in one array element.
// It should benchmark and check Insert/Find/Delete methods.
//func Benchmark_Insert_Find_Del_Insert_5(b *testing.B) { test_IFD(5, "insert", b) }
//func Benchmark_Insert_Find_Del_Insert_50(b *testing.B)  { test_IFD(50, "insert", b) }
//func Benchmark_Insert_Find_Del_Insert_500(b *testing.B) { test_IFD(500, "insert", b) }
func benchmark_Del(n int, b *testing.B) {
	am := typeArrayMap{}
	for i := 0; i < 1000; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		Insert(&am, strKey, &someStruct{
			name: fmt.Sprintf("Name %06d", i),
			age:  i,
		})
		Append(&am)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		Delete(&am, strKey)
	}
}
func benchmark_Find_Reflect(n int, b *testing.B) {
	am := typeArrayMap{}
	for i := 0; i < 1000; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		Insert(&am, strKey, &someStruct{
			name: fmt.Sprintf("Name %06d", i),
			age:  i,
		})
		Append(&am)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		_ = Value(&am, strKey)
	}
}

func benchmark_Find(n int, b *testing.B) {
	am := typeArrayMap{}
	for i := 0; i < 1000; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		Insert(&am, strKey, &someStruct{
			name: fmt.Sprintf("Name %06d", i),
			age:  i,
		})
		Append(&am)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strKey := fmt.Sprintf("Key %06d", i)
		_, _ = FindByKey(&am, strKey)
	}
}

// IFD = Insert_Find_Delete
func test_IFD(n int, opt string, b *testing.B) {
	for j := 0; j < b.N; j++ {
		am := typeArrayMap{}
		keyFormat := "Target %08d"
		// Insert datas to the array map.
		for i := 0; i < n; i++ {
			target := typeTarget{
				name: "the name",
				ip:   "192.168.0.10",
				port: 199,
			}
			strKey := fmt.Sprintf(keyFormat, i)
			Insert(&am, strKey, &target)
			if opt == "append" {
				Append(&am)
			}
		}
		// Delete one half.
		for i := 0; i <= int(n/2); i++ {
			key := fmt.Sprintf(keyFormat, i)
			Delete(&am, key)
		}
		// Check Find
		for i := 0; i < n; i++ {
			key := fmt.Sprintf(keyFormat, i)
			_, exs := FindByKey(&am, key)
			// Check Delete
			if exs == true && i <= int(n/2) {
				b.Errorf("Find deleted element, key:'%v'.", key)
			}
			// Check Insert
			if exs == false && i > int(n/2) {
				b.Errorf("Can't find inserted element, key:'%v'.", key)
			}
		}
	}
}
