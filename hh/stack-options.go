package hh

type StackOptions struct {
	Size int
}

type StackOption func(*StackOptions)

func StackSize(size int) StackOption{
	return func(options *StackOptions) {
		options.Size = size
	}
}