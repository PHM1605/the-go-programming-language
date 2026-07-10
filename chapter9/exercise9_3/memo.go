package exercise93

import "errors"

// type of Function which we will called and memoized
// NEW: "done" channel to cancel the call
type Func func(key string, done <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// result-with-readyflag
type entry struct {
	res       result
	ready     chan struct{} // closed when res is ready
	cancelled bool          // mark when receive cancel signal
}

// entry fills its data
// NEW: "done" channel to cancel call early
func (e *entry) call(f Func, key string, done <-chan struct{}) {
	e.res.value, e.res.err = f(key, done)

	select {
	case <-done:
		e.cancelled = true
	default: // do nothing as usual, "cancelled" flag still "false"
	}

	// flag "ready" to return data
	close(e.ready)
}

// entry waits until it has data => put data to Response Channel
// NEW: "done" channel to cancel call early
func (e *entry) deliver(
	response chan<- result,
	done <-chan struct{},
) {
	select { // hanging to wait for result OR cancel response
	case <-e.ready:
		response <- e.res
	case <-done:
	}
}

// request will have the URL we must get; and channel to send Response to
// NEW: a "done" channel to cancel that request anytime from outside
type request struct {
	key      string // URL
	response chan<- result
	done     <-chan struct{}
}

// Memo now no cache; but a SERVER with a channel to receive Requests
type Memo struct {
	requests chan request
}

// memo is now a SERVER
func (memo *Memo) server(f Func) {
	// here is our cache
	cache := make(map[string]*entry)
	// server hangs here to wait to process Requests
	// RULE OF THUMB: server must create a goroutine for each request
	for req := range memo.requests {
		e := cache[req.key]

		// NEW: if an entry exists, check if it's "ready"
		// - if "ready", check if it's cancelled flag was turned ON or still OFF
		if e != nil {
			select {
			case <-e.ready:
				if e.cancelled {
					e = nil
				}
			default: // normal - do nothing
			}
		}

		// call that URL for 1st time
		if e == nil {
			e = &entry{ready: make(chan struct{})} // entry contains value-from-URL and ready-flag
			cache[req.key] = e
			go e.call(f, req.key, req.done) // fill data to entry
		}
		// return data to Client
		go e.deliver(req.response, req.done)
	}
}

func (memo *Memo) Close() {
	close(memo.requests)
}

// NEW: "done" channel to cancel this request
func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	// create a response channel to receive the result
	response := make(chan result)
	// create a Request (with URL needed and where to send Response), send it to Requests channel
	memo.requests <- request{key, response, done}

	// NEW: get result OR cancel
	select {
	case res := <-response:
		return res.value, res.err

	case <-done:
		return nil, errors.New("cancelled")
	}
}

// function to receive a function and return a pointer-to-Memoize (which has the Requests channel)
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f) // NOTE: here is a background goroutine to handle Requests
	return memo
}
