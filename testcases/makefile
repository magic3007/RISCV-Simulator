CC=riscv64-unknown-elf-gcc
OBJDUMP=riscv64-unknown-elf-objdump
CC_FLAG=-Wa,-march=rv64im -static -Wl,--no-relax -g
OBJDUMP_FLAG=-Sd

SRC_PATH=.


SRC=$(foreach x, $(SRC_PATH), $(wildcard $(addprefix $(x)/*,.c)))
OBJ=$(addsuffix .out, $(notdir $(basename $(SRC))))
ASM=$(addsuffix .asm, $(notdir $(basename $(SRC))))

%.out: %.c
	$(CC) $(CC_FLAG) -o $@ $<

%.asm: %.out
	$(OBJDUMP) $(OBJDUMP_FLAG) $< > $@

.PHONY: all

all: $(OBJ) $(ASM)
