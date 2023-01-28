package dtpkg2

import "fmt"

type Humaner interface {
	SayHi()
}

type Personer interface {
	Humaner
	Sing()
}

type Student struct {
	name string
}

func (s *Student) SayHi()  {
	fmt.Println("Hi I am ",s.name)
}
func (s *Student) Sing() {
	fmt.Println("Hi sing a song",s.name)
}

func Demo() {
	s:=&Student{"lxb"}
	var h Humaner
	h =s
	h.SayHi()
	var p Personer
	p = s
	p.Sing()
	p.SayHi()
	fmt.Println("===============")
	h =p
	h.SayHi()
	var(
		a int
		b float64
	)
	a =22
	b= 333.3
	a=int(b)
	fmt.Println(a)
}


