
#1. 审题

给予一个二叉树的根节点，验证该树是否是二叉搜索树，（在O(n)时间内），请用你熟悉的语言写出算法；

两个关键点：1. 判断二叉搜索树的算法； 2. 复杂度要求在O（n）内；

##什么是二叉搜索树
Binary Search Tree(bst)是树内任一节点，小于其key的值都放在左子树，大于其key的值都放在右子树；

##O(n)
也就是说，n个节点，每个节点只能使用1次关键操作，就可以完成任务；这里关键操作就是遍历节点和比较操作，也就是说，每个节点只遍历一次，就能完成二叉搜索树的判断；


#2.方案选择
备选：1. 广度优先遍历； 2. 深度优先遍历

广度优先遍历树，先将访问上层节点，再访问其子节点，之后再访问其子节点的节点；
深度优先遍历是，如果当前节点左子树非空，则进入左子树；左子树为空，右子树非空，则进入右子树；

**采用广度优先遍历**，原因：


#3.伪代码

深度搜索伪代码

>     traverse_depth_first(node)
>     {
>         visit(node);
> 
>         if (node.left != NULL) {
>             traverse(node.left);
>         }
> 
>         if (node.right != NULL) {
>             traverse(node.right);
>         }
> 
>     }


广度搜索为代码

>     bool traverse_breath_first(node)
>     {
>         
>         stack.push(node)
>         while(!stack.empty()) {
>         current = stack.pop();
>         if (current.left != NULL) {
>             if (current.key > current.left.key) {
>                 stack.push(current.left)
>             }
>             else {
>                 return false
>             }  
>         }       
>         if (current.right != NULL) {
>             if (current.key >= current.right.key) {
>                 stack.push(current.right)
>             }
>             else {
>                 return false
>             }  
>         }         
>     }


#用法
运行bst

cd ./src/bst

go test