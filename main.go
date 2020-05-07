package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	
	data := []float64{
		4,
		4,
		3,
		2,
		3,
		4,
		5,
		2,
		15,
		2,
		3,
		0,
		7,
		4,
		10,
		6,
		7,
		5,
		4,
		2,
		9,
		3,
		6,
		4,
		1,
		1,
		0,
		6,
		5,
		5,
		0,
		2,
		0,
		1,
		7,
		7,
	}

	avg := AverageFloat(data)
	title:= struct {
		Title: "New cases per day in Hays County",
		Heading: fmt.Printf("average daily cases over %d days: %f\n", len(data), avg),
	}
	
	
	//fmt.Printf("moving summaries:\n%#v\n", MovingSummary(data))
	for _, v := range MovingSummary(data) {
		// fmt.Println(v.String())
		if b, err := json.Marshal(v); err == nil {
			fmt.Printf(string(b))
		}
	}

}

// Median finds the median value of a slice of float64
func Median(data []float64) float64 {
	if len(data) <= 2 {
		return AverageFloat(data)
	}
	ds := data
	sort.Float64s(ds)
	mid := len(ds)/2 - 1
	if len(ds)%2 == 1 {
		return data[mid]
	}

	return (ds[mid] + ds[mid+1]) / 2

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// SumFloat finds the sum of a slice of float64
func SumFloat(data []float64) float64 {
	var sum float64
	for _, v := range data {
		sum += v
	}
	return sum
}

// AverageFloat finds the simple average of aslice of float64
func AverageFloat(data []float64) float64 {

	return SumFloat(data) / float64(len(data))
}

// ReverseFloat reverses the values of a slice
func ReverseFloat(a []float64) []float64 {

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a

}

// Summary holds a set of basic statistical properties of a slice of float64
type Summary struct {
	DataIdx    int
	Value      float64
	data       []float64
	Min        float64
	Max        float64
	Mean       float64
	Median     float64
	MovingAvg5 float64
	MovingAvg7 float64
}

// Summary is a Stringer
func (s Summary) String() string {
	return fmt.Sprintf("Index: %v\n\tCases: %v\n\tMean: %v\n\tMedian: %v\n\tM5: %v\n\tM6: %v\n\t",
		s.DataIdx,
		s.Value,
		s.Mean,
		s.Median,
		s.MovingAvg5,
		s.MovingAvg7)

}

// MovingSummary builds a slice of moving summaries from a slice of float64
// the idea is to show how the dataset evolves; especially geared toward time-series}
func MovingSummary(data []float64) []Summary {

	var sumz []Summary

	for i, v := range data {

		tmp := append([]float64{}, data[0:i+1]...)
		sort.Float64s(tmp)
		offset5 := max(i-5, 0)
		offset7 := max(i-7, 0)

		sumz = append(sumz, Summary{
			DataIdx:    i,
			Value:      v,
			data:       data[0 : i+1],
			Min:        tmp[0],
			Max:        tmp[i],
			Mean:       AverageFloat(tmp),
			Median:     Median(tmp),
			MovingAvg5: AverageFloat(data[offset5 : i+1]),
			MovingAvg7: AverageFloat(data[offset7 : i+1]),
		})
	}
	return sumz
}

// calculate moving averages of a set of numbers
// over a given frame; reverse direction by giving a
// negative number for the 'over' interval parameter
func MovingAverage(data []float64, over int) []float64 {

	window := data
	if over == 0 {
		return []float64{0}
	}

	if over < 0 {
		window = ReverseFloat(data)
		over *= -1
	}

	if over >= len(window) {
		avg := AverageFloat(window)
		return []float64{avg}
	}

	//periods are calculated backwards
	//so any truncated period is always
	//first in the list
	sizeOfFirstPeriod := len(window) % over
	fmt.Println(sizeOfFirstPeriod)
	//numPeriods := (len(window) / over)
	averages := []float64{} //make([]float64, numPeriods)
	averages = append(averages, AverageFloat(window[0:sizeOfFirstPeriod]))
	//fmt.Println(window[0:sizeOfFirstPeriod])
	for i := 1; i < len(window)-over; i++ {
		//fmt.Println(window[i : i+over])
		averages = append(averages, AverageFloat(window[i:i+over]))
	}

	return averages

}
