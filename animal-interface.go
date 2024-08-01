/*Em Golang, interfaces são tipos que especificam um conjunto de métodos, mas não implementam diretamente esses métodos. Interfaces permitem que você defina comportamentos que diferentes tipos podem compartilhar, sem impor uma hierarquia rígida. Isso facilita a criação de código flexível e reutilizável.*/
package main

import "fmt"

/*Nesse exemplo, a interface Animal exige que qualquer tipo que a implemente tenha um método Falar que retorne uma string.*/
type Animal interface {
	Falar() string
	Comer() string
}

/*
Implementando uma Interface
Para que um tipo implemente uma interface, ele deve definir todos os métodos que a interface especifica. Diferente de outras linguagens, em Golang, não é necessário declarar explicitamente que um tipo implementa uma interface. A implementação é implícita, o que significa que se um tipo define os métodos necessários, ele automaticamente implementa a interface.
*/
/*Aqui, os tipos Cachorro e Gato implementam a interface Animal porque ambos têm o método Falar.*/
type Cachorro struct{}

func (c Cachorro) Falar() string {
	return "Au Au"
}

func (c Cachorro) Comer() string {
	return "Ração"
}

type Gato struct{}

func (g Gato) Falar() string {
	return "Miau"
}

func (g Gato) Comer() string {
	return "Peixe"
}

/*
Usando Interfaces
Uma vez que você tem tipos que implementam uma interface, pode escrever funções que trabalham com essa interface:
*/
func FazerFalar(a Animal) {
	fmt.Println(a.Falar())
}

func FazerComer(a Animal) {
	fmt.Println(a.Comer())
}

func ExecuteAnimalInterface() {
	cachorro := Cachorro{}
	gato := Gato{}

	FazerFalar(cachorro)
	FazerFalar(gato)
	FazerComer(cachorro)
	FazerComer(gato)

}

// Interfaces Embutidas
// Golang também suporta a composição de interfaces, permitindo que uma interface seja composta de outras interfaces:
/*
Aqui, Atleta é uma interface que requer que qualquer tipo que a implemente também implemente as interfaces Corredor e Nadador.

type Corredor interface {
    Correr() string
}

type Nadador interface {
    Nadar() string
}

type Atleta interface {
    Corredor
    Nadador
}

*/

/*
Interface Vazia
A interface vazia, interface{}, é uma interface que não possui métodos e, por isso, é implementada por todos os tipos em Go. Ela é usada para armazenar qualquer valor:

func PrintarTudo(val interface{}) {
    fmt.Println(val)
}

// Nesse exemplo, PrintarTudo pode aceitar qualquer valor como argumento.
*/
