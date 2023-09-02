CURRENT=$(wildcard *.go) $(wildcard */*.go) $(wildcard */*/*.go) $(wildcard */*/*/*.go)
CONFIG=$(wildcard ../../shared/config/*.go)
INIT=$(wildcard ../../shared/initialize/*.go)
UTILS=$(wildcard ../../shared/utils/*.go) $(wildcard ../../shared/utils/*/*.go)
TAR=$(patsubst $(shell dirname ${PWD})/%,%,$(shell pwd))
run: $(TAR)
	tmux rename-window $(TAR)
	./$(TAR) $$[RANDOM % 55535 + 10000]

$(TAR): $(CURRENT) $(CONFIG) $(INIT) $(UTILS)
	echo $(CURRENT)
	grep -n -R TODO
	go build
