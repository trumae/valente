package elements

import (
	"log"
	"testing"
)

func TestSparkline(t *testing.T) {
	sl := Sparkline{}
	s := sl.String()
	log.Println(s)

	//if s != "" {
	//	t.Error("Expected '<span></span>', got", s)
	//}

	sl = Sparkline{
		Values: []int{1, 2, 3, 4, 5, 6},
	}
	s = sl.String()
	log.Println(s)

	sl = Sparkline{
		Values: []int{1, 2, 3, 4, 5, 6},
		Options: map[string]string{
			"type":     "bar",
			"barColor": "blue",
		},
	}
	s = sl.String()
	log.Println(s)
}
