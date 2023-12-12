package puppy

import (
	"github.com/juhayahoo/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof Woo f!!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBark() string {
	return dog.WhenGrownUp(Barks())
}
