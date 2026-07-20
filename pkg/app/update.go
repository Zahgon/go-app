package app

type updateManager struct {
	pending []map[Composer]int
}

func (m *updateManager) Add(c Composer, v int) { _ = "STUB: not implemented"; return }

func (m *updateManager) Done(v Composer) { _ = "STUB: not implemented"; return }

func (m *updateManager) UpdateForEach(do func(Composer)) { _ = "STUB: not implemented"; return }
