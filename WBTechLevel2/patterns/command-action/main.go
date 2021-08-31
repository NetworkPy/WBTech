package main

func main() {
	// tv := &tv{}
	pc := &pc{}

	onCommand := &onCommand{
		device: pc,
	}

	offCommand := &offCommand{
		device: pc,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
