package modelos

//Esto es herencia
type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar() { m.respirando = true }

func (m *Mujer) Pensar() { m.pensando = true }

func (m *Mujer) Comer() { m.comiendo = true }

func (m *Mujer) Sexo() string { return "Mujer" }

func (m *Mujer) EstaVivo() bool { return true }
