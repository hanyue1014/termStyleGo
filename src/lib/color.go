package lib

import "fmt"

/*
escape sequence: \x1b or \e or \033 (will be going with \x1b)

example: \x1b[codem must end with 'm' to indicate escape code ended

\x1b[0m clears the previous code effects
*/

func TestColor() {
	fmt.Println("\x1b[38;2;255;0;100mHello Idk what this color will be\x1b[0m")
}
