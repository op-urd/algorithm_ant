package P38

import "fmt"

var queue []string

func BFS(n, m int, cells [][]string) int {
	sx, sy, err := getStartXY(cells)
	if err != nil {
		panic(err)
	}

	

	return 0
}

func getStartXY(cells [][]string)(int, int, error) {
	for i, line := range cells {
		for j, s := range line {
			if s == "S" {
				return i, j, nil
			}
		}
	}
	return -1, -1, fmt.Errorf("They have not start position")
}