package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
	"strings"
)

func ToUpperOrLower() {
	gjson.AddModifier("case", func(json, arg string) string {
		if arg == "upper" {
			return strings.ToUpper(json)
		}

		if arg == "lower" {
			return strings.ToLower(json)
		}

		return json
	})

	const json = `{"children": ["Sara", "Alex", "Jack"]}`
	fmt.Println(gjson.Get(json, "children|@case:upper"))
	fmt.Println(gjson.Get(json, "children|@case:lower"))
}
