package parsy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInspector_with_nil_inspector_ops(t *testing.T) {
	var ops *InspectorOptions

	ins := NewInspector(ops)

	assert.NotNil(t, ins)
	assert.NotNil(t, ins.ops)
	assert.False(t, ins.ops.AvoidNil)
}

func TestNewInspector_with_inspector_ops(t *testing.T) {
	ops := &InspectorOptions{
		AvoidNil: true,
	}

	ins := NewInspector(ops)

	assert.NotNil(t, ins)
	assert.NotNil(t, ins.ops)
	assert.True(t, ins.ops.AvoidNil)
}
