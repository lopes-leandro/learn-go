package controllers

//Storage Interface de armazenamento oferece
// suporte para obter e colocar um único valor
type Storage interface{
	Get() string
	Put(string)
}

//MemStorage implementa armazenamento
type MemStorage struct{
	value string
}

//Get obtenha nosso valor na memória
func (m *MemStorage) Get() string  {
	return m.value
}

//Put coloque nosso valor na memória
func (m *MemStorage) Put(s string)  {
	m.value = s
}