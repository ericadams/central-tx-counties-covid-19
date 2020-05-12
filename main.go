package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//TODO : update data files, Load data from files
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
		12,
		4,
		3,
		1,
	}

	avg := AverageFloat(data)
	title := struct {
		Title     string
		Heading   string
		Summaries []Summary
	}{
		"New cases per day in Hays County",
		fmt.Sprintf("average daily cases over %d days: %f\n", len(data), avg),
		[]Summary{},
	}

	//fmt.Printf("moving summaries:\n%#v\n", MovingSummary(data))
	for _, v := range MovingSummary(data) {
		title.Summaries = append(title.Summaries, v)
	}

	b, err := json.Marshal(title)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
