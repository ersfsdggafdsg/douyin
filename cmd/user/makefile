.PHONY: TODO
CURRENT=$(wildcard *.go) $(wildcard */*.go) $(wildcard */*/*.go) $(wildcard */*/*/*.go)
CONFIG=$(wildcard ../../shared/config/*.go)
INIT=$(wildcard ../../shared/initialize/*.go)
UTILS=$(wildcard ../../shared/utils/*.go) $(wildcard ../../shared/utils/*/*.go)
TAR=$(patsubst $(shell dirname ${PWD})/%,%,$(shell pwd))
run: TODO $(TAR)
	./$(TAR) $$[RANDOM % 55535 + 10000]

TODO:
	grep -n -R TODO
	tmux rename-window $(TAR)

$(TAR): $(CURRENT) $(CONFIG) $(INIT) $(UTILS)
	echo $(CURRENT)
	go build
