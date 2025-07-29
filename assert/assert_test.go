package assert

import (
	"testing"
)

type mockT struct {
	testing.TB
	failed bool
	logs   []string
}

func (m *mockT) Helper() {}
func (m *mockT) Errorf(format string, args ...any) {
	m.failed = true
	m.logs = append(m.logs, format)
}

func newMockT() *mockT {
	return &mockT{}
}

func TestEqual_Assert(t *testing.T) {
	mt := newMockT()
	Equal(mt, 1, 1)
	if mt.failed {
		t.Error("Equal failed on equal values")
	}

	mt = newMockT()
	Equal(mt, 1, 2)
	if !mt.failed {
		t.Error("Equal did not fail on unequal values")
	}
}

func TestNotEqual_Assert(t *testing.T) {
	mt := newMockT()
	NotEqual(mt, 1, 2)
	if mt.failed {
		t.Error("NotEqual failed on different values")
	}

	mt = newMockT()
	NotEqual(mt, 1, 1)
	if !mt.failed {
		t.Error("NotEqual did not fail on equal values")
	}
}

func TestGreater_Assert(t *testing.T) {
	mt := newMockT()
	Greater(mt, 3, 2)
	if mt.failed {
		t.Error("Greater failed on greater value")
	}

	mt = newMockT()
	Greater(mt, 2, 3)
	if !mt.failed {
		t.Error("Greater did not fail on lesser value")
	}
}

func TestLess_Assert(t *testing.T) {
	mt := newMockT()
	Less(mt, 2, 3)
	if mt.failed {
		t.Error("Less failed on lesser value")
	}

	mt = newMockT()
	Less(mt, 3, 2)
	if !mt.failed {
		t.Error("Less did not fail on greater value")
	}
}

func TestGreaterOrEqual_Assert(t *testing.T) {
	mt := newMockT()
	GreaterOrEqual(mt, 3, 3)
	if mt.failed {
		t.Error("GreaterOrEqual failed on equal values")
	}

	mt = newMockT()
	GreaterOrEqual(mt, 2, 3)
	if !mt.failed {
		t.Error("GreaterOrEqual did not fail on lesser value")
	}
}

func TestLessOrEqual_Assert(t *testing.T) {
	mt := newMockT()
	LessOrEqual(mt, 2, 2)
	if mt.failed {
		t.Error("LessOrEqual failed on equal values")
	}

	mt = newMockT()
	LessOrEqual(mt, 3, 2)
	if !mt.failed {
		t.Error("LessOrEqual did not fail on greater value")
	}
}

func TestTrue_Assert(t *testing.T) {
	mt := newMockT()
	True(mt, true)
	if mt.failed {
		t.Error("True failed on true")
	}

	mt = newMockT()
	True(mt, false)
	if !mt.failed {
		t.Error("True did not fail on false")
	}
}

func TestFalse_Assert(t *testing.T) {
	mt := newMockT()
	False(mt, false)
	if mt.failed {
		t.Error("False failed on false")
	}

	mt = newMockT()
	False(mt, true)
	if !mt.failed {
		t.Error("False did not fail on true")
	}
}

func TestIsClose_Assert(t *testing.T) {
	mt := newMockT()
	IsClose(mt, 1.01, 1.00, 0.02)
	if mt.failed {
		t.Error("IsClose failed within tolerance")
	}

	mt = newMockT()
	IsClose(mt, 1.01, 1.00, 0.001)
	if !mt.failed {
		t.Error("IsClose did not fail outside tolerance")
	}
}

func TestZero_Assert(t *testing.T) {
	mt := newMockT()
	Zero(mt, 0)
	if mt.failed {
		t.Error("Zero failed on zero value")
	}

	mt = newMockT()
	Zero(mt, 1)
	if !mt.failed {
		t.Error("Zero did not fail on non-zero value")
	}
}

func TestNotZero_Assert(t *testing.T) {
	mt := newMockT()
	NotZero(mt, 1)
	if mt.failed {
		t.Error("NotZero failed on non-zero value")
	}

	mt = newMockT()
	NotZero(mt, 0)
	if !mt.failed {
		t.Error("NotZero did not fail on zero value")
	}
}

// === ERRORS ===

func TestNil_Assert(t *testing.T) {
	var ptr *int
	mt := newMockT()
	Nil(mt, nil) // Can't use ptr directly as it is not nil when dereferenced
	if mt.failed {
		t.Error("Nil failed on nil value")
	}

	mt = newMockT()
	val := 5
	ptr = &val
	Nil(mt, ptr)
	if !mt.failed {
		t.Error("Nil did not fail on non-nil value")
	}
}

func TestNotNil_Assert(t *testing.T) {
	var ptr *int
	mt := newMockT()
	NotNil(mt, nil) // Can't use ptr as it is not nil when dereferenced
	if !mt.failed {
		t.Error("NotNil did not fail on nil value")
	}

	mt = newMockT()
	val := 5
	ptr = &val
	NotNil(mt, ptr)
	if mt.failed {
		t.Error("NotNil failed on non-nil value")
	}
}

func TestPanic_Assert(t *testing.T) {
	mt := newMockT()
	Panic(mt, func() { panic("boom") })
	if mt.failed {
		t.Error("Panic failed on actual panic")
	}

	mt = newMockT()
	Panic(mt, func() {})
	if !mt.failed {
		t.Error("Panic did not fail when no panic occurred")
	}
}

func TestSafe_Assert(t *testing.T) {
	mt := newMockT()
	Safe(mt, func() {})
	if mt.failed {
		t.Error("Safe failed on non-panicking function")
	}

	mt = newMockT()
	Safe(mt, func() { panic("boom") })
	if !mt.failed {
		t.Error("Safe did not fail on panicking function")
	}
}

// === CONTAINERS ===

func TestContains_Assert(t *testing.T) {
	mt := newMockT()
	Contains(mt, []int{1, 2, 3}, 2)
	if mt.failed {
		t.Error("Contains failed on existing item")
	}

	mt = newMockT()
	Contains(mt, []int{1, 2, 3}, 4)
	if !mt.failed {
		t.Error("Contains did not fail on non-existing item")
	}
}

func TestNotContains_Assert(t *testing.T) {
	mt := newMockT()
	NotContains(mt, []int{1, 2, 3}, 4)
	if mt.failed {
		t.Error("NotContains failed on non-existing item")
	}

	mt = newMockT()
	NotContains(mt, []int{1, 2, 3}, 2)
	if !mt.failed {
		t.Error("NotContains did not fail on existing item")
	}
}

func TestLength_Assert(t *testing.T) {
	mt := newMockT()
	Length(mt, []int{1, 2, 3}, 3)
	if mt.failed {
		t.Error("Length failed on correct length")
	}

	mt = newMockT()
	Length(mt, []int{1, 2, 3}, 4)
	if !mt.failed {
		t.Error("Length did not fail on incorrect length")
	}
}

/// === STRINGS ===

func TestStringContains_Assert(t *testing.T) {
	mt := newMockT()
	StringContains(mt, "hello world", "world")
	if mt.failed {
		t.Error("ContainsString failed on existing substring")
	}

	mt = newMockT()
	StringContains(mt, "hello world", "hill")
	if !mt.failed {
		t.Error("ContainsString did not fail on non-existing substring")
	}

	mt = newMockT()
	StringContains(mt, "hello", "world")
	if !mt.failed {
		t.Error("ContainsString did not fail on non-existing substring")
	}
}

func TestStringNotContains_Assert(t *testing.T) {
	mt := newMockT()
	NotStringContains(mt, "hello world", "world")
	if !mt.failed {
		t.Error("NotContainsString did not fail on existing substring")
	}

	mt = newMockT()
	NotStringContains(mt, "hello world", "hill")
	if mt.failed {
		t.Error("NotContainsString failed on non-existing substring")
	}

	mt = newMockT()
	NotStringContains(mt, "hello", "world")
	if mt.failed {
		t.Error("NotContainsString failed on non-existing substring")
	}
}

func TestPrefixString_Assert(t *testing.T) {
	mt := newMockT()
	PrefixString(mt, "hello world", "hello")
	if mt.failed {
		t.Error("PrefixString failed on correct prefix")
	}

	mt = newMockT()
	PrefixString(mt, "hello world", "world")
	if !mt.failed {
		t.Error("PrefixString did not fail on incorrect prefix")
	}
}

func TestNotPrefixString_Assert(t *testing.T) {
	mt := newMockT()
	NotPrefixString(mt, "hello world", "world")
	if mt.failed {
		t.Error("NotPrefixString failed on incorrect prefix")
	}

	mt = newMockT()
	NotPrefixString(mt, "hello world", "hello")
	if !mt.failed {
		t.Error("NotPrefixString did not fail on correct prefix")
	}
}

func TestSuffixString_Assert(t *testing.T) {
	mt := newMockT()
	SuffixString(mt, "hello world", "world")
	if mt.failed {
		t.Error("SuffixString failed on correct suffix")
	}

	mt = newMockT()
	SuffixString(mt, "hello world", "hello")
	if !mt.failed {
		t.Error("SuffixString did not fail on incorrect suffix")
	}
}

func TestNotSuffixString_Assert(t *testing.T) {
	mt := newMockT()
	NotSuffixString(mt, "hello world", "hello")
	if mt.failed {
		t.Error("NotSuffixString failed on incorrect suffix")
	}

	mt = newMockT()
	NotSuffixString(mt, "hello world", "world")
	if !mt.failed {
		t.Error("NotSuffixString did not fail on correct suffix")
	}
}
