// Code generated by typeshare 1.1.0. DO NOT EDIT.
package proto

import "encoding/json"

// Generated type representing the anonymous struct variant `Us` of the `AutofilledBy` Rust enum
type AutofilledByUsInner struct {
	// The UUID for the fill
	Uuid string `json:"uuid"`
}
// Generated type representing the anonymous struct variant `SomethingElse` of the `AutofilledBy` Rust enum
type AutofilledBySomethingElseInner struct {
	// The UUID for the fill
	Uuid string `json:"uuid"`
	// Some other thing
	Thing int `json:"thing"`
}
// Enum keeping track of who autofilled a field
type AutofilledByTypes string
const (
	// This field was autofilled by us
	AutofilledByTypeVariantUs AutofilledByTypes = "Us"
	// Something else autofilled this field
	AutofilledByTypeVariantSomethingElse AutofilledByTypes = "SomethingElse"
)
type AutofilledBy struct{ 
	Type AutofilledByTypes `json:"type"`
	content interface{}
}

func (a *AutofilledBy) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag    AutofilledByTypes   `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	a.Type = enum.Tag
	switch a.Type {
	case AutofilledByTypeVariantUs:
		var res AutofilledByUsInner
		a.content = &res
	case AutofilledByTypeVariantSomethingElse:
		var res AutofilledBySomethingElseInner
		a.content = &res

	}
	if err := json.Unmarshal(enum.Content, &a.content); err != nil {
		return err
	}

	return nil
}

func (a AutofilledBy) MarshalJSON() ([]byte, error) {
    var enum struct {
		Tag    AutofilledByTypes   `json:"type"`
		Content interface{} `json:"content,omitempty"`
    }
    enum.Tag = a.Type
    enum.Content = a.content
    return json.Marshal(enum)
}

func (a AutofilledBy) Us() *AutofilledByUsInner {
	res, _ := a.content.(*AutofilledByUsInner)
	return res
}
func (a AutofilledBy) SomethingElse() *AutofilledBySomethingElseInner {
	res, _ := a.content.(*AutofilledBySomethingElseInner)
	return res
}

func NewAutofilledByTypeVariantUs(content *AutofilledByUsInner) AutofilledBy {
    return AutofilledBy{
        Type: AutofilledByTypeVariantUs,
        content: content,
    }
}
func NewAutofilledByTypeVariantSomethingElse(content *AutofilledBySomethingElseInner) AutofilledBy {
    return AutofilledBy{
        Type: AutofilledByTypeVariantSomethingElse,
        content: content,
    }
}

// Generated type representing the anonymous struct variant `AnonVariant` of the `EnumWithManyVariants` Rust enum
type EnumWithManyVariantsAnonVariantInner struct {
	Uuid string `json:"uuid"`
}
// Generated type representing the anonymous struct variant `AnotherAnonVariant` of the `EnumWithManyVariants` Rust enum
type EnumWithManyVariantsAnotherAnonVariantInner struct {
	Uuid string `json:"uuid"`
	Thing int `json:"thing"`
}
// This is a comment (yareek sameek wuz here)
type EnumWithManyVariantsTypes string
const (
	EnumWithManyVariantsTypeVariantUnitVariant EnumWithManyVariantsTypes = "UnitVariant"
	EnumWithManyVariantsTypeVariantTupleVariantString EnumWithManyVariantsTypes = "TupleVariantString"
	EnumWithManyVariantsTypeVariantAnonVariant EnumWithManyVariantsTypes = "AnonVariant"
	EnumWithManyVariantsTypeVariantTupleVariantInt EnumWithManyVariantsTypes = "TupleVariantInt"
	EnumWithManyVariantsTypeVariantAnotherUnitVariant EnumWithManyVariantsTypes = "AnotherUnitVariant"
	EnumWithManyVariantsTypeVariantAnotherAnonVariant EnumWithManyVariantsTypes = "AnotherAnonVariant"
)
type EnumWithManyVariants struct{ 
	Type EnumWithManyVariantsTypes `json:"type"`
	content interface{}
}

func (e *EnumWithManyVariants) UnmarshalJSON(data []byte) error {
	var enum struct {
		Tag    EnumWithManyVariantsTypes   `json:"type"`
		Content json.RawMessage `json:"content"`
	}
	if err := json.Unmarshal(data, &enum); err != nil {
		return err
	}

	e.Type = enum.Tag
	switch e.Type {
	case EnumWithManyVariantsTypeVariantUnitVariant:
		return nil
	case EnumWithManyVariantsTypeVariantTupleVariantString:
		var res string
		e.content = &res
	case EnumWithManyVariantsTypeVariantAnonVariant:
		var res EnumWithManyVariantsAnonVariantInner
		e.content = &res
	case EnumWithManyVariantsTypeVariantTupleVariantInt:
		var res int
		e.content = &res
	case EnumWithManyVariantsTypeVariantAnotherUnitVariant:
		return nil
	case EnumWithManyVariantsTypeVariantAnotherAnonVariant:
		var res EnumWithManyVariantsAnotherAnonVariantInner
		e.content = &res

	}
	if err := json.Unmarshal(enum.Content, &e.content); err != nil {
		return err
	}

	return nil
}

func (e EnumWithManyVariants) MarshalJSON() ([]byte, error) {
    var enum struct {
		Tag    EnumWithManyVariantsTypes   `json:"type"`
		Content interface{} `json:"content,omitempty"`
    }
    enum.Tag = e.Type
    enum.Content = e.content
    return json.Marshal(enum)
}

func (e EnumWithManyVariants) TupleVariantString() string {
	res, _ := e.content.(*string)
	return *res
}
func (e EnumWithManyVariants) AnonVariant() *EnumWithManyVariantsAnonVariantInner {
	res, _ := e.content.(*EnumWithManyVariantsAnonVariantInner)
	return res
}
func (e EnumWithManyVariants) TupleVariantInt() int {
	res, _ := e.content.(*int)
	return *res
}
func (e EnumWithManyVariants) AnotherAnonVariant() *EnumWithManyVariantsAnotherAnonVariantInner {
	res, _ := e.content.(*EnumWithManyVariantsAnotherAnonVariantInner)
	return res
}

func NewEnumWithManyVariantsTypeVariantUnitVariant() EnumWithManyVariants {
    return EnumWithManyVariants{
        Type: EnumWithManyVariantsTypeVariantUnitVariant,
    }
}
func NewEnumWithManyVariantsTypeVariantTupleVariantString(content string) EnumWithManyVariants {
    return EnumWithManyVariants{
        Type: EnumWithManyVariantsTypeVariantTupleVariantString,
        content: &content,
    }
}
func NewEnumWithManyVariantsTypeVariantAnonVariant(content *EnumWithManyVariantsAnonVariantInner) EnumWithManyVariants {
    return EnumWithManyVariants{
        Type: EnumWithManyVariantsTypeVariantAnonVariant,
        content: content,
    }
}
func NewEnumWithManyVariantsTypeVariantTupleVariantInt(content int) EnumWithManyVariants {
    return EnumWithManyVariants{
        Type: EnumWithManyVariantsTypeVariantTupleVariantInt,
        content: &content,
    }
}
func NewEnumWithManyVariantsTypeVariantAnotherUnitVariant() EnumWithManyVariants {
    return EnumWithManyVariants{
        Type: EnumWithManyVariantsTypeVariantAnotherUnitVariant,
    }
}
func NewEnumWithManyVariantsTypeVariantAnotherAnonVariant(content *EnumWithManyVariantsAnotherAnonVariantInner) EnumWithManyVariants {
    return EnumWithManyVariants{
        Type: EnumWithManyVariantsTypeVariantAnotherAnonVariant,
        content: content,
    }
}

