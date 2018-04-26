# count-24-point
## Description
### 速算24点相信绝大多数人都玩过。就是随机给你四张牌，包括 A(1),2,3,4,5,6,7,8,9,10,J(11),Q(12),K(13)。要求只用'+','-','*','/'运算符以及括号改变运算 顺序，使得最终运算结果为24(每张牌必须且仅能用一次)。游戏很简单，但遇到无解的情况往往让人很郁闷。你的任务就是针对每一组随机产生的四张牌，判断 是否有解。我们另外规定，整个计算过程中都不能出现小数。

### 本速算24使用Go语言实现，使用递归的方式进行全列举,算法参考: https://blog.csdn.net/wangqiulin123456/article/details/8145545
### 浦发银行信用卡的速算24小游戏模式为\*A\*\*B\*\*\*C\*\*D\*，其中*为运算符，运算符的patter有限制，所以输出特定的结果
### Example
### Input 
#### 7 9 3 4
### Output
#### (9-7)\*(3\*4)
#### (9-7)\*(4\*3)
#### (9\*3)-(7-4)
#### ...


      
        
