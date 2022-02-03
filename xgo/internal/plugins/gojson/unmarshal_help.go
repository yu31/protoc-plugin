package gojson

func (p *plugin) genVariableOneofIsStore(oneofName string) string {
	return "oneof" + oneofName + "isStore"
}

func (p *plugin) unmarshalObjectBeforeReadKey(loopLabel string) {
	p.g.P("if decoder.ObjectBeforeReadKey() { // before read object key")
	p.g.P("    break ", loopLabel)
	p.g.P("}")
}

func (p *plugin) unmarshalObjectBeforeReadValue() {
	p.g.P("decoder.ObjectBeforeReadValue() // Before read object value")
}

func (p *plugin) unmarshalObjectAfterReadValue(loopLabel string) {
	p.g.P("if decoder.ObjectAfterReadValue() { // After read object value")
	p.g.P("    break ", loopLabel)
	p.g.P("}")
}

func (p *plugin) unmarshalArrayBeforeReadValue(loopLabel string) {
	p.g.P("if decoder.ArrayBeforeReadValue() { // Before read array value.")
	p.g.P("    break ", loopLabel)
	p.g.P("}")
}

func (p *plugin) unmarshalArrayAfterReadValue(loopLabel string) {
	p.g.P("if decoder.ArrayAfterReadValue() { // After read array value.")
	p.g.P("    break ", loopLabel)
	p.g.P("}")
}
