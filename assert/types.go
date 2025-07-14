package assert

type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type container interface {
	~string | ~[]byte | ~[]rune | ~[]any |
		~[]int | ~[]int8 | ~[]int16 | ~[]int64 |
		~[]uint | ~[]uint16 | ~[]uint32 | ~[]uint64 | ~[]uintptr
}

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}
