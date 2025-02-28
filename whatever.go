package zog

import (
	"reflect"

	"github.com/udfordria/zog/conf"
	p "github.com/udfordria/zog/internals"
	"github.com/udfordria/zog/zconst"
)

var _ ComplexZogSchema = &WhateverSchema{}

type WhateverSchema struct {
	// preTransforms  []p.PreTransform
	tests []p.Test
	// schema   ZogSchema
	// required *p.Test
	// postTransforms []p.PostTransform
	// defaultVal     *any
	// catch          *any
}

func (v *WhateverSchema) getType() zconst.ZogType {
	return zconst.TypePtr
}

func (v *WhateverSchema) setCoercer(c conf.CoercerFunc) {
}

// Whatever creates a Whatever ZogSchema
func Whatever() *WhateverSchema {
	return &WhateverSchema{
		tests: []p.Test{},
	}
}

// Parse the data into the destination Whatever
func (v *WhateverSchema) Parse(data any, dest any, options ...ExecOption) p.ZogIssueMap {
	errs := p.NewErrsMap()
	defer errs.Free()
	ctx := p.NewExecCtx(errs, conf.IssueFormatter)
	defer ctx.Free()
	for _, opt := range options {
		opt(ctx)
	}
	path := p.NewPathBuilder()
	defer path.Free()
	v.process(ctx.NewSchemaCtx(data, dest, path, v.getType()))

	return errs.M
}

func (v *WhateverSchema) process(ctx *p.SchemaCtx) {
	d := reflect.ValueOf(ctx.Val)
	if !d.IsValid() {
		return
	}

	k := d.Kind()
	if (k == reflect.Ptr || k == reflect.Slice || k == reflect.Map || k == reflect.Chan) && d.IsNil() {
		return
	}

	rv := reflect.ValueOf(ctx.DestPtr)
	rv.Elem().Set(d)
}

// Validates a Whatever Whatever
func (v *WhateverSchema) Validate(data any, options ...ExecOption) p.ZogIssueMap {
	errs := p.NewErrsMap()
	defer errs.Free()
	ctx := p.NewExecCtx(errs, conf.IssueFormatter)
	defer ctx.Free()
	for _, opt := range options {
		opt(ctx)
	}
	path := p.NewPathBuilder()
	defer path.Free()
	v.validate(ctx.NewValidateSchemaCtx(data, path, v.getType()))
	return errs.M
}

func (v *WhateverSchema) validate(ctx *p.SchemaCtx) {
}

// Validate Existing Whatever
