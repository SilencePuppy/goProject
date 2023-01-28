package dtpkg4

import "encoding/json"

type IT struct {
	company string
	subjects []string
	isOk bool
	price float64
}

func demo()  {
	s:=IT{"lxb",[]string{"java","go"},true,33.3}
	marshal, err := json.Marshal(s)
	if err != nil {
		return 
	}
}
