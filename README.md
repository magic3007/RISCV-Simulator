# RISCV-Simulator

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

- If you would like to compile your customized C source codes into RISC-V executable file, first of all, you must have <u>RISC-V Compiler Toolchain</u>. You can refer to respective online resources to install them in your environment. You can built the toolchain on your environment by yourself, or download a RISC-V GCC toolchain prebuilt under compatible environment. [Here](https://www.sifive.com/boards) is a recommended prebuilt toolchain provided by `SiFive`.

  To further compile your customized C source codes, please refer to [this section](#how-to-compile-your-customized-c-source-codes-into-risc-v-executable-file) below.

  


## How to compile your customized C source codes into RISC-V executable file?



