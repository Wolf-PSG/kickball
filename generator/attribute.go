package generator

import (
	"fmt"
	"math/rand"
	"time"
)

func Attribute(modifier string) int32 {
	rand.Seed(time.Now().UnixNano())
	var attribute int32
	switch modifier {
	case "BRONZE":
		attribute = int32(rand.Intn((30 - 20)) + 20)
	case "SILVER":
		attribute = int32(rand.Intn((50 - 30)) + 30)
	case "GOLD":
		attribute = int32(rand.Intn((60 - 40)) + 40)
	case "PLATINUM":
		attribute = int32(rand.Intn((70 - 50)) + 50)
	case "DIAMOND":
		fmt.Println(modifier)
		attribute = int32(rand.Intn((90 - 70)) + 70)
	default:
		attribute = 30
	}

	return attribute
}
