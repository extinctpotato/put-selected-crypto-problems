BIN_DIR = ./bin/
CMD_DIR = ./cmd/

PROGS = lsbstegenc lsbstegdec visenc visdec
LIST = $(addprefix $(BIN_DIR), $(PROGS))

.PHONY: clean

all: $(LIST)

clean:
	rm -rf $(BIN_DIR)

$(BIN_DIR)%: $(CMD_DIR)%
	mkdir -p $(BIN_DIR)
	go build -o ./$@ ./$<
