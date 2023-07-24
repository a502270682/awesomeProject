package gjson

import (
	"fmt"
	"github.com/tidwall/gjson"
)

type S struct {
	ID string `json:"id"`
}

func GetByLevel() {
	json := `{"name":{"first":"li","last":"dj"},"age":18}`
	lastName := gjson.Get(json, "name.last")
	fmt.Println("last name:", lastName.String())

	age := gjson.Get(json, "age")
	fmt.Println("age:", age.Int())
}
