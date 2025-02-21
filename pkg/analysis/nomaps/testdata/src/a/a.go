package a

type DefinedType map[string]Component

type TypeAlias = map[string]Component

type TypeAlias2 = DefinedType

type NoPhaseTestStruct struct {
	Primitive     int32                `json:"primitive"`
	Components    []Component          `json:"components"`
	MapComponents map[string]Component `json:"mapComponents"` // want "MapComponents should not use a map type, use list type instead"
	DefinedType   DefinedType          `json:"definedType"`   // want "DefinedType should not use a map type, use list type instead"
	TypeAlias     TypeAlias            `json:"typeAlias"`     // want "TypeAlias should not use a map type, use list type instead"
	TypeAlias2    TypeAlias2           `json:"typeAlias2"`    // want "TypeAlias2 should not use a map type, use list type instead"
	Labels        map[string]string    `json:"specialCase"`
}

type Component struct {
	Key   string `json:"key"`
	Value int32  `json:"value"`
}
