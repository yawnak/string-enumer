package stringenumer_test

import (
	"io"
	"log"
	"os"

	"github.com/yawnak/string-enumer/pkg/stringenumer"
)

func Example() {
	r, err := stringenumer.Generate(
		stringenumer.Paths("."),
		stringenumer.TypeNames("MyType", "YourType"),
		stringenumer.TextUnmarshaling(true),
	)
	if err != nil {
		log.Fatalln(err)
	}
	_, _ = io.Copy(os.Stdout, r)
}
