package context_learn

import "testing"

func TestNewCtx(t *testing.T) {
	NewCtx()
}

func TestCtxWithValue(t *testing.T) {
	CtxWithValue()
}

func TestAccessParentContextValue(t *testing.T) {
	AccessParentContextValue()
}
