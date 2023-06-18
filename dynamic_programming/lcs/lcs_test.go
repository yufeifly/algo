package lcs

import (
	"fmt"
	"testing"
)

func TestLCS(t *testing.T) {
	l := LCS("abcde", "acde")
	fmt.Println(l)
}
