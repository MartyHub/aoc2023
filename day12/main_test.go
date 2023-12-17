package main

import "testing"

func Test_Rec(t *testing.T) {
	tests := []struct {
		row  Row
		want int
	}{
		{
			row:  NewRow("???.### 1,1,3").Expand(),
			want: 1,
		},
		{
			row:  NewRow(".??..??...?##. 1,1,3").Expand(),
			want: 16384,
		},
		{
			row:  NewRow("?#?#?#?#?#?#?#? 1,3,1,6").Expand(),
			want: 1,
		},
		{
			row:  NewRow("????.#...#... 4,1,1").Expand(),
			want: 16,
		},
		{
			row:  NewRow("????.######..#####. 1,6,5").Expand(),
			want: 2500,
		},
		{
			row:  NewRow("?###???????? 3,2,1").Expand(),
			want: 506250,
		},
	}

	for _, tt := range tests {
		t.Run(tt.row.Springs, func(t *testing.T) {
			if got := tt.row.Compute(make(map[string]int)); got != tt.want {
				t.Errorf("Compute1() = %v, want %v for %v", got, tt.want, tt.row)
			}
		})
	}
}
