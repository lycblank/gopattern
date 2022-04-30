package visitor

// 工程师
type Engineer struct {
	name    string
	CodeNum int32
}

func NewEngineer(name string) *Engineer {
	return &Engineer{
		name: name,
	}
}

func (e *Engineer) Name() string {
	return e.name
}

func (e *Engineer) Accept(v Visitor) {
	v.VisitEngineer(e)
}
