package zog

import (
	"reflect"

	"github.com/Oudwins/zog/conf"
	p "github.com/Oudwins/zog/internals"
	"github.com/Oudwins/zog/zconst"
)

var _ ComplexZogSchema = &WhateverSchema{}

type WhateverSchema struct {
	// preTransforms  []p.PreTransform
	tests []p.Test
	// schema   ZogSchema
	required *p.Test
	// postTransforms []p.PostTransform
	// defaultVal     *any
	// catch          *any
}

func (v *WhateverSchema) getType() zconst.ZogType {
	return zconst.TypePtr
}

func (v *WhateverSchema) setCoercer(c conf.CoercerFunc) {
	// v.schema.setCoercer(c)
}

// Whatever creates a Whatever ZogSchema
func Whatever() *WhateverSchema {
	return &WhateverSchema{
		tests: []p.Test{},
		// schema: schema,
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
	// v.process(ctx.NewSchemaCtx(data, dest, path, v.getType()))
	d := reflect.ValueOf(data)
	if !d.IsValid() {
		return errs.M
	}
	k := d.Kind()
	if (k == reflect.Ptr || k == reflect.Slice || k == reflect.Map || k == reflect.Chan) && d.IsNil() {
		return errs.M
	}

	rv := reflect.ValueOf(dest)
	rv.Elem().Set(d)

	return errs.M
}

func (v *WhateverSchema) process(ctx *p.SchemaCtx) {
	/*
		// TODO this is a mess. But couldn't figure out a simple way to support top level optional structs without doing this.
		// Companion code to this codde is in struct.go > process
		subCtx := ctx.NewSchemaCtx(ctx.Val, nil, ctx.Path, v.schema.getType())
		var err error
		if fn, ok := ctx.Val.(p.DpFactory); ok {
			ctx.Val, err = fn()
			if err != nil {
				ctx.AddIssue(subCtx.IssueFromUnknownError(err))
				return
			}
		}
		// End of messy code

		isZero := p.IsParseZeroValue(ctx.Val, ctx)
		if isZero {
			if v.required != nil {
				// We set the destination type to the schema type because Whatever doesn't have any issue messages. They pass through to the schema type
				ctx.AddIssue(ctx.IssueFromTest(v.required, ctx.Val).SetDType(v.schema.getType()))
			}
			return
		}
		rv := reflect.ValueOf(ctx.DestPtr)
		destPtr := rv.Elem()
		if destPtr.IsNil() {
			// this sets the primitive also
			newVal := reflect.New(destPtr.Type().Elem())
			// this generates a new nil Whatever
			//newVal := reflect.Zero(destPtr.Type())
			destPtr.Set(newVal)
		}
		di := destPtr.Interface()
		subCtx.DestPtr = di
		v.schema.process(subCtx)
	*/

	/*
		rv := reflect.ValueOf(ctx.DestPtr)
		destPtr := rv.Elem()
		if destPtr.IsNil() {
			// this sets the primitive also
			newVal := reflect.New(destPtr.Type().Elem())
			// this generates a new nil Whatever
			//newVal := reflect.Zero(destPtr.Type())
			destPtr.Set(newVal)
		}
	*/
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
	rv := reflect.ValueOf(ctx.Val)
	destPtr := rv.Elem()
	/*
		if !destPtr.IsValid() || destPtr.IsNil() {
			if v.required != nil {
				// We set the destination type to the schema type because Whatever doesn't have any issue messages. They pass through to the schema type
				ctx.AddIssue(ctx.IssueFromTest(v.required, ctx.Val).SetDType(v.schema.getType()))
			}
			return
		}*/
	di := destPtr.Interface()
	ctx.Val = di
	// v.schema.validate(ctx.NewValidateSchemaCtx(di, ctx.Path, v.schema.getType()))
}

// Validate Existing Whatever
