package visitor

// 经理
type Manager struct {
	name           string
	RequirementNum int32
}

func NewManager(name string) *Manager {
	return &Manager{
		name: name,
	}
}

func (m *Manager) Name() string {
	return m.name
}

func (m *Manager) Accept(v Visitor) {
	v.VisitManager(m)
}
