#!../geist

// a simple example geist script:

begin := time.Now()

fmt.Println("Welcome to Geist!")

fmt.Printf("elapsed time is: %v\n", time.Since(begin))
