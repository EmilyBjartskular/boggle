package mode

// NewStandardMode returns a board for Standard Boggle
func NewStandardMode() *GameMode {
	name := "Standard Boggle"

	// 4x4 Grid
	size := Size{4, 4}

	// 16 Die
	diceList := DiceList{
		Die{"A", "A", "C", "I", "O", "T"},
		Die{"A", "B", "I", "L", "T", "Y"},
		Die{"A", "B", "J", "M", "O", "Qu"},
		Die{"A", "C", "D", "E", "M", "P"},
		Die{"A", "C", "E", "L", "R", "S"},
		Die{"A", "D", "E", "N", "V", "Z"},
		Die{"A", "H", "M", "O", "R", "S"},
		Die{"B", "I", "F", "O", "R", "X"},
		Die{"D", "E", "N", "O", "S", "W"},
		Die{"D", "K", "N", "O", "T", "U"},
		Die{"E", "E", "F", "H", "I", "Y"},
		Die{"E", "G", "K", "L", "U", "Y"},
		Die{"E", "G", "I", "N", "T", "V"},
		Die{"E", "H", "I", "N", "P", "S"},
		Die{"E", "L", "P", "S", "T", "U"},
		Die{"G", "I", "L", "R", "U", "W"},
	}

	// Match scores per die count
	matchScore := MatchScore{
		3: 1,
		4: 1,
		5: 2,
		6: 3,
		7: 5,
		8: 11,
	}

	result := &GameMode{
		Name:       name,
		Size:       size,
		MatchScore: matchScore,
		DiceList:   diceList,
		ValidMatch: standardValidMatch,
	}

	return result
}

func standardValidMatch(match Match) bool {

	return true
}
