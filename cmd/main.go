package main

import "kafka-by-momin/app"

func main() {
	application := app.NewApplication()
	application.Init()
	application.Run()
	application.Wait()
	application.Cleanup()
}
