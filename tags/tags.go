package tags

// Tag16
type Tag16 struct {
	name string
}

func NewTag16(s string) Tag16 {
	if len(s) > 16 {
		panic("Tag16: tag name too long")
	}
	return Tag16{name: s}
}

// Tag8
type Tag8 struct {
	name string
}

func NewTag8(s string) Tag8 {
	if len(s) > 8 {
		panic("Tag8: tag name too long")
	}
	return Tag8{name: s}
}
