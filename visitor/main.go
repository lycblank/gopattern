package main

import (
	"fmt"
	"github.com/lycblank/gopattern/visitor/visitor"
)

type StaffManager struct {
	staffs []visitor.Staff
}

func (sm *StaffManager) Accept(vs ...visitor.Visitor) {
	for _, staff := range sm.staffs {
		for _, v := range vs {
			staff.Accept(v)
		}
	}
}

func main() {
	sm := &StaffManager{}
	m1 := visitor.NewManager("张三")
	m1.RequirementNum = 5
	m2 := visitor.NewManager("李四")
	m2.RequirementNum = 15
	m3 := visitor.NewManager("王麻子")
	m3.RequirementNum = 25
	sm.staffs = append(sm.staffs, m1,m2,m3)

	e1 := visitor.NewEngineer("小明")
	e1.CodeNum = 600
	e2 := visitor.NewEngineer("小红")
	e2.CodeNum = 6000
	e3 := visitor.NewEngineer("小王")
	e3.CodeNum = 60000
	sm.staffs = append(sm.staffs, e1,e2,e3)

	cv := &visitor.CTOVisitor{}
	sm.Accept(cv)

	fmt.Println(cv.GenReport())
}
