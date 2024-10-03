package structUtil

import (
	"github.com/go-estar/types/mapUtil"
	"testing"
)

type TestStruct struct {
	A map[string]string
	B *map[string]string
	C map[string]string
	D []string
	E *[]string
	F []string
}

func TestStructToSortString(t *testing.T) {
	s := TestStruct{
		A: nil,
		B: nil,
		C: map[string]string{},
		D: nil,
		E: nil,
		F: nil,
	}
	v := ToSortString(s)
	t.Log(v)
}

func TestMapToSortString(t *testing.T) {
	var (
		A map[string]string
		B *map[string]string
		C map[string]string
		D []string
		E *[]string
		F []string
		G TestStruct
	)
	s := map[string]any{
		"A": A,
		"B": B,
		"C": C,
		"D": D,
		"E": E,
		"F": F,
		"G": G,
	}
	v := mapUtil.ToSortString(s)
	t.Log(v)
}
