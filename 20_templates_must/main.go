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
	t := template.Must(template.New("CursoTemplate").Parse("Nome: {{.Nome}} - Carga Horária: {{.CargaHoraria}}h"))
	t.Execute(os.Stdout, curso)
}
