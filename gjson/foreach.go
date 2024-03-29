package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func GetForEach() {
	const json = `
{
  "name":"dj",
  "age":18,
  "pets": ["cat", "dog"],
  "contact": {
    "phone": "123456789",
    "email": "dj@example.com"
  }
}`

	pets := gjson.Get(json, "pets")
	pets.ForEach(func(_, pet gjson.Result) bool {
		fmt.Println(pet)
		return true
	})

	contact := gjson.Get(json, "contact")
	contact.ForEach(func(key, value gjson.Result) bool {
		fmt.Println(key, value)
		return true
	})

}
