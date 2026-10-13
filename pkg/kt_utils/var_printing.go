package kt_utils

import (
	"fmt"

	"github.com/sanity-io/litter"
)

var (
	NIL_VALUE = "<nil>"
)

// LogStringer is an optional log-oriented string form for a value.
//
// It is honored only by VarPrinter (Print / PrintS / String). PrintVar / PrintVarS do not call
// LogString() — they keep fmt.Stringer / error / litter — so LogString and String can differ.
type LogStringer interface {
	LogString() string
}

// Prints the given 'theVar' value to the console nicely - even if it is a struct with pointer fields! 'prettyPrint' controls if you want one liner or more
// readable multi line.
//
// NOTE: if 'theVar' implements fmt.Stringer then .String() is used. Else if it is an error then .Error() is used. In those
// cases 'prettyPrint' is ignored. Otherwise litter is used.
//
// LogStringer is not used here — see VarPrinter.
func PrintVar(theVar any, prettyPrint bool) {
	PrintVarWithTransform(theVar, prettyPrint, nil)
}

// Same as PrintVar but applies optional 'transform' to 'theVar' before formatting. Nil transform is a no-op.
// Transform runs eagerly in this function — for lazy transform in logs use VarPrinter.Transform.
func PrintVarWithTransform(theVar any, prettyPrint bool, transform func(any) any) {
	if transform != nil {
		theVar = transform(theVar)
	}
	switch v := theVar.(type) {
	case fmt.Stringer:
		fmt.Println(v.String())
		return
	case error:
		fmt.Println(v.Error())
		return
	}

	// for now we use https://github.com/sanity-io/litter
	litter.Options{
		Compact:           !prettyPrint,
		StripPackageNames: false,
	}.Dump(theVar)
}

// Prints the given 'theVal' value into a string - even if it is a struct with pointer fields! 'prettyPrint' controls if you want one liner or more readable
// multi line.
//
// NOTE: if 'theVar' implements fmt.Stringer then .String() is used. Else if it is an error then .Error() is used. In those
// cases 'prettyPrint' is ignored. Otherwise litter is used.
//
// LogStringer is not used here — see VarPrinter.
//
// **IMPORTANT NOTE!** Do NOT use it in log lines! Because if log is omitted due to log level then you wasted resource to construct the string! In that case use
// `VarPrinter struct` instead!
func PrintVarS(theVar any, prettyPrint bool) string {
	return PrintVarSWithTransform(theVar, prettyPrint, nil)
}

// Same as PrintVarS but applies optional 'transform' to 'theVar' before formatting. Nil transform is a no-op.
// Transform runs eagerly in this function — for lazy transform in logs use VarPrinter.Transform.
func PrintVarSWithTransform(theVar any, prettyPrint bool, transform func(any) any) string {
	if transform != nil {
		theVar = transform(theVar)
	}
	switch v := theVar.(type) {
	case fmt.Stringer:
		return v.String()
	case error:
		return v.Error()
	}

	// for now we use https://github.com/sanity-io/litter
	return litter.Options{
		Compact:           !prettyPrint,
		StripPackageNames: false,
	}.Sdump(theVar)
}

type Printer interface {
	Print(prettyPrint bool)
}

type SPrinter interface {
	PrintS(prettyPrint bool)
}

// You can quickly construct this struct as wrapper around your variable (mostly structs) to print them - WHEN(!) needed.
//
// This is ideal to be used in logging - as if log is dropped due to level for example then you do not construct the string and waste it! For example you can do
// this:
//
//	myStruct := MyStruct{...}
//	LOGGER.Debug("and my struct was: %s", helper.VarPrinter{TheVar: myStruct})
//
// This way the string will be only constructed when LOGGER's level is DEBUG - othwise not at all
//
// For []byte text payloads use BytesAsString: true — do not use TheVar: string(bytes) as that converts eagerly:
//
//	LOGGER.Debug("payload: %s", kt_utils.VarPrinter{TheVar: bytes, BytesAsString: true})
//
// For expensive prepare (e.g. secret masking copy) use Transform — it runs only when formatting is invoked:
//
//	LOGGER.Debug("cfg: %s", kt_utils.VarPrinter{TheVar: cfg, Transform: func(v any) any { return MaskSecrets(v) }})
//
// If TheVar (after Transform) implements LogStringer, VarPrinter uses LogString() instead of Stringer/litter.
type VarPrinter struct {
	TheVar        any
	PrettyPrint   bool
	BytesAsString bool
	Transform     func(any) any
}

func (vp *VarPrinter) Print(prettyPrint bool) {
	if vp == nil {
		fmt.Print(NIL_VALUE)
		return
	}
	theVar := vp.TheVar
	if vp.Transform != nil {
		theVar = vp.Transform(theVar)
	}
	if vp.BytesAsString {
		if b, ok := theVar.([]byte); ok {
			fmt.Println(string(b))
			return
		}
	}
	if ls, ok := theVar.(LogStringer); ok {
		fmt.Println(ls.LogString())
		return
	}
	PrintVar(theVar, prettyPrint)
}

func (vp *VarPrinter) PrintS(prettyPrint bool) string {
	if vp == nil {
		return NIL_VALUE
	}
	theVar := vp.TheVar
	if vp.Transform != nil {
		theVar = vp.Transform(theVar)
	}
	if vp.BytesAsString {
		if b, ok := theVar.([]byte); ok {
			return string(b)
		}
	}
	if ls, ok := theVar.(LogStringer); ok {
		return ls.LogString()
	}
	return PrintVarS(theVar, prettyPrint)
}

// This is the fmt.Stringer interface implementation - so when someone simply does a toString() on the struct.
func (vp VarPrinter) String() string {
	return vp.PrintS(vp.PrettyPrint)
}
