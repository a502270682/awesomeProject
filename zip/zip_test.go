package zip

import (
	"bytes"
	"testing"
)

func TestZip(t *testing.T) {
	elements := []string{"https://gd4.alicdn.com/imgextra/i2/634491/O1CN01T9PnWc1j2vIMn6gk2_!!634491.jpg", "https://gd4.alicdn.com/imgextra/i2/634491/O1CN01T9PnWc1j2vIMn6gk2_!!634491.jpg"}
	err := GetZip(new(bytes.Buffer), elements)
	if err != nil {
		t.Fatal(err)
	}
}
