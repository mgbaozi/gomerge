package gomerge

import "strings"

type JsonTag struct {
	Name string
	Inline bool
	OmitEmpty bool
}

func ParseJsonTag(tag string) (result JsonTag) {
	options := strings.Split(tag, ",")
	if len(options) == 0 {
		return
	}
	result.Name = options[0]
	for _, option := range options[1:] {
		switch option {
		case "inline":
			result.Inline = true
		case "omitempty":
			result.OmitEmpty = true
		}
	}
	return
}