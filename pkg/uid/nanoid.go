package uid

import "github.com/aidarkhanov/nanoid"

type NanoGenerator interface {
	NanoId(code string) string
}

type generate struct{}

func NewNanoId() NanoGenerator {
	return &generate{}
}

func (g *generate) NanoId(code string) string {
	nano, err := nanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", nanoid.DefaultSize)
	if err != nil {
		return ""
	}

	return code + "-" + nano
}
