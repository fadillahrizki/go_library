package main

type Book struct {
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Name string
	Age  string
}

var Books = []*Book{
	&Book{Title: "Math", Author: &Author{Name: "Rizky Fadillah", Age: "17"}},
	&Book{Title: "Physics", Author: &Author{Name: "Rizky Fadillah", Age: "17"}},
	&Book{Title: "Programming", Author: &Author{Name: "Rizky Fadillah", Age: "17"}},
}
