package main

func order(fuel string, full Full) {
	err := full.ful()
	if err != nil {
		return
	}
}

type Full interface {
	ful() error
}

type diesel struct {
	pumpsNum int
}

func newDiesel(pumpsNum int) Full {
	return &diesel{
		pumpsNum: pumpsNum,
	}
}

func (f *diesel) ful() error {

	return nil
}

type electr struct {
}

func newElectr(pumpsNum int) Full {
	return &electr{}
}

func (f *electr) ful() error {

	return nil
}

type benz struct {
}

func newBenz(pumpsNum int) Full {
	return &benz{}
}

func (f *benz) ful() error {

	return nil
}

func main() {
	f := "benz"
	var full Full
	fulWay := 3

	switch fulWay {
	case 1:
		full = newDiesel(1)
	case 2:
		full = newBenz(2)
	case 3:
		full = newElectr(3)
	}
	order(f, full)
}
