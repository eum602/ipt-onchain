package main

func fanIn(i1, i2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-i1
		}
	}()

	go func() {
		for {
			c <- <-i2
		}
	}()

	return c
}
