package rules

var threeFoldCard [3]int
var PlayerA []int = make([]int, 17)
var PlayerB []int = make([]int, 17)
var PlayerC []int = make([]int, 17)

func generateFoldCards() {
	for i := 0; i < 3; i++ {
		threeFoldCard[i] = GetRandomCard()
	}
}

func GeneratePlayerCards() {
	for i := 0; i < 17; i++ {
		PlayerA[i] = GetRandomCard()
	}
	for i := 0; i < 17; i++ {
		PlayerB[i] = GetRandomCard()
	}
	for i := 0; i < 17; i++ {
		PlayerC[i] = GetRandomCard()
	}
}
