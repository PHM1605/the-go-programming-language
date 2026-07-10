package main

import "fmt"

var deposits = make(chan int)           // send amount to deposit
var balances = make(chan int)           // receive balance info
var withdrawals = make(chan withdrawal) // store withdraw money request
// NEW: type of withdraw request: "amount" and "where to send result"
type withdrawal struct {
	amount int
	result chan bool
}

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func Withdraw(amount int) bool {
	// result of withdrawal (ok or not) will be pumped to this channel
	result := make(chan bool)
	// send request to "teller()" via a channel
	withdrawals <- withdrawal{
		amount,
		result,
	}
	// capture withdrawal result
	return <-result
}

// this is the ONLY person can make changes to "balance"
func teller() {
	var balance int // "balance" is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			// done above; send balance info to output channel
		case w := <-withdrawals:
			if balance >= w.amount {
				balance -= w.amount
				w.result <- true
			} else {
				w.result <- false
			}
		}
	}
}

// NOTE: this function already runs automatically before main()
func init() {
	go teller()
}

func main() {
	Deposit(100)

	ok := Withdraw(30)
	fmt.Println(ok)        // true
	fmt.Println(Balance()) // 70

	ok = Withdraw(100)     // will not enough
	fmt.Println(ok)        // false
	fmt.Println(Balance()) // still 70
}
