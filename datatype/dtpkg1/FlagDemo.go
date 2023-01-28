package dtpkg1

import (
	"flag"
	"fmt"
	"os"
)

func TestArgs()  {
	args:= os.Args
	for i, arg := range args {
		fmt.Println(i,arg)
	}
}

func TestFlag() {
	name:=flag.String("name","defaultName", "input your name")
	age:=flag.Int("age",22,"input your age")

	flag.Parse()
	fmt.Println(*name,*age)

	fmt.Printf("-------------------------------------------------------\n")

	fmt.Printf("there are %d non-flag input param\n",flag.NArg())
	args:=flag.Args()
	for i,param := range args{
		fmt.Printf("#%d    :%s\n",i,param)
	}
}

type MyValue struct {
	value string
	prefix string
	suffix string
}

func (myValue *MyValue)String()string  {
		fmt.Println("我被调用了1")
		return "default string method return"
}
func (myValue *MyValue)Set(val string) error{
	fmt.Println("我被调用了2")
	myValue.value = val
	return nil
}


func TestFlagValue() {
	myValue :=MyValue{
		prefix: "pre",
		suffix: "suf",
	}
	flag.Var(&myValue,"Value","input your value")
	flag.Parse()
	fmt.Println(myValue)
}