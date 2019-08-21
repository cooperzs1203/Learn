package object

func Student(name string) student {
	return student{
		Name:name,
	}
}

type student struct {
	Name string
	id string
}

func (s *student) SetId(id string) {
	s.id = id
}

func (s *student) GetId() string {
	return s.id
}