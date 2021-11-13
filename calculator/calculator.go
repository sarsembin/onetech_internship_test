package calculator

type Calculator struct {
	Input  <-chan int
	Output chan<- int
}

func (c *Calculator) Start() {
	go func(c *Calculator){
		for {
			num, ok := <-c.Input
			if !ok {
				close(c.Output)
				break
			} else {
				c.Output <- num * num
			}
		}
	}(c)

}
