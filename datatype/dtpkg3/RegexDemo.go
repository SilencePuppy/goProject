package dtpkg3

import (
	"fmt"
	"os"
	"regexp"
	strconv "strconv"
)

func RegDemo1() {
	words:=[...]string{"Seven","even","Maven","Amen","eleven"}
	for _, word := range words {
		matchString, err := regexp.MatchString(".even", word)
		if err != nil {
			return
		}
		if matchString {
			fmt.Printf("%s matches\n",word)
		}else {
			fmt.Printf("%s not matches\n",word)
		}
	}

}

func RegDemo2() {
	words:=[...]string{"Seven","even","Maven","Amen","eleven"}
	re:=regexp.MustCompile(".even")
	for _, word := range words {
		found := re.MatchString(word)
		if found {
			fmt.Printf("%s matches\n",word)
		}else{
			fmt.Printf("%s not matches\n",word)
		}
	}
}
func RegDemo3() {
	var content = `Foxes are omnivorous mammals belonging to several genera
of the family Canidae. Foxes have a flattened skull, upright triangular ears,
a pointed, slightly upturned snout, and a long bushy tail. Foxes live on every
continent except Antarctica. By far the most common and widespread species of
fox is the red fox.`
	compile := regexp.MustCompile("(?i)fox(es)?")
	founds := compile.FindAllString(content, 1)
	fmt.Printf("%q\n",founds)
	if founds==nil {
		fmt.Printf("no match found\n")
		os.Exit(1)
	}
	for _, found := range founds {
		fmt.Printf("%s\n",found)
	}
}

func RegDemo4() {
	var content = `Foxes are omnivorous mammals belonging to several genera
of the family Canidae. Foxes have a flattened skull, upright triangular ears,
a pointed, slightly upturned snout, and a long bushy tail. Foxes live on every
continent except Antarctica. By far the most common and widespread species of
fox is the red fox.`
	compile := regexp.MustCompile("(?i)fox(es)?")
	indexes := compile.FindAllStringIndex(content, -1)
	fmt.Printf("%v\n",indexes)
	if indexes==nil {
		fmt.Printf("no match found\n")
		os.Exit(1)
	}
	for _, found := range indexes {
		fmt.Printf("%v\n",found)
	}
}

func RegDemo5()  {
	numStr :="22, 13,3,  3,4,5,6,77,3"
	compile := regexp.MustCompile(",\\s*")
	splitArray := compile.Split(numStr, -1)
	fmt.Printf("%v\n",splitArray)
	var sum int
	for _, s := range splitArray {
		num, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		sum+=num
	}
	fmt.Println("sum:",sum)
}

func RegDemo6()  {
	websites := [...]string{"webcode.me", "zetcode.com", "freebsd.org", "netbsd.org"}
	compile := regexp.MustCompile("(\\w+)\\.(\\w+)")
	for _, website := range websites {
		submatches := compile.FindStringSubmatch(website)
		for _,sub := range submatches {
			fmt.Println(sub)
		}
		fmt.Println("--------------------")
	}
}

func RegDemo()  {

}
