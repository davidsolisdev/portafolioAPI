package main

func main() {
	err := App().Listen(":3005")
	if err != nil {
		panic(err.Error())
	}
}
