package memo3

import "sync"

// type of Function which we will called and memoized
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// NEW: result-with-readyflag
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f     Func              // when we call this...
	mu    sync.Mutex        // guard cache for concurrency
	cache map[string]*entry // ...will cached in this. NEW: map an URL string to pointer-to-result-with-lock
}

// function to receive a function and return a pointer-to-Memoize
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// cache no exists => call function and cache result
		e = &entry{ready: make(chan struct{})} // our result-with-readyflag here
		memo.cache[key] = e                    // hold place for cache inside memoize
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key) // fill cache
		close(e.ready)                       // broadcast "ready" condition
	} else {
		// if key is ready => we do nothing but unlock
		memo.mu.Unlock()
		<-e.ready // wait for "readyflag" until that value is ready to return
	}

	return e.res.value, e.res.err
}
