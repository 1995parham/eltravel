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
	"gitlab.com/1995parham/eltravel/passenger"
	"gitlab.com/1995parham/eltravel/route"
	"time"
)

type node struct {
	ts     time.Time
	name   string
	routes []route.Route
}

// Route find a route from src to dst in the given graph g
func Route(g map[string][]route.Route, src string, dst string, dep time.Time, ps passenger.Passengers) []route.Route {
	queue := make([]node, 0)

	queue = append(queue, node{
		ts:     dep,
		name:   src,
		routes: []route.Route{},
	})

	for len(queue) != 0 {
		root := queue[0]
		queue = queue[1:]

		if root.name == dst {
			return root.routes
		}

		for _, r := range g[root.name] {
			if r.IsAvailable(root.ts, ps) {
				routes := make([]route.Route, len(root.routes))
				copy(routes, root.routes)

				queue = append(queue, node{
					ts:     r.ArrivalTime(root.ts),
					name:   r.Dst(),
					routes: append(routes, r),
				})
			}
		}
	}

	return []route.Route{}
}
