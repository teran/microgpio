package fake

type FakeDriver struct {
	LowFunc    func(int) error
	HighFunc   func(int) error
	OutputFunc func(int) error
}

func (f *FakeDriver) Output(id int) error {
	return f.OutputFunc(id)
}

func (f *FakeDriver) Low(id int) error {
	return f.LowFunc(id)
}

func (f *FakeDriver) High(id int) error {
	return f.HighFunc(id)
}
