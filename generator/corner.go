package generator

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	cornerPreference = []string{
		"outfield",
		"box",
		"closest",
	}
)

func CornerPreference() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s", cornerPreference[rand.Intn(len(cornerPreference))])
}
