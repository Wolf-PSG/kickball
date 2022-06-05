package generator

func Attribute(modifier string) int32 {
	if modifier == "bronze" {
		return 30
	}

	return 50
}
