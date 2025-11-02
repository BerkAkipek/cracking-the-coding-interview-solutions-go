package callcenter

/*
Call Center: Imagine you have a call center with three levels of employees: respondent, manager,
and director. An incoming telephone call must be first allocated to a respondent who is free. If the
respondent can't handle the call, he or she must escalate the call to a manager. If the manager is not
free or not able to handle it, then the call should be escalated to a director. Design the classes and
data structures for this problem. Implement a method dispatchCall() which assigns a call to
the first available employee.
*/

type Employee interface {
	HandleCall(call *Call)
	IsAvailable() bool
	GetRank() int
	SetAvailable(bool)
}

type EmployeeBase struct {
	isAvailable bool
	rank        int
	id          int
}

type Call struct {
	id        int
	rank      int
	completed bool
}

func (e *EmployeeBase) IsAvailable() bool { return e.isAvailable }

func (e *EmployeeBase) SetAvailable(available bool) { e.isAvailable = available }

type Respondent struct{ EmployeeBase }

type Manager struct{ EmployeeBase }

type Director struct{ EmployeeBase }

type CallManager struct {
	calls chan *Call
}

func (cm *CallManager) Start(workers []Employee) {
	for _, w := range workers {
		go cm.dispatchWorker(w)
	}
}

func (cm *CallManager) dispatchWorker(emp Employee) {
	for call := range cm.calls {
		emp.HandleCall(call)
	}
}

func (cm *CallManager) ReceiveCall(c *Call) {
	cm.calls <- c
}
