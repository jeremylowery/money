package money

import (
	"github.com/dustin/go-humanize"
)

func (m Money) Human() string {
	v := m.Rounded().Float64()
	return "$" + humanize.FormatFloat("#,###.##", v)
}

func (m Money) Print() string {
	v := m.Rounded().Float64()
	return humanize.FormatFloat("#.##", v)
}
