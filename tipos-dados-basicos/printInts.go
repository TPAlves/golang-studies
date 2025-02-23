package main

import (
	"bytes"
	"fmt"
)

// intsToString é como fmt.Sprint(values), mas com vírgulas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func main() {
	values := []int{1, 2, 3}
	a := fmt.Sprintf("%d\n", values)
	fmt.Print(a)
	fmt.Println(intsToString(values))
}

