//show declaring OMIT
// Declaring and initializing.
var c chan int
c = make(chan int)
// or
c := make(chan int) // HL
//end show declaring OMIT

//show sending OMIT
// Sending on a channel
c <- 1 // HL
//end show sending OMIT

//show receiving OMIT
// Receiving from a channel
// The "arrow" indicates the direction of data flow.
value = <-c // HL
//end show receiving OMIT
