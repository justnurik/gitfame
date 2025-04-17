package programm

import (
	"fmt"
	"strconv"
)

type PositiveInt uint

func (p *PositiveInt) Type() string {
	return "uint"
}

func (p *PositiveInt) Set(value string) error {
	v, err := strconv.Atoi(value)
	if err != nil {
		return fmt.Errorf("workers-count должен быть числом")
	}
	if v <= 0 {
		return fmt.Errorf("workers-count должен быть положительным числом")
	}
	*p = PositiveInt(v)
	return nil
}

func (p PositiveInt) String() string {
	return fmt.Sprintf("%d", p)
}
