package kt_utils

import (
	"fmt"

	"github.com/sanity-io/litter"
)

var (
	NIL_VALUE = "<nil>"
)

// Prints the given 'theVar' value to the console nicely - even if it is a struct with pointer fields! 'prettyPrint' controls if you want one liner or more
// readable multi line.
//
// NOTE: if 'theVar' is an error then .Error() function is used. If 'theVar' implements fmt.Stringer so it has a .String() function then that is used. In this
// case 'prettyPrint' is ignored.
func PrintVar(theVar any, prettyPrint bool) {

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
// NOTE: if 'theVar' is an error then .Error() function is used. If 'theVar' implements fmt.Stringer so it has a .String() function then that is used. In this
// case 'prettyPrint' is ignored.
//
// **IMPORTANT NOTE!** Do NOT use it in log lines! Because if log is omitted due to log level then you wasted resource to construct the string! In that case use
// `VarPrinter struct` instead!
func PrintVarS(theVar any, prettyPrint bool) string {

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
type VarPrinter struct {
	TheVar      any
	PrettyPrint bool
}

func (vp *VarPrinter) Print(prettyPrint bool) {
	if vp == nil {
		fmt.Print(NIL_VALUE)
	} else {
		PrintVar(vp.TheVar, prettyPrint)
	}
}

func (vp *VarPrinter) PrintS(prettyPrint bool) string {
	if vp == nil {
		return NIL_VALUE
	} else {
		return PrintVarS(vp.TheVar, prettyPrint)
	}
}

// This is the fmt.Stringer interface implementation - so when someone simply does a toString() on the struct.
func (vp VarPrinter) String() string {
	return vp.PrintS(vp.PrettyPrint)
}
