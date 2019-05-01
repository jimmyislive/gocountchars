.PHONY : test

test:
	@cd cmd/gocountchars && go test -bench . -short -v
