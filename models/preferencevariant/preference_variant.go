package preferencevariant

type PreferenceVariant struct {
	//Deprecated: this field is deprecated
	Value string `json:"value"`
	//Deprecated: this field is deprecated
	AllowOverrideInInterface bool `json:"allow_override_in_interface"`
	//Deprecated: this field is deprecated
	AllowOverrideInApi bool   `json:"allow_override_in_api"`
	Object             string `json:"object"`
}
