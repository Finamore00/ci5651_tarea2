package main

import (
	"ci5651_tarea2/src/datatypes"
	"fmt"
	"sort"
)

func optimizeTunnel(billboards []*datatypes.Interval) []*datatypes.Interval {
	//Accepted set
	included := []*datatypes.Interval{}
	//Sort slices by end point
	sort.Slice(billboards, func(i int, j int) bool {
		return billboards[i].End() < billboards[j].End()
	})

	included = append(included, billboards[0])
	billboards = billboards[1:]

	for _, e := range billboards {
		//Include next billboard with the lowest end point if it doesn't intersect with the previous included billboard
		if !e.HasIntersectionWith(included[len(included)-1]) {
			included = append(included, e)
		}
	}

	return included
}

func main() {
	billboards := []*datatypes.Interval{
		{
			Start:  1,
			Length: 4,
		},
		{
			Start:  2,
			Length: 1,
		},
		{
			Start:  3,
			Length: 3,
		},
	}

	for _, e := range optimizeTunnel(billboards) {
		fmt.Println(e.Start)
	}
}
