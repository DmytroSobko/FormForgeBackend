package athlete

import "fmt"

type Stat float64

func NewStat(value float64) (Stat, error) {
	if value < 0 || value > 100 {
		return 0, fmt.Errorf("stat must be between 0 and 100")
	}
	return Stat(value), nil
}

func (s Stat) Value() float64 {
	return float64(s)
}
