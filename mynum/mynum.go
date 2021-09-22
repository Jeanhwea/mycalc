package mynum

type MyNum struct {
	X int
	Y int
}

func Add(a, b MyNum) MyNum {
	return MyNum{a.X + b.X, a.Y + b.Y}
}

func Sub(a, b MyNum) MyNum {
	return MyNum{a.X - b.X, a.Y - b.Y}
}

func Mul(a, b MyNum) MyNum {
	x := a.X*b.X - a.Y*b.Y
	y := a.X*b.Y + a.Y*b.X
	return MyNum{x, y}
}
