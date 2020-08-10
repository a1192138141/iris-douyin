package socket

type RealizeInterface interface {
	Tets() string
}

type Realize struct{}

func Test() string {
	return "1"
}
