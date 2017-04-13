package messages

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	temp := Temp(107)
	b, err := json.Marshal(temp)
	if err != nil {
		t.Fail()
	}
	fmt.Println(string(b))
}
