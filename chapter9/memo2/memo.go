package memo2

import "sync"

// type of Function which we will called and memoized
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func              // when we call this...
	mu    sync.Mutex        // guard cache for concurrencu
	cache map[string]result // ...will cached in this
}

// function to receive a function and return a pointer-to-cache
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		// cache no exists => call function and cache result
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}
