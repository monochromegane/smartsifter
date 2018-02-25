package smartsifter

type Categorizer interface {
	Index(x []int) int
	Size(idx int) int
}

type noOpCategorizer struct {
}

func (c noOpCategorizer) Index(x []int) int {
	return int(x[0])
}

func (c noOpCategorizer) Size(idx int) int {
	return 1
}
