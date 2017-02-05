package money

// money type. Stored as 64-bit integer with 4 decimal places (hundredth of a
// cent) fixed-point representation for GAAP accurate calculation

type Money int64

const ZERO = Money(0)

func NewInt(i int) Money {
	return Money(i * 100)
}

func NewFloat64(i float64) Money {
	return Money(int(i * 10000))
}

func (m Money) Float64() float64 {
	return float64(m) / 10000.
}

func (m Money) Int() int {
	return int(m / 100)
}

// Round off the money to the nearest cent. This is important to quantize a
// final value after all intermediate calculations are done so that sub-penny
// amounts aren't added to cummulative aggregate calculations
func (m Money) Rounded() Money {
	cents := m / 100
	if m%100 > 49 {
		cents++
	}
	return cents * 100
}
