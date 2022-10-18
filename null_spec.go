package vjson

// NullFieldSpec is a type used for parsing an NullField
type NullFieldSpec struct {
	Name string    `mapstructure:"name" json:"name"`
	Type fieldType `json:"type"`
}

// NewNull receives an NullFieldSpec and returns and NullField
func NewNull(spec NullFieldSpec) *NullField {
	return &NullField{
		name: spec.Name,
	}
}
