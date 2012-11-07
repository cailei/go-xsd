package xsd

import (
	"fmt"
	"strings"

	util "github.com/metaleap/go-util"
)

var (
	PkgGen = &goPkgSrcMaker {
		BaseTypes: map[string]string {
			"decimal": "complex128",
			"float": "float32",
			"double": "float64",
			"duration": "time.Duration",
			"dateTime": "time.Time",
			"date": "time.Time",
			"hexBinary": "[]byte",
			"base64Binary": "[]byte",
			"integer": "int",
			"nonPositiveInteger": "int",
			"negativeInteger": "int",
			"long": "int64",
			"int": "int32",
			"short": "int16",
			"byte": "byte",
			"nonNegativeInteger": "uint",
			"unsignedLong": "uint64",
			"unsignedInt": "uint32",
			"unsignedShort": "uint16",
			"unsignedByte": "uint8",
			"positiveInteger": "uint",
		},
		Name: "goxsdpkg",
		BaseCodePath: util.BaseCodePath("metaleap", "go-xsd-pkg"),
		BasePath: "github.com/metaleap/go-xsd-pkg",
	}
)

type goPkgSrcMaker struct {
	BaseTypes map[string]string
	BaseCodePath, BasePath, Name, TypePrefix string

	lines []string
	imports map[string]string
}

	func (me *goPkgSrcMaker) append (lines ... string) {
		me.lines = append(me.lines, lines ...)
	}

	func (me *goPkgSrcMaker) appendFmt (format string, fmtArgs ... interface{}) {
		me.append(fmt.Sprintf(format, fmtArgs ...))
	}

	func (me *goPkgSrcMaker) insertFmt (index int, format string, fmtArgs ... interface{}) {
		me.lines = append(me.lines[: index], append([]string { fmt.Sprintf(format, fmtArgs ...) }, me.lines[index : ] ...) ...)
	}

	func (me *goPkgSrcMaker) pascalCase (name string) string {
		return strings.ToUpper(name[: 1]) + name[1 :]
	}

	func (me *goPkgSrcMaker) reinit () {
		me.lines = []string { "package " + me.Name, "" }
		me.imports = map[string]string {}
	}

type makerBag struct {
	Schema *Schema
}

func (me *Schema) makePkg (bag *makerBag) {
	var impPos int
	PkgGen.reinit()
	PkgGen.imports["xsdt"] = "github.com/metaleap/go-xsd/types"
	me.hasElemAnnotation.makePkg(bag)
	me.hasElemsImport.makePkg(bag)
	impPos = len(PkgGen.lines) + 1
	PkgGen.append("import (", ")", "")
	me.hasElemsSimpleType.makePkg(bag)
	for impName, impPath := range PkgGen.imports {
		PkgGen.insertFmt(impPos, "\t%v \"%v\"", impName, impPath)
	}
}