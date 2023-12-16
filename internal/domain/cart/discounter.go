package cart

type Discounter interface {
	calculate(c Cart) Discount
}
