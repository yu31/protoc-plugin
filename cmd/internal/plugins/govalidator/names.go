package govalidator

const (
	// TODO: supported user-defined prefix.
	prefix = "_xxx_xxx_Validator_"
)

func (p *plugin) buildVariableName(fieldInfo *FieldInfo, tagName string) string {
	x := prefix + p.message.GoIdent.GoName + "_"
	//if fieldInfo.Field.Parent.Desc.IsMapEntry() {
	//	x += string(fieldInfo.Field.Parent.Desc.Name()) + "_"
	//}
	x += tagName + "_"
	if fieldInfo.IsCheckIf {
		x += fieldInfo.Parent.Field.GoName + "_By_" + fieldInfo.Field.GoName
	} else {
		x += fieldInfo.Field.GoName
	}
	if fieldInfo.Field.Parent.Desc.IsMapEntry() {
		x += "_" + string(fieldInfo.Field.Parent.Desc.Name())
	}
	return x
}

func (p *plugin) buildVariableNameForTagIn(fieldInfo *FieldInfo) string {
	return p.buildVariableName(fieldInfo, "In")
}

func (p *plugin) buildVariableNameForTagNotIn(fieldInfo *FieldInfo) string {
	return p.buildVariableName(fieldInfo, "NotIn")
}

func (p *plugin) buildVariableNameForTagRegex(fieldInfo *FieldInfo) string {
	return p.buildVariableName(fieldInfo, "Regex")
}

func (p *plugin) buildVariableNameForTagInEnums(fieldInfo *FieldInfo) string {
	return p.buildVariableName(fieldInfo, "InEnums")
}

func (p *plugin) buildMethodNameForFieldValidate(fieldInfo *FieldInfo) string {
	return prefix + "Validate_" + fieldInfo.Name
}

func (p *plugin) buildMethodNameForFieldCheckIf(fieldInfo *FieldInfo) string {
	return prefix + "CheckIf_" + fieldInfo.Name
}
