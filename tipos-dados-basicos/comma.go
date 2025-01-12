// comma insere virgulas em uma string que repesenta um inteiro decimal
// não negativo
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buf bytes.Buffer
	values := strings.Split(s, ".")
	n := len(s)
	if n <= 3 {
		return s
	}
	n = len(values[0])
	for i := 0; i < n; i++ {
		if i > 0 && (n-i)%3 == 0 && s[i-1] != '.' {
			buf.WriteByte('.')
		}
		buf.WriteByte(s[i])
	}
	for i, v := range values {
		if i != 0 {
			buf.WriteString(fmt.Sprintf(".%s", v))
		}
	}
	return buf.String()
}

func main() {
	number := comma("15203.40")
	fmt.Println(number)

}
