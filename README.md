# sRISCV-Simulator

A Simple RISC-V ISA Simulator that supports *RV64I* Base Instruction Set and *RV64M* Standard Extension, and can run in both <u>single-instruction mode</u> and <u>pipeline mode</u>. ✨

## Overview

This simulator functions in interactive mode like `gdb`. It provides following interactive operations that have functions similar with those in `gdb`. 

-  `c`: Continue running until the program comes to the end
-  `reg`: Display register information
-  `info`: Display the address of this symbol defined in ELF file
-  `si`: In single-instruction mode, the function of this operations is the same as that of `gdb`, namely running a single machine instruction. However, in pipeline mode, it means running a single pipeline stage.
-  ` x/<length><format> [address]`: Display the memory contents at a given address using the specified format.

- `status`(only in pipeline mode): Display the status of each pipeline register.

*<mark>Note</mark>: Operation `si` has different function in single-instruction mode and pipeline mode. We will explain why we have such design in the following section.*

## Dependencies

- If you would like to compile your customized C source codes into RISC-V executable file, you must have <u>RISC-V Compiler Toolchain</u> at first. You c to install them in your environment. You can built the toolchain on your environment by yourself, or download a RISC-V GCC toolchain prebuilt under compatible environment. [Here](https://www.sifive.com/boards) is a recommended prebuilt toolchain provided by `SiFive`.

  To further compile your customized C source codes, please refer to [this section](#how-to-compile-your-customized-c-source-codes-into-risc-v-executable-file) below.
  
- <u>[Go](https://golang.org/doc/install) with version 1.12 or later</u>: please refer to the official website to install Go in your environment. 

## How to compile your customized C source codes into RISC-V executable file?

First of all, ensure that you could find `riscv64-unknown-elf-gcc` and `riscv64-unknown-elf-objdump` in your `PATH` environment variable. Please refer to perspective online resources to learn how to add them into `$PATH` if you couldn’t find them.

```bash
$ riscv64-unknown-elf-gcc --version
riscv64-unknown-elf-gcc (SiFive GCC 8.3.0-2019.08.0) 8.3.0
Copyright (C) 2018 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
$ riscv64-unknown-elf-objdump --version
GNU objdump (SiFive Binutils 2.32.0-2019.08.0) 2.32
Copyright (C) 2019 Free Software Foundation, Inc.
This program is free software; you may redistribute it under the terms of
the GNU General Public License version 3 or (at your option) any later version.
This program has absolutely no warranty.
```

A [makefile](./testcases/makefile) script has be provided under directory [testcases/](./testcases). Switch to this directory and type `make` , Any file with extension `*.c` under this directory will be compiled the RISC-V executable file as well as disassembled to corresponding assembly language file.

```bash
$ cd testcases
$ make
riscv64-unknown-elf-gcc -Wa,-march=rv64im -static -Wl,--no-relax -g -o hello.out hello.c
riscv64-unknown-elf-objdump -Sd hello.out > hello.asm
```

Note that several compiling flags are necessary

- ` -Wa,-march=rv64im`: compulsorily compile the source file into RV64I executable file
- `-static`: statically linking
- `-Wl,--no-relax`: To start running from `main`, we have to forbid the compiler to leverage the global pointer to optimize

*<mark>Node</mark>: This simulator currently only support 32-bit instruction. However, some linked library functions in prebuilt toolchain use 16-bit compressed Instructions, therefore this simulator currently starts running from `main` and does not support system calls.*🤯🤯

See more info in [makefile](./testcases/makefile).

## Compilation

**Ubuntu18.04**
```bash
mkdir bin
export GOPATH=$(pwd)
go build -o bin/sim src/main.go
```

## How to configure my customized Instruction?

One sparkle point of this simulator is that it leverages <u>the design idea of data-driving</u>. In other words, by editing the table [src/action_table.csv](./src/action_table.csv) with portability, you could add, delete and modify the behavior or even the delay in each pipeline stage of any instruction, So you could design your customized instruction!😜

With little effort, the scope of supported instructions could be expanded with ease. To see how currently supported instruction functions, please refer to [src/action_table.csv](./src/action_table.csv).

After you have configured your customized instruction, you have to generate code as following, and then recompile your code as illustrated in [this previous section](#compilation).

```bash
export GOPATH=$(pwd)
go run src/action_parser.go -f src/action_table.csv -t src/action.tmpl
go run src/microaction_parser.go -f src/action_table.csv -t src/microaction.tmpl
```

