package mapsreview

import (
	"sync"
)

// Source: Stadt Nürnberg Bezirksatlas InstantAtlas layer
// https://online-service2.nuernberg.de/geoinf/ia_bezirksatlas/
//

type Bezirk struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BezirkBoundary struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Label    string        `json:"label"`
	Polygons [][][]float64 `json:"polygons"`
}

type bezirkMapSource struct {
	PixelWidth  float64               `json:"pixelWidth"`
	PixelHeight float64               `json:"pixelHeight"`
	BoundingBox string                `json:"boundingBox"`
	Features    []bezirkFeatureSource `json:"features"`
}

type bezirkFeatureSource struct {
	ID    string      `json:"d"`
	Name  string      `json:"n"`
	Paths [][]float64 `json:"p"`
}

type bezirkPolygon struct {
	Bezirk
	Rings [][]point
	MinX  float64
	MinY  float64
	MaxX  float64
	MaxY  float64
}

type bezirkIndex struct {
	MinX        float64
	MinY        float64
	MaxX        float64
	MaxY        float64
	PixelWidth  float64
	PixelHeight float64
	Polygons    []bezirkPolygon
}

type point struct {
	X float64
	Y float64
}

var (
	bezirkOnce  sync.Once
	bezirke     *bezirkIndex
	bezirkError error
)

func AssignBezirk(lat, lng float64) *Bezirk {
	return assignBezirk(lat, lng, "", false)
}

func AssignBezirkForPostcode(lat, lng float64, postcode string) *Bezirk {
	return assignBezirk(lat, lng, postcode, true)
}

func AllBezirke() []Bezirk {
	return make([]Bezirk, 0)
}

func BezirkBoundaries() []BezirkBoundary {
	return make([]BezirkBoundary, 0)
}

func assignBezirk(lat, lng float64, postcode string, allowFallback bool) *Bezirk {
	return nil
}