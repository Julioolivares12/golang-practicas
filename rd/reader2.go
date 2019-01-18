package rd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//
func R1() {
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Simple shell")
	fmt.Println("-------------------------")

	for {
		fmt.Print("->")
		txt, _ := r.ReadString('\n')
		// convert CRLF to LF
		txt = strings.Replace(txt, "\r\n", "", -1)
		if strings.Compare("hi", txt) == 0 {
			fmt.Println("hello, Yourself")
		} else {
			fmt.Printf("Hola %s", txt)
		}
	}
}

//
func Rune1() {
	r := bufio.NewReader(os.Stdin)
	char, _, err := r.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(char)
	switch char {
	case 'A':
		fmt.Println("A key Pressed")
		break
	case 'a':
		fmt.Println("a key Pressed")
		break
	}
}

//
func Scanner1() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
