package main

import "fmt"

func dice(sl [][]int, a, b, c int) {
	tempslice := [][]int{{}, {}, {}}
	score := [][]int{{a}, {b}, {c}}

	for i := range sl {
		for _, v2 := range sl[i] {
			if v2 == 6 {
				score[i][0] = score[i][0] + 1
			}
			if i == len(sl)-1 {
				if v2 == 1 {
					tempslice[0] = append(tempslice[0], 1)
				}
				if v2 > 1 && v2 < 6 {
					tempslice[i] = append(tempslice[i], v2)
				}

			} else {
				if v2 == 1 {
					tempslice[i+1] = append(tempslice[i+1], v2)
				}
				if v2 > 1 && v2 < 6 {
					tempslice[i] = append(tempslice[i], v2)
				}

			}
		}
	}
	fmt.Println("Player 1 =", tempslice[0], " Score =", score[0])
	fmt.Println("Player 2 =", tempslice[1], " Score =", score[1])
	fmt.Println("Player 3 =", tempslice[2], " Score =", score[2])
	fmt.Println()
}

func main() {
	dice([][]int{{3, 6, 1, 3}, {2, 4, 5, 5}, {1, 2, 5, 6}}, 0, 0, 0)
	dice([][]int{{1, 2, 6}, {4, 3, 1, 3, 3}, {1, 6}}, 1, 0, 1)
	dice([][]int{{1, 6}, {2, 5, 6, 4, 6}, {1}}, 2, 0, 2)
	dice([][]int{{1}, {3, 4, 5, 5}, {}}, 3, 2, 2)
}
