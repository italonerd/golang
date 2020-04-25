package estrutura

type ObjetoA struct {
	Num  int
	Desc string
}

func (obj ObjetoA) New(numero int, descricao string) ObjetoA {
	return ObjetoA{numero, descricao}
}

func (obj ObjetoA) GetDesc() string {
	return obj.Desc
}

func (obj *ObjetoA) SetDesc(descricao string) string {
	obj.Desc = descricao
	return obj.Desc
}
