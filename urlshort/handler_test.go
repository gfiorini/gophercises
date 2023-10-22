package urlshort

import (
	"fmt"
	"testing"
)
import "github.com/go-yaml/yaml"

func TestYmlParsing(t *testing.T) {

	yml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	source := []byte("mapper:" + yml)
	var mapper Mapper
	err := yaml.Unmarshal(source, &mapper)
	if err != nil {
		panic(err)
	}
	fmt.Printf("--- mapper:\n%v\n\n", mapper)

}
