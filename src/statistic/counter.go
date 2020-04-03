package statistic

type Counter struct {
	Count uint
}

func (c *Counter) Tick(){
	c.Count += 1
}

type RatioCounter struct {
	Numerator, Denominator Counter
}

func (c *RatioCounter) TickNumerator(){
	c.Numerator.Tick()
}

func (c *RatioCounter) TickDenominator(){
	c.Denominator.Tick()
}