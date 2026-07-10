package memo3

// type of Function which we will called and memoized
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// result-with-readyflag
type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

// entry fills its data
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	// flag "ready"
	close(e.ready)
}

// entry waits until it has data => put data to Response Channel
func (e *entry) deliver(response chan<- result) {
	<-e.ready // hanging here
	response <- e.res
}

// NOTE: request will have the URL we must get; and channel to send Response to
type request struct {
	key      string // URL
	response chan<- result
}

// NEW: Memo now no cache; but a SERVER with a channel to receive Requests
type Memo struct {
	requests chan request
}

// NEW: memo is now a SERVER
func (memo *Memo) server(f Func) {
	// here is our cache
	cache := make(map[string]*entry)
	// server hangs here to wait to process Requests
	// RULE OF THUMB: server must create a goroutine for each request
	for req := range memo.requests {
		e := cache[req.key]
		// call that URL for 1st time
		if e == nil {
			e = &entry{ready: make(chan struct{})} // entry contains value-from-URL and ready-flag
			cache[req.key] = e
			go e.call(f, req.key) // fill data to entry
		}
		// return data to Client
		go e.deliver(req.response)
	}
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	// create a response channel to receive the result
	response := make(chan result)
	// create a Request (with URL needed and where to send Response), send it to Requests channel
	memo.requests <- request{key, response}

	// get result
	res := <-response
	return res.value, res.err
}

// function to receive a function and return a pointer-to-Memoize (which has the Requests channel)
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f) // NOTE: here is a background goroutine to handle Requests
	return memo
}
