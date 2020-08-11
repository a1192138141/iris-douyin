package rabbitmq

import "fmt"

type Test struct {
	
}

func (self *Test) Push() string  {
	return  "1111"
}

func (self *Test) Consume(data []byte) bool  {
	fmt.Println(data)
	return  true
}
