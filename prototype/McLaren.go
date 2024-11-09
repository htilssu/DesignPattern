package prototype

type McLaren struct {
	Name  string
	Speed int
}

func (m McLaren) Clone() interface{} {
	return &McLaren{
		Name:  m.Name,
		Speed: m.Speed,
	}
}

func NewMcLaren() *McLaren {
	return &McLaren{}
}
