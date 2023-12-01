DAY ?= 1

.PHONY: day
day: day$(DAY) day$(DAY)/input.txt

day$(DAY):
	mkdir -p $@
	cp template/main.go $@
	cp template/main_test.go $@

day$(DAY)/input.txt:
	curl --cookie "session=$(shell cat sessiontoken.txt)" -o $@ https://adventofcode.com/2023/day/$(DAY)/input

.PHONY: run
run: runday$(DAY)

runday$(DAY):
	go run day$(DAY)/main.go

.PHONY: test
test: testday$(DAY)

testday$(DAY):
	cd day$(DAY) && go test
