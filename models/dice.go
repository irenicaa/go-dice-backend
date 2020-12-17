package models

import (
	"fmt"
	"strconv"
)

// Dice ...
type Dice struct {
	Tries int
	Faces int
}

// String ...
func (dice Dice) String() string {
	if dice.Tries == 1 {
		return "d" + strconv.Itoa(dice.Faces)
	}

	return fmt.Sprintf("%dd%d", dice.Tries, dice.Faces)
}
