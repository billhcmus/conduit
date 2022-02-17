package server

type Option interface {
	apply(h *Http)
}

type optionFunc func(*Http)

func (f optionFunc) apply(h *Http) {
	f(h)
}

func Option1(value int) Option {
	return optionFunc(func(h *Http) {
		h.value1 = value
	})
}
