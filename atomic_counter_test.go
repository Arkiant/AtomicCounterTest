package benchuuid

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {

	tc := []struct {
		Name     string
		Quantity int
	}{
		{"100 items", 100},
		{"50 items", 50},
		{"1 items", 1},
	}

	for _, tt := range tc {
		t.Run(tt.Name, func(t *testing.T) {
			counter := New()
			var wg sync.WaitGroup
			wg.Add(tt.Quantity)
			for i := 0; i < tt.Quantity; i++ {
				go func(val *Counter) {
					val.Increment()
					time.Sleep(1 * time.Second)
					wg.Done()
				}(&counter)
			}

			wg.Wait()

			if counter.Get() != uint32(tt.Quantity) {
				t.Fatalf("Counter don't work correctly, expected value %d given value %d", tt.Quantity, counter)
			}
		})
	}

}

func BenchmarkCounter(b *testing.B) {
	b.ReportAllocs()
	counter := New()
	for i := 0; i < b.N; i++ {
		counter.Increment()
	}

}
