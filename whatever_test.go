package zog

import (
	"testing"

	"github.com/Oudwins/zog/zconst"
	"github.com/stretchr/testify/assert"
)

func TestWhateverPrimitive(t *testing.T) {
	// in := 10
	var out any
	s := Whatever()

	err := s.Parse("papa", &out)
	assert.Equal(t, "papa", out)
	assert.Nil(t, err)

	err = s.Parse(22, &out)
	assert.Equal(t, 22, out)
	assert.Nil(t, err)
}

func TestWhateverParseSetCoercerPassThrough(t *testing.T) {
	var dest any
	validator := Whatever()
	errs := validator.Parse("5", &dest)
	assert.Empty(t, errs)
}

func TestWhateverInStruct(t *testing.T) {
	type TestStruct struct {
		Value *int
	}

	s := Struct(Schema{
		"value": Whatever(),
	})
	in := map[string]any{
		"value": 10,
	}
	var out TestStruct
	err := s.Parse(in, &out)

	assert.Nil(t, err)
	assert.NotNil(t, out)
	assert.NotNil(t, out.Value)
	assert.Equal(t, 10, *out.Value)
}

func TestWhateverWhateverInStruct(t *testing.T) {
	type TestStruct struct {
		Value **int
	}

	s := Struct(Schema{
		"value": Whatever(),
	})
	in := map[string]any{
		"value": 10,
	}
	var out TestStruct
	// empty input
	err := s.Parse(nil, &out)
	assert.NotNil(t, err)
	assert.Equal(t, zconst.IssueCodeCoerce, err[zconst.ISSUE_KEY_ROOT][0].Code())
	assert.Nil(t, out.Value)

	// good input
	err = s.Parse(in, &out)

	assert.Nil(t, err)
	assert.NotNil(t, out)
	assert.NotNil(t, out.Value)
	assert.NotNil(t, *out.Value)
	assert.Equal(t, 10, **out.Value)
}

func TestWhateverNestedStructs(t *testing.T) {
	type Inner struct {
		Value *int
	}
	type Outer struct {
		Inner *Inner
	}

	schema := Struct(Schema{
		"inner": Whatever(),
	})

	var out Outer
	data := map[string]any{
		"inner": map[string]any{
			"value": 10,
		},
	}

	err := schema.Parse(data, &out)
	assert.Nil(t, err)
	assert.NotNil(t, out.Inner)
	assert.NotNil(t, out.Inner.Value)
	assert.Equal(t, 10, *out.Inner.Value)
}

func TestWhateverInSlice(t *testing.T) {
	schema := Slice(Whatever())
	var out []*int

	data := []any{10, 20, 30}
	err := schema.Parse(data, &out)

	assert.Nil(t, err)
	assert.Len(t, out, 3)
	assert.Equal(t, 10, *out[0])
	assert.Equal(t, 20, *out[1])
	assert.Equal(t, 30, *out[2])
}

func TestWhateverSliceStruct(t *testing.T) {
	type TestStruct struct {
		Value int
	}

	schema := Slice(Whatever())
	var out []*TestStruct

	data := []any{
		map[string]any{"value": 10},
		map[string]any{"value": 20},
		map[string]any{"value": 30},
	}
	err := schema.Parse(data, &out)

	assert.Nil(t, err)
	assert.Len(t, out, 3)
	assert.Equal(t, 10, out[0].Value)
	assert.Equal(t, 20, out[1].Value)
	assert.Equal(t, 30, out[2].Value)
}

func TestWhateverToStruct(t *testing.T) {
	var dest any
	s := Whatever()
	in := map[string]any{
		"value": 10,
	}
	err := s.Parse(in, &dest)
	assert.Nil(t, err)
	assert.NotNil(t, dest)
	t.Log(dest)
}

func TestWhateverToSlice(t *testing.T) {
	var dest *[]*int
	s := Whatever()
	err := s.Parse([]any{10, 20, 30}, &dest)
	assert.Nil(t, err)
	assert.NotNil(t, dest)
	assert.Len(t, *dest, 3)
	assert.Equal(t, 10, *(*dest)[0])
	assert.Equal(t, 20, *(*dest)[1])
	assert.Equal(t, 30, *(*dest)[2])
}
