package dice

type DiceSideMeaning struct {
	Description string
}

type Dice struct {
	sides map[int]DiceSideMeaning
}

func CreateNewDice(sides map[int]DiceSideMeaning) Dice {
	var newDice Dice
	newDice.sides = sides
	return newDice
}
