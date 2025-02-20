package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	curso := Curso{"Go", 40}
	tmp := template.New("CursoTemplate")
	tmp, _ = tmp.Parse("Nome: {{.Nome}} - Carga Horária: {{.CargaHoraria}}h")
	tmp.Execute(os.Stdout, curso)
}
