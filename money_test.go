package money

import "testing"

type TestStruct struct {
	name      string
	num       int64
	formatted string
}

type RoundTestStruct struct {
	name    string
	num     int64
	rounded int64
}

type FloatTestStruct struct {
	name string
	num  float64
}

func TestHuman(t *testing.T) {
	tests := []TestStruct{
		{"default", 12345600, "$1,234.56"},
		{"millions", 188812345600, "$18,881,234.56"},
		{"billions", 22188812345600, "$2,218,881,234.56"},
		{"pennies", 500, "$0.05"},
		{"dollars", 340500, "$34.05"},
	}

	for _, test := range tests {
		m := Money(test.num)
		got := m.Human()
		if got != test.formatted {
			t.Errorf("On %v (%v), got %v, wanted %v",
				test.name, test.num, got, test.formatted)
		}
	}

}

func TestRounded(t *testing.T) {
	tests := []RoundTestStruct{
		{"none", 123400, 123400},
		{"down", 123449, 123400},
		{"up", 123450, 123500},
		{"up2", 123460, 123500},
	}

	for _, test := range tests {
		i := Money(test.num)
		e := Money(test.rounded)
		if i.Rounded() != e {
			t.Errorf("On %v, got %v, wanted %v", test.name, i.Rounded(), e)
		}
	}
}

func TestPrint(t *testing.T) {
	tests := []TestStruct{
		{"default", 12345600, "1234.56"},
		{"round down close", 12345649, "1234.56"},
		{"round up close", 12345650, "1234.57"},
		{"round up not close", 12345699, "1234.57"},
	}

	for _, test := range tests {
		m := Money(test.num)
		got := m.Print()
		if got != test.formatted {
			t.Errorf("On %v (%v), got %v, wanted %v",
				test.name, test.num, got, test.formatted)
		}
	}

}

func TestFloat(t *testing.T) {
	tests := []FloatTestStruct{
		{"pennies", 3445.43},
		{"dollars", 12.00},
		{"subpennis", 12.0099},
		{"subpennis2", 0.0001},
	}

	for _, test := range tests {
		m := NewFloat64(test.num)
		got := m.Float64()
		if got != test.num {
			t.Errorf("On %v, got %v, wanted %v", test.name, got, test.num)
		}
	}
}
