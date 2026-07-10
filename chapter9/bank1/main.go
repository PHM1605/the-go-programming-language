package main

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance info

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // "balance" is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			// send balance info to output channel
		}
	}
}

// NOTE: this function already runs automatically before main()
func init() {
	go teller()
}

func main() {
	Deposit(100)
	Deposit(50)
	println(Balance())
}
