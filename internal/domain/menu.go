package domain

type Menu struct {
	ID   int
	Name string
}

type Submenu struct {
	Name string
}

type ResultMenu struct {
	ID      int
	Name    string
	Submenu []Submenu
}
