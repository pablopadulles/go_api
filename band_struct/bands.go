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

