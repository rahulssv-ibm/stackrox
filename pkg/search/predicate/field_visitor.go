package predicate

import (
	"reflect"
)

// FieldPath represents the fields we need to access to get to the field we care about.
type FieldPath []reflect.StructField

// VisitFields calls the input function on all paths to a field in the input toWalk type.
func VisitFields(toWalk interface{}, visitField func(fieldPath FieldPath)) {
	visitChildrenRec(nil, reflect.TypeOf(toWalk), []reflect.StructField{}, visitField)
}

// Search children for search tags. This lists all of the types that may have children with search fields.
func visitChildrenRec(parentType, currentType reflect.Type, path FieldPath, visitField func(fieldPath FieldPath)) {
	switch currentType.Kind() {
	case reflect.Struct:
		visitStructFields(currentType, path, visitField)
	case reflect.Ptr:
		visitElemField(currentType, path, visitField)
	case reflect.Interface:
		visitInterfaceFields(parentType, currentType, path, visitField)
	case reflect.Array, reflect.Slice:
		visitElemField(currentType, path, visitField)
	case reflect.Map:
		visitMapFields(currentType, path, visitField)
	}
}

func visitStructFields(currentType reflect.Type, path FieldPath, visitField func(fieldPath FieldPath)) {
	// For each field of the input type.
	for i := 0; i < currentType.NumField(); i++ {
		field := currentType.Field(i)

		// Create a new path through this field.
		newPath := append(path, field)

		// Visit a copy of the field path, so that visitor users can store/modify it.
		pathCopy := append(FieldPath{}, newPath...)
		visitField(pathCopy)

		// Recursively visit the fields children.
		visitChildrenRec(currentType, field.Type, newPath, visitField)
	}
}

// If the parent is a map type, search the child types recursively.
func visitMapFields(currentType reflect.Type, path FieldPath, visitField func(fieldPath FieldPath)) {
	visitChildrenRec(currentType, currentType.Key(), path, visitField)
	visitChildrenRec(currentType, currentType.Elem(), path, visitField)
}

// If the parent is a slice, array, or pointer type, search it's element(s) recursively.
func visitElemField(currentType reflect.Type, path FieldPath, visitField func(fieldPath FieldPath)) {
	visitChildrenRec(currentType, currentType.Elem(), path, visitField)
}

// Assumes that the interface type is a OneOf field, because everything else compiled from a proto will be a Ptr to a
// concrete type.
func visitInterfaceFields(parentType, currentType reflect.Type, path FieldPath, visitField func(fieldPath FieldPath)) {
	ptrToParent := reflect.PtrTo(parentType)
	method, ok := ptrToParent.MethodByName("XXX_OneofFuncs")
	if !ok {
		panic("XXX_OneofFuncs should exist for all protobuf oneofs")
	}
	out := method.Func.Call([]reflect.Value{reflect.New(parentType)})
	actualOneOfFields := out[3].Interface().([]interface{})
	for _, f := range actualOneOfFields {
		typ := reflect.TypeOf(f)
		if typ.Implements(currentType) {
			visitChildrenRec(currentType, typ, path, visitField)
		}
	}
}
