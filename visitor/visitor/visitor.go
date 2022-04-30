package visitor

import (
	"fmt"
	"strings"
)

type Visitor interface {
	VisitEngineer(engineer *Engineer)
	VisitManager(manager *Manager)
}

type CTOVisitor struct {
	performances []string
}

func (cv *CTOVisitor) VisitEngineer(engineer *Engineer) {
	level := "D"
	switch {
	case engineer.CodeNum >= 0 && engineer.CodeNum < 1000:
		level = "D"
	case engineer.CodeNum >= 1000 && engineer.CodeNum < 10000:
		level = "C"
	case engineer.CodeNum >= 10000 && engineer.CodeNum < 100000:
		level = "B"
	case engineer.CodeNum > 100000:
		level = "A"
	}
	cv.performances = append(cv.performances, fmt.Sprintf("工程师:%s 代码量:%d 评级:%s",
		engineer.Name(), engineer.CodeNum, level))
}

func (cv *CTOVisitor) VisitManager(manager *Manager) {
	level := "D"
	switch {
	case manager.RequirementNum >= 0 && manager.RequirementNum < 10:
		level = "D"
	case manager.RequirementNum >= 10 && manager.RequirementNum < 20:
		level = "C"
	case manager.RequirementNum >= 20 && manager.RequirementNum < 30:
		level = "B"
	case manager.RequirementNum > 30:
		level = "A"
	}
	cv.performances = append(cv.performances, fmt.Sprintf("产品经理:%s 需求量:%d 评级:%s",
		manager.Name(), manager.RequirementNum, level))
}

func (cv *CTOVisitor) GenReport() string {
	return strings.Join(cv.performances, "\n")
}
