package main

import (
	"countries-states-cities-mongo/bootstrap"
	"os"
	"os/signal"
)

func main() {
	v, err := bootstrap.SetValues()
	if err != nil {
		panic(err)
	}
	app, err := App(v)
	if err != nil {
		panic(err)
	}
	if err = app.Run(); err != nil {
		panic(err)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
}
