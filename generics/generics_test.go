package generics

import "testing"

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})

	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "hello", "hello")
		AssertNotEqual(t, "hello", "goodbye")
	})

	t.Run("assert bools", func(t *testing.T) {
		AssertTrue(t, true)
		AssertFalse(t, false)
	})
}

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		s := new(Stack[int])

		// check stack is empty
		AssertTrue(t, s.IsEmpty())

		// Push & check stack is not empty
		s.Push(1)
		AssertFalse(t, s.IsEmpty())

		// Add another thing, pop it back
		s.Push(2)
		value, ok := s.Pop()
		AssertEqual(t, value, 2)
		AssertTrue(t, ok)
		value, ok = s.Pop()
		AssertEqual(t, value, 1)
		AssertTrue(t, ok)
		AssertTrue(t, s.IsEmpty())

		// Cannot pop off an empty stack
		val, ok := s.Pop()
		AssertEqual(t, val, 0)
		AssertFalse(t, ok)
	})

	t.Run("stack of strings", func(t *testing.T) {
		s := new(Stack[string])
		AssertTrue(t, s.IsEmpty())

		// Push & check not empty
		s.Push("hello")
		AssertFalse(t, s.IsEmpty())

		// Push again & pop both off
		s.Push("world")
		val, ok := s.Pop()
		AssertEqual(t, val, "world")
		AssertTrue(t, ok)
		val, ok = s.Pop()
		AssertEqual(t, val, "hello")
		AssertTrue(t, ok)
		AssertTrue(t, s.IsEmpty())

		// Cannot pop off empty stack
		val, ok = s.Pop()
		AssertEqual(t, val, "")
		AssertFalse(t, ok)
	})
}
