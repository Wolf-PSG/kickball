package generator

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	style = []string{
		"attack",
		"defence",
		"all-rounder",
	}
)

func Style() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%s", style[rand.Intn(len(style))])
}
