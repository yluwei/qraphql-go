package models

type Hello struct {
	Id   int
	Name string
}

func (h *Hello) Query(id int, name string) ([]Hello, error) {
	hellos := make([]Hello, 0)
	hello := Hello{
		Id:   id,
		Name: name,
	}
	hellos = append(hellos, hello)
	return hellos, nil
}
