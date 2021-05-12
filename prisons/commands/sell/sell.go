package sell

type SubSell string

type Sell struct {
	Sub    SubSell
	Amount int
}

func (s SubSell) SubName() string {
	return "hand"
}
