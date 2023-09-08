.PHONY: test-voter test-vote test-poll

test-voter:
	go test -v ./voter/testcases.go
test-vote:
    go test -v ./vote/votetestcases.go
test-poll:
	go test -v ./poll/polltestcases.go