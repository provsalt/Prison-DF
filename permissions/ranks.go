package permissions

const (
	A = iota
	B
	C
	D
	E
	F
	G
	H
	I
	J
	K
	L
	M
	N
	O
	P
	Q
	R
	S
	T
	U
	V
	W
	X
	Y
	Z
)

const (
	Nons = iota + 100
	Coal
	Copper
	Gold
	Diamond
	Emerald
	Youtuber
)

const (
	Helper = iota + 1000
	Moderator
	Manager
	Owner
)

func GetAll() map[string]int {
	ranks := map[string]int{
		"A":        A,
		"B":        B,
		"C":        C,
		"D":        D,
		"E":        E,
		"F":        F,
		"G":        G,
		"H":        H,
		"I":        I,
		"J":        J,
		"K":        K,
		"L":        L,
		"M":        M,
		"N":        N,
		"O":        O,
		"P":        P,
		"Q":        Q,
		"R":        R,
		"S":        S,
		"T":        T,
		"U":        U,
		"V":        V,
		"W":        W,
		"X":        X,
		"Y":        Y,
		"Z":        Z,
		"Nons":     Nons,
		"Coal":     Coal,
		"Copper":   Copper,
		"Gold":     Gold,
		"Diamond":  Diamond,
		"Emerald":  Emerald,
		"Youtuber": Youtuber,
	}
	return ranks
}
