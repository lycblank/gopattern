package visitor

type Staff interface {
	Name() string
	Accept(v Visitor)
}
