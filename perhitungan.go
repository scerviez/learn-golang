package main

import (
	"fmt"
)

func main() {
	angka := 2

  angka3 := 90

  printer := angka3 - angka

  angka8 := printer + angka3

  angka7 := angka8 - printer

  angkawk := angka8 * angka8

  wow := angka7 * angkawk

  fmt.Println(fmt.Sprintf("%d - %d = %d + %d - %d x %d = -%d", angka3, angka, printer, angka8, angka7, angkawk, wow))
}
