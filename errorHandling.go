package main

func showErr(e error) {
	if e != nil {
		panic(e)
	}
}
