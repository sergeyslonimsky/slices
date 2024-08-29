package slices_test

import (
	"fmt"
	"testing"

	"github.com/sergeyslonimsky/slices"
)

func TestArrayMap(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		arr      []int
		callback func(int, int) string
		want     []string
	}{
		{
			name: "converting int to string with keys",
			arr:  []int{1, 2, 3, 4, 5},
			callback: func(i, v int) string {
				return fmt.Sprintf("%d%d", i, v)
			},
			want: []string{"01", "12", "23", "34", "45"},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := slices.ArrayMap(tt.arr, tt.callback)

			if len(got) != len(tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}

			for i, v := range tt.want {
				if got[i] != v {
					t.Errorf("got %v, want %v", got[i], v)
				}
			}
		})
	}
}
