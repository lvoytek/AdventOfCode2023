DAY ?= 1

.PHONY: day
day: day$(DAY) day$(DAY)/input.txt

day$(DAY):
	mkdir -p $@
	cp template/main.go $@
	cp template/main_test.go $@

day$(DAY)/input.txt:
	curl --cookie "session=$(shell cat sessiontoken.txt)" -o $@ https://adventofcode.com/2023/day/$(DAY)/input

