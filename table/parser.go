package table

import (
	"reflect"
)

type ParserConfig struct {
	TagsFieldsOnly   bool // if true, only tagged fields will be parsed.
	UnexportedFields bool // if true, unexported fields will be parsed too.
	MaxDepth         int  // max depth of the struct, -1 means no limit.
}

// Parser should be implemented by all available reflect-based parsers.
//
// See `StructParser`(struct{} type), `SliceParser`(slice[] type), `MapParser`(map type) and `JSONParser`(any type).
// Manually registering of a parser is a valid option (although not a recommendation), see `RegisterParser` for more.
type Parser interface {
	// Parse Why not `ParseRows` and `ParseHeaders`?
	// Because type map has not a specific order, order can change at different runtimes,
	// so we must keep record on the keys order the first time we fetche them (=> see `MapParser#ParseRows`, `MapParser#ParseHeaders`).
	Parse(v reflect.Value, filters []RowFilter) (headers []TableRowInterface, rows [][]string, numbers []int)

	// SetConfig sets the parser's configuration.
	SetConfig(config *ParserConfig)
}

// The built'n type parsers, all except `JSONParser` are directly linked to the `Print/PrintHeadList` functions.
var (
	DefaultConfig = ParserConfig{
		TagsFieldsOnly:   false,
		UnexportedFields: false,
		MaxDepth:         -1,
	}
	StructParser = &structParser{Config: &DefaultConfig}
	SliceParser  = &sliceParser{Config: &DefaultConfig}
	MapParser    = &mapParser{Config: &DefaultConfig}
	JSONParser   = new(jsonParser)
)

// WhichParser returns the available `Parser` for the "typ" type; Slice, Map, Struct...
func WhichParser(typ reflect.Type, config *ParserConfig) Parser {
	if p, ok := availableParsers[typ.Kind()]; ok {
		p.SetConfig(config)
		return p
	}
	return nil // it can return nil.
}

var availableParsers = map[reflect.Kind]Parser{
	reflect.Struct: StructParser,
	reflect.Slice:  SliceParser,
	reflect.Map:    MapParser,
}

// RegisterParser sets a parser based on its kind of type.
// It overrides any existing element on that kind, each Parser reflects a single kind of type.
//
// It can be used at the initialization of the program to register a custom Parser, see `StructParser` for example.
// It's not designed to be safe to use it inside many different routines at the same time.
func RegisterParser(kind reflect.Kind, parser Parser) {
	availableParsers[kind] = parser
}

// like reflect.Indirect but for types and reflect.Interface types too.
func indirectType(typ reflect.Type) reflect.Type {
	if kind := typ.Kind(); kind == reflect.Interface || kind == reflect.Ptr {
		return typ.Elem()
	}

	return typ
}

// like reflect.Indirect but reflect.Interface values too.
func indirectValue(val reflect.Value) reflect.Value {
	if kind := val.Kind(); kind == reflect.Interface || kind == reflect.Ptr {
		return val.Elem()
	}

	return val
}
