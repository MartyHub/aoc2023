package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntervals_Split(t *testing.T) {
	tests := []struct {
		name string
		itvs Intervals
		want Intervals
	}{
		{
			name: "in",
			itvs: Intervals{
				newInt(0, 100),
			},
			want: Intervals{
				newInt(0, 24),
				newInt(25, 75),
				newInt(76, 100),
			},
		},
		{
			name: "out",
			itvs: Intervals{
				newInt(0, 10),
			},
			want: Intervals{
				newInt(0, 10),
			},
		},
		{
			name: "left",
			itvs: Intervals{
				newInt(0, 50),
			},
			want: Intervals{
				newInt(0, 24),
				newInt(25, 50),
			},
		},
		{
			name: "right",
			itvs: Intervals{
				newInt(50, 100),
			},
			want: Intervals{
				newInt(50, 75),
				newInt(76, 100),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			itvs := tt.itvs

			for i, pt := range []int{25, 75} {
				itvs = itvs.Split(pt, i%2 == 0)
			}

			if !reflect.DeepEqual(itvs, tt.want) {
				panic(fmt.Sprintf("Expected %v, got %v", tt.want, itvs))
			}
		})
	}
}
