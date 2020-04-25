package pacote

import (
	"ProjetoDemo/pacote/subpacote"
	"fmt"
	"reflect"
	"strconv"
)

const CONSTANTE_PACOTE string = "PACOTE"

func Looping() {
	println("	Looping")
	vowels := []string{"A", "E", "I", "O", "U"}

	for index, value := range vowels {
		fmt.Println(index, value)
	}
}

func PacotePrint(numero int) {
	fmt.Println("PacotePrint" + strconv.Itoa(numero))
	subpacoteResult, _ := subpacote.SubPacotePrint("numero")

	fmt.Println("	subpacote: " + subpacoteResult)
	fmt.Println("	constantes: " + GOPATH)

	//	Looping()
}

func Listas() {

}

func Maps() {
	fmt.Println("	Maps")
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	prsa, prs := m["aisygfiuahsga"]
	fmt.Println("prsa:", prsa)
	fmt.Println("prs:", prs)

	prs1a, prs1b := m["k1"]
	fmt.Println("prs1a:", prs1a)
	fmt.Println("prs1b:", prs1b)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}

func Slices() {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}

func Pointers() {
	var ponteiroVar int = 10
	fmt.Println("Pointers:1 ", strconv.Itoa(ponteiroVar))
	ponteiroFunc(&ponteiroVar)
	fmt.Println(reflect.TypeOf(&ponteiroVar))
	fmt.Println(&ponteiroVar)

	fmt.Println("Pointers:2 ", strconv.Itoa(ponteiroVar))
}

func ponteiroFunc(ponteiro *int) {
	*ponteiro++
	var copia int = *ponteiro
	copia++
	fmt.Println("ponteiroFunc: " + strconv.Itoa(copia))

	copiaRef := ponteiro
	fmt.Println(reflect.TypeOf(copiaRef))
	fmt.Println(copiaRef)

}
