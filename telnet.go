package main

func telnetInit() {
}

func telnetReset(enterBoot bool) {
	serialSend <- []byte("soft-reset\n")
}
