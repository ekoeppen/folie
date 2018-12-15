package main

func telnetInit() {
}

func telnetReset(enterBoot bool) {
	serialSend <- []byte("COLD\n")
}
