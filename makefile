DIR=$(patsubst %/,%,$(wildcard cmd/*/))
TAR=$(foreach d,$(DIR),$(d)/$(notdir $(d)))
.PHONY: $(TAR) build
build: $(TAR)

$(TAR):
	bash -c "cd $(dir $@) && make $(notdir $@)"

run: build
	./run.sh $(DIR)

clean:
	rm $(TAR)
