run: 
	# Should runner with the given arguments to run the streamer

run-runner: 
	# does the same as run
	make run

run-streamer: 
	go build clistreamer.go
	go run .\clistreamer.go -args="Title,Message 1,Message 2,Stream Delay,Run Times\nCLI Invoke1,First Msg1,Second Msg2,2,10"