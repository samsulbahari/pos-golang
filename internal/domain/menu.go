package domain

type Menu struct {
	ID   int
	Name string
}

type Submenu struct {
	Name string
}

type Result struct {
	ID      int
	Name    string
	Submenu []Submenu
}
