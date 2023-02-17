package langs

type PropNValue struct {
	name  string
	value []interface{}
}

type Prop struct {
	name     string
	proptype string
}

func newProp(name string, proptype string) Prop {
	return Prop{
		name:     name,
		proptype: proptype,
	}
}
