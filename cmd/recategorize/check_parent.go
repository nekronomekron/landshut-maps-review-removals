package main

import (
	"encoding/json"
	"fmt"
	"os"

	"nuernberg-maps-review-removals/internal/mapsreview"
)

func main() {
	// Quick check: read a few raw rows
	data, err := os.ReadFile("output/places.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var raw []struct {
		Name           string `json:"name"`
		Category       string `json:"category"`
		ParentCategory string `json:"parentCategory"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	count := 0
	for _, r := range raw[:10] {
		if r.ParentCategory != "" {
			count++
			fmt.Printf("%s → parentCategory=%s\n", r.Name, r.ParentCategory)
		}
	}
	total := 0
	for _, r := range raw {
		if r.ParentCategory != "" {
			total++
		}
	}
	fmt.Printf("\nRaw JSON: %d / %d have parentCategory\n", total, len(raw))

	// Now read via mapsreview.ReadJSON
	rows, err := mapsreview.ReadJSON("output/places.json", []mapsreview.Place{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	viaRead := 0
	for _, r := range rows {
		if pc := mapsreview.StringValue(r.ParentCategory); pc != "" {
			viaRead++
		}
	}
	fmt.Printf("Via ReadJSON: %d / %d have parentCategory\n", viaRead, len(rows))
}
