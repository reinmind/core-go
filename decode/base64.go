package decode

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func DecodeBase64(encoded string) string {
	b, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error decoding %v", err)
	}
	return string(b)
}
func EncodeBase64(raw []byte) string {
	s := base64.StdEncoding.EncodeToString(raw)
	return s
}

func Main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s := input.Text()
		fmt.Println(DecodeBase64(s))
	}

}
