package utility

import (
	"fmt"
	"strings"
)

func Echo(args []string) {
	str := strings.Join(args, " ")
	if len(str) > 1 && ((str[0] == '"' && str[len(str)-1] == '"') || (str[0] == '\'' && str[len(str)-1] == '\'')) {
		str = str[1 : len(str)-1]
	}
	fmt.Println(str)
}
