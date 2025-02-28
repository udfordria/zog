package internals

import (
	zconst "github.com/udfordria/zog/zconst"
)

// Zog Context interface. This is the interface that is passed to schema tests, pre and post transforms
type Ctx interface {
	// Get a value from the context
	Get(key string) any
	// Deprecated: Use Ctx.AddIssue() instead
	// Please don't depend on this interface it may change
	NewError(p *PathBuilder, e ZogIssue)
	// Adds an issue to the schema execution.
	AddIssue(e ZogIssue)
	// Please don't depend on this interface it may change
	HasErrored() bool
}

func NewExecCtx(errs ZogIssues, fmter IssueFmtFunc) *ExecCtx {
	c := ExecCtxPool.Get().(*ExecCtx)
	c.Fmter = fmter
	c.Errors = errs
	return c
}

type ExecCtx struct {
	Fmter  IssueFmtFunc
	Errors ZogIssues
	m      map[string]any
}

func (c *ExecCtx) HasErrored() bool {
	return !c.Errors.IsEmpty()
}

func (c *ExecCtx) SetIssueFormatter(fmter IssueFmtFunc) {
	c.Fmter = fmter
}

func (c *ExecCtx) Set(key string, val any) {
	if c.m == nil {
		c.m = make(map[string]any)
	}
	c.m[key] = val
}

func (c *ExecCtx) Get(key string) any {
	return c.m[key]
}

// Adds a ZogIssue to the execution context.
func (c *ExecCtx) AddIssue(e ZogIssue) {
	if e.Message() == "" {
		c.Fmter(e, c)
	}
	c.Errors.Add(e.Path(), e)
}

// Deprecated: Use Ctx.AddIssue() instead
// This is old interface. It will be removed soon
func (c *ExecCtx) NewError(path *PathBuilder, e ZogIssue) {
	c.Errors.Add(path.String(), e)
}

// Internal. Used to format errors
func (c *ExecCtx) FmtErr(e ZogIssue) {
	if e.Message() != "" {
		return
	}
	c.Fmter(e, c)
}

func (c *ExecCtx) NewSchemaCtx(val any, destPtr any, path *PathBuilder, dtype zconst.ZogType) *SchemaCtx {
	c2 := SchemaCtxPool.Get().(*SchemaCtx)
	c2.ExecCtx = c
	c2.Val = val
	c2.DestPtr = destPtr
	c2.Path = path
	c2.DType = dtype
	c2.CanCatch = false
	c2.HasCaught = false
	return c2
}

func (c *ExecCtx) NewValidateSchemaCtx(valPtr any, path *PathBuilder, dtype zconst.ZogType) *SchemaCtx {
	c2 := SchemaCtxPool.Get().(*SchemaCtx)
	c2.ExecCtx = c
	c2.Val = valPtr
	c2.DestPtr = nil
	c2.Path = path
	c2.DType = dtype
	c2.CanCatch = false
	c2.HasCaught = false
	return c2
}

func (c *ExecCtx) Free() {
	ExecCtxPool.Put(c)
}

type SchemaCtx struct {
	*ExecCtx
	Val       any
	DestPtr   any
	Path      *PathBuilder
	DType     zconst.ZogType
	CanCatch  bool
	HasCaught bool
}
type TestCtx struct {
	*SchemaCtx
	Test *Test
}

func (c *SchemaCtx) Issue() ZogIssue {
	// TODO handle catch here
	e := ZogIssuePool.Get().(*ZogErr)
	e.C = ""
	e.EPath = c.Path.String()
	e.Err = nil
	e.Msg = ""
	e.ParamsM = nil
	e.Typ = c.DType
	e.Val = c.Val
	return e
}

// Please don't depend on this method it may change
func (c *SchemaCtx) IssueFromTest(test *Test, val any) ZogIssue {
	e := ZogIssuePool.Get().(*ZogErr)
	e.C = test.IssueCode
	e.EPath = c.Path.String()
	e.Err = nil
	e.Msg = ""
	e.Typ = c.DType
	e.Val = val
	e.ParamsM = test.Params
	if test.IssueFmtFunc != nil {
		test.IssueFmtFunc(e, c)
	}
	if test.IssuePath != "" {
		e.EPath = test.IssuePath
	}
	return e
}

// Please don't depend on this method it may change
func (c *SchemaCtx) IssueFromCoerce(err error) ZogIssue {
	e := ZogIssuePool.Get().(*ZogErr)
	e.C = zconst.IssueCodeCoerce
	e.EPath = c.Path.String()
	e.Err = nil
	e.Msg = ""
	e.Typ = c.DType
	e.Val = c.Val
	e.Err = err
	return e
}

// Please don't depend on this method it may change
// Wraps an error in a ZogIssue if it is not already a ZogIssue
func (c *SchemaCtx) IssueFromUnknownError(err error) ZogIssue {
	zerr, ok := err.(ZogIssue)
	if !ok {
		return c.Issue().SetError(err)
	}
	return zerr
}

// Frees the context to be reused
func (c *SchemaCtx) Free() {
	SchemaCtxPool.Put(c)
}

func (c *TestCtx) Issue() ZogIssue {
	// TODO handle catch here
	zerr := ZogIssuePool.Get().(*ZogErr)
	zerr.C = c.Test.IssueCode
	zerr.EPath = c.Path.String()
	zerr.Err = nil
	zerr.Msg = ""
	zerr.ParamsM = c.Test.Params
	zerr.Typ = c.DType
	zerr.Val = c.Val
	return zerr
}

func (c *TestCtx) FmtErr(e ZogIssue) {
	if e.Message() != "" {
		return
	}

	if c.Test.IssueFmtFunc != nil {
		c.Test.IssueFmtFunc(e, c)
		return
	}

	c.SchemaCtx.FmtErr(e)
}

func (c *TestCtx) AddIssue(e ZogIssue) {
	c.FmtErr(e)
	c.Errors.Add(c.Path.String(), e)
}
