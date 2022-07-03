package piscine

func DivMod(a int, b int, div *int, mod *int) {
	c := a / b
	*div = c
	d := a % b
	*mod = d
}
