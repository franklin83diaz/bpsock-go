package tags

import "unicode"

// Tag16
type Tag16 struct {
	name string
}

func NewTag16(s string) Tag16 {
	if len(s) > 16 {
		panic("Tag16: tag name too long")
	}
	//check if start with a number
	if unicode.IsDigit(rune(s[0])) {
		panic("Tag16: tag name must start with a letter")
	}
	return Tag16{name: s}
}

// new Tag16 ephemera
func NewTag16Eph(s string) Tag16 {
	if len(s) > 16 {
		panic("Tag16: tag name too long")
	}
	return Tag16{name: s}
}

func (t Tag16) Name() string {
	return t.name
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

func (t Tag8) Name() string {
	return t.name
}
