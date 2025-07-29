package assert

import (
	"testing"
)

func TestEqual_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).Equal(1, 1)
	if mt.failed {
		t.Error("Equal failed on equal values")
	}

	mt = newMockT()
	That(mt).Equal(1, 2)
	if !mt.failed {
		t.Error("Equal did not fail on unequal values")
	}
}

func TestNotEqual_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).NotEqual(1, 2)
	if mt.failed {
		t.Error("NotEqual failed on different values")
	}

	mt = newMockT()
	That(mt).NotEqual(1, 1)
	if !mt.failed {
		t.Error("NotEqual did not fail on equal values")
	}
}

func TestTrue_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).True(true)
	if mt.failed {
		t.Error("True failed on true")
	}

	mt = newMockT()
	That(mt).True(false)
	if !mt.failed {
		t.Error("True did not fail on false")
	}
}

func TestFalse_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).False(false)
	if mt.failed {
		t.Error("False failed on false")
	}

	mt = newMockT()
	That(mt).False(true)
	if !mt.failed {
		t.Error("False did not fail on true")
	}
}

/// === ERRORS ===

func TestNil_Fluent(t *testing.T) {
	var ptr *int
	mt := newMockT()
	That(mt).Nil(nil)
	if mt.failed {
		t.Error("Nil failed on nil value")
	}

	mt = newMockT()
	val := 5
	ptr = &val
	That(mt).Nil(ptr)
	if !mt.failed {
		t.Error("Nil did not fail on non-nil value")
	}
}

func TestNotNil_Fluent(t *testing.T) {
	var ptr *int
	mt := newMockT()
	That(mt).NotNil(nil)
	if !mt.failed {
		t.Error("NotNil did not fail on nil value")
	}

	mt = newMockT()
	val := 5
	ptr = &val
	That(mt).NotNil(ptr)
	if mt.failed {
		t.Error("NotNil failed on non-nil value")
	}
}

func TestPanic_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).Panic(func() { panic("boom") })
	if mt.failed {
		t.Error("Panic failed on actual panic")
	}

	mt = newMockT()
	That(mt).Panic(func() {})
	if !mt.failed {
		t.Error("Panic did not fail when no panic occurred")
	}
}

func TestSafe_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).Safe(func() {})
	if mt.failed {
		t.Error("Safe failed on non-panicking function")
	}

	mt = newMockT()
	That(mt).Safe(func() { panic("boom") })
	if !mt.failed {
		t.Error("Safe did not fail on panicking function")
	}
}

// === CONTAINERS ===

func TestContains_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).Contains([]any{1, 2, 3}, 2)
	if mt.failed {
		t.Error("Contains failed on existing item")
	}

	mt = newMockT()
	That(mt).Contains([]any{1, 2, 3}, 4)
	if !mt.failed {
		t.Error("Contains did not fail on non-existing item")
	}
}

// === STRINGS ===

func TestNotContains_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).NotContains([]any{1, 2, 3}, 4)
	if mt.failed {
		t.Error("NotContains failed on non-existing item")
	}

	mt = newMockT()
	That(mt).NotContains([]any{1, 2, 3}, 2)
	if !mt.failed {
		t.Error("NotContains did not fail on existing item")
	}
}

func TestLength_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).Length([]any{1, 2, 3}, 3)
	if mt.failed {
		t.Error("Length failed on correct length")
	}

	mt = newMockT()
	That(mt).Length([]any{1, 2, 3}, 4)
	if !mt.failed {
		t.Error("Length did not fail on incorrect length")
	}
}

func TestStringContains_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).StringContains("hello world", "world")
	if mt.failed {
		t.Error("StringContains failed on existing substring")
	}

	mt = newMockT()
	That(mt).StringContains("hello world", "test")
	if !mt.failed {
		t.Error("StringContains did not fail on non-existing substring")
	}
}

func TestNotStringContains_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).NotStringContains("hello world", "test")
	if mt.failed {
		t.Error("NotStringContains failed on non-existing substring")
	}

	mt = newMockT()
	That(mt).NotStringContains("hello world", "world")
	if !mt.failed {
		t.Error("NotStringContains did not fail on existing substring")
	}
}

func TestPrefixString_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).PrefixString("hello world", "hello")
	if mt.failed {
		t.Error("PrefixString failed on correct prefix")
	}

	mt = newMockT()
	That(mt).PrefixString("hello world", "world")
	if !mt.failed {
		t.Error("PrefixString did not fail on incorrect prefix")
	}
}

func TestNotPrefixString_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).NotPrefixString("hello world", "world")
	if mt.failed {
		t.Error("NotPrefixString failed on incorrect prefix")
	}

	mt = newMockT()
	That(mt).NotPrefixString("hello world", "hello")
	if !mt.failed {
		t.Error("NotPrefixString did not fail on correct prefix")
	}
}

func TestSuffixString_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).SuffixString("hello world", "world")
	if mt.failed {
		t.Error("SuffixString failed on correct suffix")
	}

	mt = newMockT()
	That(mt).SuffixString("hello world", "hello")
	if !mt.failed {
		t.Error("SuffixString did not fail on incorrect suffix")
	}
}

func TestNotSuffixString_Fluent(t *testing.T) {
	mt := newMockT()
	That(mt).NotSuffixString("hello world", "hello")
	if mt.failed {
		t.Error("NotSuffixString failed on incorrect suffix")
	}

	mt = newMockT()
	That(mt).NotSuffixString("hello world", "world")
	if !mt.failed {
		t.Error("NotSuffixString did not fail on correct suffix")
	}
}
