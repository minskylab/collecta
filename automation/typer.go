package main

import (

	"strings"

	"github.com/facebookincubator/ent/schema/field"
	"github.com/minskylab/collecta/errors"
)

func gqlTypeFromTypeInfo(t *field.TypeInfo) (string, error) {
	if t.Ident != "" {
		if strings.HasPrefix(t.Ident, "map") {
			return "Map", nil
		}

		switch t.Ident {
		case "uuid.UUID":
			return "ID", nil
		case "[]string":
			return "[String!]", nil
		case "[]int":
			return "[Int!]", nil
		case "[]bool":
			return "[Boolean!]", nil
		default:
			return t.Ident, nil
		}
	}

	switch t.Type {
	case field.TypeInvalid:
		return "", errors.New("invalid type")
	case field.TypeBool:
		return "Boolean", nil
	case field.TypeTime:
		return "Time", nil
	case field.TypeJSON:
		return "Map", nil
	case field.TypeUUID:
		return "ID", nil
	case field.TypeBytes:
		return "String", nil
	case field.TypeEnum:
		return strings.ToUpper(string(t.Type.String()[0])) + t.Type.String()[1:], nil
	case field.TypeString:
		return "String", nil
	case field.TypeInt |
		 field.TypeInt8|
		 field.TypeInt16|
		 field.TypeInt32|
		 field.TypeInt64|
		 field.TypeUint|
		 field.TypeUint8|
		 field.TypeUint16|
		 field.TypeUint32|
		 field.TypeUint64:
		return "Int", nil
	case field.TypeFloat32 |
		 field.TypeFloat64:
		return "Float", nil
	default:
		return "", errors.New("invalid type")
	}
}