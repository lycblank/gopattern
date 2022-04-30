# 访问者模式

### 使用场景

1. 对象结构比较稳定，但经常需要在此对象结构上定义新的操作；
2. 需要对一个对象结构中的对象进行很多不同的并且不相关的操作，而需要避免这些操作“污染”这些对象的类，也不希望在增加新操作时修改这些类。

### 优缺点
#### 优点
* 访问者模式 增加新的操作很容易, 只需要增加一个新的访问者即可。
* 将相关的行为, 封装到一个访问者中。

#### 缺点
* 增加新数据结构，比较困难。
* 元素变更比较困难，如为被访问的对象增加/减少一些属性, 相应的访问者也需要进行修改。


### 访问者模式的UML类图
![avatar](docs/visitor-uml.png)

### 角色介绍
* Visitor：接口或者抽象类，定义了对每个 Element 访问的行为，它的参数就是被访问的元素，它的方法个数理论上与元素的个数是一样的，因此，访问者模式要求元素的类型要稳定，如果经常添加、移除元素类，必然会导致频繁地修改 Visitor 接口，如果出现这种情况，则说明不适合使用访问者模式。
* ConcreteVisitor：具体的访问者，它需要给出对每一个元素类访问时所产生的具体行为。
* Element：元素接口或者抽象类，它定义了一个接受访问者（accept）的方法，其意义是指每一个元素都要可以被访问者访问。
* ElementA、ElementB：具体的元素类，它提供接受访问的具体实现，而这个具体的实现，通常情况下是使用访问者提供的访问该元素类的方法。
* ObjectStructure：定义当中所提到的对象结构，对象结构是一个抽象表述，它内部管理了元素集合，并且可以迭代这些元素提供访问者访问。

### 样例

#### 背景

CTO对员工进行绩效考评，针对工程师与产品经理的考评指标不一样
* 工程师主要考评代码量，并针对工程师的代码量进行评级；
* 产品经理主要考评需求量，并对产品经理的需求量进行评级；
* 输出所有员工的绩效评级报告

#### 员工定义
```go
type Staff interface {
    Name() string
    Accept(v Visitor)
}
```

#### 工程师
```go
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
```

#### 产品经理
```go
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
```

#### 访问者接口
```go
type Visitor interface {
    VisitEngineer(engineer *Engineer)
    VisitManager(manager *Manager)
}
```

#### CTO访问者实现
```go
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
```

#### 客户端实现
```go
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
```

#### 运行demo
```shell
go get -u -v github.com/lycblank/gopattern
cd visitor
go build -o main
./main
```

#### 输出
```shell
产品经理:张三 需求量:5 评级:D
产品经理:李四 需求量:15 评级:C
产品经理:王麻子 需求量:25 评级:B
工程师:小明 代码量:600 评级:D
工程师:小红 代码量:6000 评级:C
工程师:小王 代码量:60000 评级:B
```

## 如果觉得有帮助
<img src="https://github.com/lycblank/gopattern/blob/master/docs/WechatIMG48.jpeg" width="300px"/>
