package matematica 

// Resumao: Mainuscula para acesso fora do arquivo (export do js),
// seja para metodos, structs, etc
func Sum[T int | float64](a, b T) T{
    return a + b
}
