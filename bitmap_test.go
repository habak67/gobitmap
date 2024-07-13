package gobitmap

import "testing"

// Tests Set and Has
func Test_Set(t *testing.T) {
	tests := []struct {
		name     string
		setBit   []int
		checkBit []int
	}{
		{"set single", []int{2}, []int{2}},
		{"set multiple bits", []int{2, 4, 23}, []int{2, 4, 23}},
		{"set multiple times", []int{2, 2}, []int{2}},
		{"set edge", []int{0, 63}, []int{0, 63}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := EmptyBitMap
			for _, bit := range test.setBit {
				m = m.Set(bit)
			}
			checkMap(t, uint64(m), 64, test.checkBit)
		})
	}
}

// Tests Clear and Has
func Test_Clear(t *testing.T) {
	tests := []struct {
		name     string
		setBit   []int
		clearBit []int
		checkBit []int
	}{
		{"clear single", []int{2}, []int{2}, []int{}},
		{"clear single other set", []int{2, 4}, []int{2}, []int{4}},
		{"clear not set", []int{2}, []int{4}, []int{2}},
		{"clear multiple bits", []int{2, 4, 6}, []int{2, 6}, []int{4}},
		{"clear multiple times", []int{2, 4}, []int{2, 2}, []int{4}},
		{"clear edge", []int{0, 4, 63}, []int{0, 63}, []int{4}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := EmptyBitMap
			for _, bit := range test.setBit {
				m = m.Set(bit)
			}
			for _, bit := range test.clearBit {
				m = m.Clear(bit)
			}
			checkMap(t, uint64(m), 64, test.checkBit)
		})
	}
}

// Tests Toggle and Has
func Test_Toggle(t *testing.T) {
	tests := []struct {
		name      string
		setBit    []int
		toggleBit []int
		checkBit  []int
	}{
		{"toggle single set", []int{2, 4}, []int{2}, []int{4}},
		{"toggle single cleared", []int{4}, []int{2}, []int{2, 4}},
		{"double toggle set", []int{2, 4}, []int{2, 2}, []int{2, 4}},
		{"double toggle cleared", []int{4}, []int{2, 2}, []int{4}},
		{"toggle multiple bits set", []int{2, 4, 6}, []int{2, 6}, []int{4}},
		{"toggle multiple bits cleared", []int{4}, []int{2, 6}, []int{2, 4, 6}},
		{"toggle compound", []int{4}, []int{2, 6, 2}, []int{4, 6}},
		{"toggle edge", []int{0, 4, 63}, []int{0, 63}, []int{4}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := EmptyBitMap
			for _, bit := range test.setBit {
				m = m.Set(bit)
			}
			for _, bit := range test.toggleBit {
				m = m.Toggle(bit)
			}
			checkMap(t, uint64(m), 64, test.checkBit)
		})
	}
}

// Tests Empty and variable EmptyBitMap
func Test_Empty(t *testing.T) {
	m := EmptyBitMap
	if !m.Empty() {
		t.Errorf("expected empty bitmap (got=%v)", m)
	}
	m = m.Set(1)
	if m.Empty() {
		t.Errorf("expected non-empty bitmap")
	}
	m = m.Clear(1)
	if !m.Empty() {
		t.Errorf("expected empty bitmap (got=%v)", m)
	}
}

func Test_String(t *testing.T) {
	tests := []struct {
		name   string
		setBit []int
		s      string
	}{
		{
			"empty", []int{},
			"[]",
		},
		{
			"multiple bits set", []int{0, 2, 4, 23, 63},
			"[0,2,4,23,63]",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			m := EmptyBitMap
			for _, bit := range test.setBit {
				m = m.Set(bit)
			}
			if m.String() != test.s {
				t.Errorf("unexpected string\nexp=%v\ngot=%v", test.s, m.String())
			}
		})
	}
}

func Test_Set_OutOfRange_low(t *testing.T) {
	defer func() { recover() }()
	m := EmptyBitMap
	m.Set(-1)
	t.Errorf("expected test to panic")
}

func Test_Set_OutOfRange_high(t *testing.T) {
	defer func() { recover() }()
	m := EmptyBitMap
	m.Set(64)
	t.Errorf("expected test to panic")
}

func Test_Clear_OutOfRange_low(t *testing.T) {
	defer func() { recover() }()
	m := EmptyBitMap
	m.Clear(-1)
	t.Errorf("expected test to panic")
}

func Test_Clear_OutOfRange_high(t *testing.T) {
	defer func() { recover() }()
	m := EmptyBitMap
	m.Clear(64)
	t.Errorf("expected test to panic")
}

func Test_Toggle_OutOfRange_low(t *testing.T) {
	defer func() { recover() }()
	m := EmptyBitMap
	m.Toggle(-1)
	t.Errorf("expected test to panic")
}

func Test_Toggle_OutOfRange_high(t *testing.T) {
	defer func() { recover() }()
	m := EmptyBitMap
	m.Toggle(64)
	t.Errorf("expected test to panic")
}

func Test_Has_OutOfRange_low(t *testing.T) {
	defer func() { recover() }()
	m := EmptyBitMap
	m.Has(-1)
	t.Errorf("expected test to panic")
}

func Test_Has_OutOfRange_high(t *testing.T) {
	defer func() { recover() }()
	m := EmptyBitMap
	m.Has(64)
	t.Errorf("expected test to panic")
}

func checkMap(t *testing.T, m uint64, size int, cb []int) {
	checkBits := make(map[int]bool, size)
	for i := 0; i < size; i++ {
		checkBits[i] = false
	}
	for _, set := range cb {
		checkBits[set] = true
	}
	for i := 0; i < 64; i++ {
		if m&(1<<i) != 0 != checkBits[i] {
			t.Errorf("bit %v differs (has=%v, expected=%v)", i, m&(1<<i), checkBits[i])
		}
	}
}
