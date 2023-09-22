package calculate

import "fmt"

func Kuis1() {
	// hitung rata-rata
	scores := [...]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int


	sum := 0
	for _, value := range scores {
		// length := len(scores)
		result := value + sum
		sum = result 

		if value >= 90 {
			goodScores = append(goodScores, value)
		}
	}

	result := float64(sum) / float64(len(scores))

	fmt.Println(result)
	fmt.Println(goodScores)



}