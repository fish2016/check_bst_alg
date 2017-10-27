#关键点

1. 判断不平衡的类型LL、LR、RR、RL;即判断不平衡的程度
2. 更新树的高度
3. 旋转算法


空树的高度为0
只有root节点树的高度为1


#旋转对树高度的影响
1. 单旋能使LL型和RR型高度减少；
2. 但单旋对LR型和RL型无影响，所以此两种类型需要先转换成LL型或RR型

算法：

1. 遍历过程中，更新height，并检测有无失衡
 
2. 如果失衡开始矫正，矫正过程中注意更新height，矫正失衡直到该子树达到平衡


用法：

cd ./src/avl
go test

> DJdeMacBook-Pro:avl dj$ go test
> 
> 1, 3, 2, 5, 110, 6, 4, 201, 205, 203, 202, 111, ========
> 
> PASS
> 
> ok  	avl	0.006s
> 
> DJdeMacBook-Pro:avl dj$ 