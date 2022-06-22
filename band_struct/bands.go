package band_struct

type Member struct {
	Name     string
	LastName string
}

type Band struct {
	Name    string
	Members []Member
}

type Disc struct {
	Name string
	Type string
}

func (m Member) get_nombre_completo() string {

	res := ""
	res += m.LastName
	res += ", "
	res += m.Name
	return res

}