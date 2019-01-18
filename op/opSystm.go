package op

import (
	"fmt"
	"os/exec"
	"runtime"
)

//estas funciones sirven en windows con MINGW64
func GetSystemOP() {

	var os = runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("your sistem is windows")
	}
}
func Dir() {
	out, err := exec.Command("dir").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)

	out2, err := exec.Command("pwd").Output()
	if err != nil {
		fmt.Println("%s", err)
	}
	output2 := string(out2[:])
	fmt.Println(output2)
}

func Ls() {
	out, err := exec.Command("ls", "-ltr").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	output := string(out[:])
	fmt.Println(output)
}
