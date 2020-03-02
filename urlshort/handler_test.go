package urlshort

import (
	"fmt"
	"testing"
)

func Test_ParseYaml(*testing.T){
	b := []byte(`
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`)
	m := parseYaml(b)
	fmt.Println(m)
}

