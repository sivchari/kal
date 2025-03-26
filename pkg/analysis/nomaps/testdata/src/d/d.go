package d

type (
	MapStringComponent    map[string]Component
	PtrMapStringComponent *map[string]Component
	MapStringInt          map[string]int
	MapIntString          map[int]string
)

type (
	MapStringComponentAlias           = map[string]Component
	MapStringPtrComponentAlias        = *map[string]Component
	MapStringIntAlias                 = map[string]int
	DefinedMapStringComponentAlias    = MapStringComponent
	DefinedMapStringComponentPtrAlias = *MapStringComponent
)

type (
	MapStringGenerics[V any]             map[string]V
	MapIntGenerics[V any]                map[int]V
	MapComparableKeyString[K comparable] map[K]string
	MapComparableKeyInt[K comparable]    map[K]int
)

type NoMapsTestStruct struct {
	Primitive        int32                 `json:"primitive"`
	Components       []Component           `json:"components"`
	MapComponents    map[string]Component  `json:"mapComponents"` // want "MapComponents should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapComponents *map[string]Component `json:"mapComponents"` // want "MapComponents should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt     map[string]int        `json:"mapStringInt"`
	Labels           map[string]string     `json:"specialCase"`
}

type NoMapsTestStructWithDefiningType struct {
	MapStringComponent    MapStringComponent    `json:"mapStringComponent"`    // want "MapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	PtrMapStringComponent PtrMapStringComponent `json:"ptrMapStringComponent"` // want "PtrMapStringComponent should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInt          MapStringInt          `json:"mapStringInt"`
	MapIntString          MapIntString          `json:"mapIntString"`
}

type NoMapsTestStructWithAlias struct {
	MapStringComponentAlias           MapStringComponentAlias           `json:"mapStringComponentAlias"`    // want "MapStringComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	MapStringPtrComponentAlias        MapStringPtrComponentAlias        `json:"mapStringPtrComponentAlias"` // want "MapStringPtrComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	MapStringIntAlias                 MapStringIntAlias                 `json:"mapStringIntAlias"`
	DefinedMapStringComponentAlias    DefinedMapStringComponentAlias    `json:"definedMapStringComponentAlias"`    // want "DefinedMapStringComponentAlias should not use a map type, use a list type with a unique name/identifier instead"
	DefinedMapStringComponentPtrAlias DefinedMapStringComponentPtrAlias `json:"definedMapStringComponentPtrAlias"` // want "DefinedMapStringComponentPtrAlias should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithGenerics[K comparable, V any] struct {
	MapStringGenerics      MapStringGenerics[V]      `json:"mapStringGenerics"`      // want "MapStringGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapIntGenerics         MapIntGenerics[V]         `json:"mapIntGenerics"`         // want "MapIntGenerics should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyString MapComparableKeyString[K] `json:"mapComparableKeyString"` // want "MapComparableKeyString should not use a map type, use a list type with a unique name/identifier instead"
	MapComparableKeyInt    MapComparableKeyInt[K]    `json:"mapComparableKeyInt"`    // want "MapComparableKeyInt should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestCompositeLiteral struct {
	MapStringArray     map[string][5]string         `json:"arrayMap"`     // want "MapStringArray should not use a map type, use a list type with a unique name/identifier instead"
	MapStringStruct    map[string]struct{}          `json:"structMap"`    // want "MapStringStruct should not use a map type, use a list type with a unique name/identifier instead"
	MapStringPointer   map[string]*string           `json:"pointerMap"`   // want "MapStringPointer should not use a map type, use a list type with a unique name/identifier instead"
	MapStringFunc      map[string]func()            `json:"funcMap"`      // want "MapStringFunc should not use a map type, use a list type with a unique name/identifier instead"
	MapStringInterface map[string]any               `json:"interfaceMap"` // want "MapStringInterface should not use a map type, use a list type with a unique name/identifier instead"
	MapStringSlice     map[string][]string          `json:"sliceMap"`     // want "MapStringSlice should not use a map type, use a list type with a unique name/identifier instead"
	MapStringMap       map[string]map[string]string `json:"mapMap"`       // want "MapStringMap should not use a map type, use a list type with a unique name/identifier instead"
	MapStringChan      map[string]chan string       `json:"chanMap"`      // want "MapStringChan should not use a map type, use a list type with a unique name/identifier instead"
}

type NoMapsTestStructWithEmbedded struct {
	NoMapsTestStruct
	NoMapsTestStructWithDefiningType
	NoMapsTestStructWithGenerics[string, Component]
	NoMapsTestStructWithAlias
}

type Component struct {
	Key   string `json:"key"`
	Value int32  `json:"value"`
}
