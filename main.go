package main

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")

	svc = NextLoggingService(NewMetricService(svc))

	server := NewApiServer(svc, ":3000")

	server.Run()
}
