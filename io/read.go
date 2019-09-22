/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 22-09-2019
 * |
 * | File Name:     read.go
 * +===============================================
 */

package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"p1/road"
)

// Reader reads the input graph
type Reader interface {
	Read() (map[string][]road.Road, error)
}

// FileReader reads the input graph from the file
type FileReader struct {
	Name string
}

// Read the input graph from the file with given filename
func (fr FileReader) Read() (map[string][]road.Road, error) {
	content, err := ioutil.ReadFile(fr.Name)
	if err != nil {
		return nil, fmt.Errorf("cannot read the input graph (%w)", err)
	}

	var result struct {
		RoadDetails []road.Road `json:"road_details"`
	}

	if err := json.Unmarshal(content, &result); err != nil {
		return nil, fmt.Errorf("cannot read the input graph (%w)", err)
	}

	g := make(map[string][]road.Road)
	for _, road := range result.RoadDetails {
		g[road.Source] = append(g[road.Source], road)
	}

	return g, nil
}
