/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 22-09-2019
 * |
 * | File Name:     route.go
 * +===============================================
 */

package road

// Road represents a route in country with its details
type Road struct {
	Name        string `json:"road_name"`
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Length      int    `json:"length"`
	SpeedLimit  int    `json:"speed_limit"`
}
