package domain

var PayoffMatrix = [3][3]int{
	//		Share  Hold  Take
	/* Share*/ {+4, 0, -3},
	/* Hold */ {+1, 0, -1},
	/* Take */ {+7, 0, -2},
}

// my - мой ход, op (opponent) - ход противника
// счетчик очков на один ход
func Payoff(my, op Act) (int, int) {
	myPayoff := PayoffMatrix[my][op]
	opPayoff := PayoffMatrix[op][my]
	return myPayoff, opPayoff
}
