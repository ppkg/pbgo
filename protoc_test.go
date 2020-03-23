package pbgo

import "testing"

func Test_Generate(t *testing.T) {
	if err := Generate(); err != nil {
		println(err.Error())
	}
}
