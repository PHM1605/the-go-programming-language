package memo1

// type of Function which we will called and memoized
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func              // when we call this...
	cache map[string]result // ...will cached in this
}

// function to receive a function and return a pointer-to-cache
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		// cache no exists => call function and cache result
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
