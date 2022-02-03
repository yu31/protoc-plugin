package protodefaults

type Defaults interface {
	SetDefaults()
}

func CallDefaultsIfExists(candidate interface{}) {
	if candidate == nil {
		return
	}
	if dt, ok := candidate.(Defaults); ok {
		dt.SetDefaults()
	}
}
