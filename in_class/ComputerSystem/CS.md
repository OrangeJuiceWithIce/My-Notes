## flynn分类
S-Single;M-Multiple;I-Instruction;D-DataStream
SISD: 
SIMD: 
MISD:
MIMD:
## ISA(Instruction Set Architecture)
### 分类
1.CISC(Complex Instruction Set Computer):复杂指令集计算机，如x86
2.RISC(Reduced Instruction Set Computer):精简指令集计算机，如ARM

register-memory ISA:很多指令可以直接访问内存，例如x86
load-store ISA:内存必须先加载到寄存器，然后才能操作寄存器，例如RISC-V

## 功耗
### dynamic power
信号0-1的变化所需的能量
power = 活动因子 * 负载电容 * 电压平方 
#### overclocking
### static power
保持信号0或1所需的能量
power=current * voltage
## 成本
