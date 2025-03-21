# 数据库系统

---

## Key
superkey(超码)：能够唯一识别一条记录的属性或属性集
candidate key(候选码)：能够唯一识别一条记录的最小属性集
primary(主码)：在候选码中人为选择
foreign key(外键)：一个关系的外键指向另一个关系的主键，前者称为refrencing relation(参照表，从表)，后者称为referenced relation(被参照表，主表)

## Relational Algebra
1.slect:满足特定要求的row 符号:σ条件(relation)
2.project(投影):一张表,选择特定column,注意删去重复的row 符号:π列名(relation)
3.Union:表1和表2的并集 符号:∪
4.Set Difference:表1和表2的差集 符号:−
5.Cartesian Product:笛卡尔积,表1的每一row与表2的每一row组合 符合:x
6.rename:相当于创造一个新表,原表的名字变为新表的名字  符号:p'name'(attribute1,attribute2,...)(Expression)
7.Natural join:将同名属性值相同的列叠加起来
Theta join:先笛卡尔积，再select 符号:r '横着的沙漏'条件 s=σ条件(r x s)
## extended Relation Algebra Operations
instersection:
insertion:交,符号:∩
assignment:赋值,符号:←
division:除,p=r÷s,p拥有s没有的属性,pxs是r的子集 符号:÷

>arity:关系的属性个数
compatible relation:关系的属性的个数和类型相同
## Join
Inner Join(Natrual Join) 两个表都不存在的属性值舍去
Left Out Join ,左边的每一行都得有结果，不存在的值为null
Right Out Join,右边的每一行都得有结果，不存在的值为null
Full Outer Join
