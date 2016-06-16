package arraymap

import "fmt"
import "testing"
import "math/rand"

type someStruct struct {
	name string
	age  int
}

// Benchmark usual method
func Benchmark_For_Find_5_maps_1(b *testing.B)     { bench_function(5, 1, "For_Find", b) }
func Benchmark_For_Find_500_maps_1(b *testing.B)   { bench_function(500, 1, "For_Find", b) }
func Benchmark_For_Find_50000_maps_1(b *testing.B) { bench_function(50000, 1, "For_Find", b) }

// Benchmark Index
func Benchmark_Index_5_maps_1(b *testing.B)     { bench_function(5, 1, "Index", b) }
func Benchmark_Index_500_maps_1(b *testing.B)   { bench_function(500, 1, "Index", b) }
func Benchmark_Index_50000_maps_1(b *testing.B) { bench_function(50000, 1, "Index", b) }

// Benchmark IndexS
func Benchmark_IndexS_5_maps_1(b *testing.B)     { bench_function(5, 1, "IndexS", b) }
func Benchmark_IndexS_500_maps_1(b *testing.B)   { bench_function(500, 1, "IndexS", b) }
func Benchmark_IndexS_50000_maps_1(b *testing.B) { bench_function(50000, 1, "IndexS", b) }

// Benchmark usual method
func Benchmark_For_Find_1_maps_50(b *testing.B)    { bench_function(1, 5, "For_Find", b) }
func Benchmark_For_Find_1_maps_500(b *testing.B)   { bench_function(1, 500, "For_Find", b) }
func Benchmark_For_Find_1_maps_50000(b *testing.B) { bench_function(1, 50000, "For_Find", b) }

// Benchmark IndexI
func Benchmark_Index_1_maps_5(b *testing.B)     { bench_function(1, 5, "Index", b) }
func Benchmark_Index_1_maps_500(b *testing.B)   { bench_function(1, 500, "Index", b) }
func Benchmark_Index_1_maps_50000(b *testing.B) { bench_function(1, 50000, "Index", b) }

// Benchmark IndexS
func Benchmark_IndexS_1_maps_5(b *testing.B)     { bench_function(1, 5, "IndexS", b) }
func Benchmark_IndexS_1_maps_500(b *testing.B)   { bench_function(1, 500, "IndexS", b) }
func Benchmark_IndexS_1_maps_50000(b *testing.B) { bench_function(1, 50000, "IndexS", b) }

// Benchmark and check results
func bench_function(n, m int, f string, b *testing.B) {
	nm := n * m
	r := rand.New(rand.NewSource(99))
	am, keys := insertData(n, m)
	b.ResetTimer()
	switch f {
	case "Index":
		for i := 0; i < b.N; i++ {
			fullInd := r.Intn(nm)
			sliceIndWant := fullInd % n
			key := keys[fullInd]
			if getN := Index(am, key); getN != sliceIndWant {
				b.Errorf("Index incorrect. GET:%d, WANT:%d ", getN, sliceIndWant)
			}
		}
	case "IndexS":
		for i := 0; i < b.N; i++ {
			fullInd := r.Intn(nm)
			sliceIndWant := fullInd % n
			key := keys[fullInd]
			if getN := IndexS(am, key); getN != sliceIndWant {
				b.Errorf("Index incorrect. GET:%d, WANT:%d ", getN, sliceIndWant)
			}
		}
	case "Value":
		for i := 0; i < b.N; i++ {
			fullInd := r.Intn(nm)
			key := keys[fullInd]
			if getN := Value(am, key); getN != nil {
				b.Errorf("Value incorrect. ")
			}
		}
	case "For_Find":
		for i := 0; i < b.N; i++ {
			fullInd := r.Intn(nm)
			key := keys[fullInd]
			if getN, _ := FindByKey(am, key); getN == nil {
				b.Errorf("Value incorrect. ")
			}
		}
	}
}

func insertData(n, m int) (*typeArrayMap, []string) {
	am := typeArrayMap{}
	keys := []string{}
	c := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			strKey := fmt.Sprintf("Key %06d", c)
			Insert(&am, strKey, &someStruct{
				name: fmt.Sprintf("Name %06d", c),
				age:  i,
			})
			keys = append(keys, strKey)
			c++
		}
		Append(&am)
	}
	return &am, keys
}
