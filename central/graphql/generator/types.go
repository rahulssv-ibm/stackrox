package generator

import "reflect"

// TypeAndField defines a field on a specific type.
type TypeAndField struct {
	ParentType reflect.Type
	FieldName  string
}

// TypeWalkParameters tells generator how to walk types
type TypeWalkParameters struct {
	IncludedTypes []reflect.Type
	SkipResolvers []reflect.Type
	SkipFields    []TypeAndField
}
