package main

import "github.com/johnwr-response/celeritas"

type application struct {
	App *celeritas.Celeritas
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
