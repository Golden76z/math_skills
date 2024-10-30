package main

func PrintNbr(nbr int) int {
	return nbr
}

func ForEach(f func(int), a []int) {
	for b := 0; b < len(a)-1; b++ {
		f(a[b])
	}
}
