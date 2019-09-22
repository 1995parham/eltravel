/*
 * In The Name of God
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 22-09-2019
 * |
 * | File Name:     bfs.go
 * +===============================================
 */

package alg

import (
	"p1/road"
)

type node struct {
	name  string
	roads []road.Road
}

// Route find a route from src to dst in the given graph g
func Route(g map[string][]road.Road, src string, dst string) []road.Road {
	queue := make([]node, 0)

	queue = append(queue, node{
		name:  src,
		roads: []road.Road{},
	})

	for len(queue) != 0 {
		root := queue[0]
		queue = queue[1:]

		if root.name == dst {
			return root.roads
		}

		for _, rd := range g[root.name] {
			roads := make([]road.Road, len(root.roads))
			copy(roads, root.roads)

			queue = append(queue, node{
				name:  rd.Destination,
				roads: append(roads, rd),
			})
		}
	}

	return []road.Road{}
}
