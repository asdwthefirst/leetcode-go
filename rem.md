# 数据结构
## 树
### 递归，
11，12，13，14  
[404. Sum of Left Leaves (Easy)](datastructure/tree/recursion/404.go)  
[687. Longest Univalue Path (Easy)](datastructure/tree/recursion/687.go)  
[337. House Robber III (Medium)](datastructure/tree/recursion/337.go)  
[671. Second Minimum Node In a Binary Tree (Easy)](datastructure/tree/recursion/671.go)  

### 前中后序  
[中序](datastructure/tree/traversing/94.go)  
[后序](datastructure/tree/traversing/145.go)  
[前序](datastructure/tree/traversing/144.go)  
### bst
2，3，5，8，10（一点印象都无）  
[230. Kth Smallest Element in a BST (Medium)](datastructure/tree/BST/230.go)  
[538. Convert BST to Greater Tree (Easy)](datastructure/tree/BST/538.go)  
[236. Lowest Common Ancestor of a Binary Tree (Medium)](datastructure/tree/BST/236.go)  
[653. Two Sum IV - Input is a BST (Easy)](datastructure/tree/BST/653.go)  
[501. Find Mode in Binary Search Tree (Easy)](datastructure/tree/BST/501.go)  
### trie
2  
[677. Map Sum Pairs (Medium)](datastructure/tree/Trie/677.go)  
## 栈与队列
2，3（最小栈的思想，完全没印象），5（单调栈），6  
[225. Implement Stack using Queues (Easy)](datastructure/stackqueue/225.go)  
[155. Min Stack (Easy)](datastructure/stackqueue/155.go)  
[739. Daily Temperatures (Medium)](datastructure/stackqueue/739.go)  
[503. Next Greater Element II (Medium)](datastructure/stackqueue/503.go)  

## 哈希
3（其实好像很简单，但还是在hash留一题），  
[594. Longest Harmonious Subsequence (Easy)](datastructure/hashset/594.go)
## 数组矩阵
6，7（毫无印象，时间换空间），8（太特化了想不起来），11  
[645. Set Mismatch (Easy)](datastructure/arraymatrix/645.go)一个数组元素在 [1, n] 之间，其中一个数被替换为另一个数，找出重复的数和丢失的数  
[287. Find the Duplicate Number (Medium)](datastructure/arraymatrix/287.go)找出数组中重复的数，数组值在 [1, n] 之间  
[667. Beautiful Arrangement II (Medium)](datastructure/arraymatrix/667.go)数组相邻差值的个数  
[565. Array Nesting (Medium)](datastructure/arraymatrix/565.go)嵌套数组  

## 图
全看，是dfs的内容
684，好复杂，我自己想不出来要强记的题  
[785. Is Graph Bipartite? (Medium)](datastructure/graph/785.go)判断是否为二分图  
[207. Course Schedule (Medium)](datastructure/graph/207.go)  
[210. Course Schedule II (Medium)](datastructure/graph/210.go)  
[684. Redundant Connection (Medium)](datastructure/graph/684.go)冗余连接并差集  




# 算法
## 排序
堆1，容易忘记。[215. Kth Largest Element in an Array (Medium)](algorithm/sort/215.go)0624成功凭空做出  
荷兰国旗1  [75. Sort Colors (Medium)](algorithm/sort/75.go)0624直接做出  
## 贪心
4，[406. Queue Reconstruction by Height(Medium)](algorithm/greed/406.go)靠记忆的  
9，[665. Non-decreasing Array (Easy)](algorithm/greed/665.go)细节处理差了一点，但是也是调调能通的  
10，[53. Maximum Subarray (Easy)](algorithm/greed/53.go)还是有一点忘记  
## 分治
都要，我完全忘记分治是怎么治的了  
1，[241. Different Ways to Add Parentheses (Medium)](algorithm/divideandconquer/241.go)  
2，[95. Unique Binary Search Trees II (Medium)](algorithm/divideandconquer/95.go)  
我完全想起来分治是怎么回事了，its easy，枚举所有情况
## 搜索
各看一题
### backtracking
5，[46. Permutations (Medium)](algorithm/search/backtracking/46.go) [leetcode](https://leetcode.cn/problems/palindrome-partitioning/description/)排列  
6，[47. Permutations II (Medium)](algorithm/search/backtracking/47.go)含有相同元素排列  
7，[77. Combinations (Medium)](algorithm/search/backtracking/77.go)组合  
8，[39. Combination Sum (Medium)](algorithm/search/backtracking/39.go)组合求和  
13，[131. Palindrome Partitioning (Medium)](algorithm/search/backtracking/131.go)分割字符串使得每个部分都是回文数  
### BFS
1，[1091. Shortest Path in Binary Matrix(Medium)](algorithm/search/BFS/1091.go)计算在网格中从原点到特定点的最短路径长度
### DFS
5，[417. Pacific Atlantic Water Flow (Medium)](algorithm/search/DFS/417.go)
## dp
也是各看一题
### 斐波那契
[213. House Robber II (Medium)](algorithm/dp/fibonacci/213.go)环形街区抢劫
### 矩阵路径
[62. Unique Paths (Medium)](algorithm/dp/matrixpath/62.go)矩阵的总路径数
### 数组区间
[413. Arithmetic Slices (Medium)](algorithm/dp/arrayrange/413.go)数组中等差递增子区间的个数
### 分割整数
[279. Perfect Squares(Medium)](algorithm/dp/cutint/279.go)
### 最长递增子序列
1，3
[300. Longest Increasing Subsequence (Medium)](algorithm/dp/longestincreasingsubsequence/300.go)最长递增子序列
[376. Wiggle Subsequence (Medium)](algorithm/dp/longestincreasingsubsequence/376.go)最长摆动子序列
### 最长公共子序列
[1143. Longest Common Subsequence](algorithm/dp/longestcommonsubsequence/1143.go)最长公共子序列
### 01背包，
4，5，7
[322. Coin Change (Medium)](algorithm/dp/0-1bag/322.go)找零钱的最少硬币数  
[518. Coin Change 2 (Medium)](algorithm/dp/0-1bag/518.go)找零钱的硬币数组合  
[377. Combination Sum IV (Medium)](algorithm/dp/0-1bag/377.go)组合总和  
### 股票交易，
4
[188. Best Time to Buy and Sell Stock IV (Hard)](algorithm/dp/stocktrade/188.go)只能进行 k 次的股票交易  
### 字符串编辑，
1，2，3
[583. Delete Operation for Two Strings (Medium)](algorithm/dp/stringedit/583.go)删除两个字符串的字符使它们相等  
[72. Edit Distance (Hard)](algorithm/dp/stringedit/72.go)编辑距离  
[650. 2 Keys Keyboard (Medium)](algorithm/dp/stringedit/650.go)复制粘贴字符  


# 其他代码
## 设计模式
1. 工厂模式的三种区别容易忘记，但是就是  
* 简单工厂：一个确切工厂生产一种产品接口。  
* 工厂模式：一个工厂接口生产一种产品接口。  
* 抽象工厂模式：一个工厂接口生产多种产品接口。  
2. 适配器模式，就是用一个Adapter包裹Adaptee，Adapter实现需要实现的接口。
3. 代理模式和装饰器模式很像，（其实感觉它们没有区别！）区别：
* 装饰器模式强调的是增强自身，在被装饰之后你能够在被增强的类上使用增强后的功能。增强后你还是你，只不过能力更强了而已；代理模式强调要让别人帮你去做一些本身与你业务没有太多关系的职责（记录日志、设置缓存）。代理模式是为了实现对象的控制，因为被代理的对象往往难以直接获得或者是其内部不想暴露出来。
* 装饰模式是以对客户端透明的方式扩展对象的功能，是继承方案的一个替代方案；代理模式则是给一个对象提供一个代理对象，并由代理对象来控制对原有对象的引用；
* 装饰模式是为装饰的对象增强功能；而代理模式对代理的对象施加控制，但不对对象本身的功能进行增强；
* 另说：技术是一样的，代理和特性增强两个应用场景使两者有区别。比如decorator可以多层装饰，代理一般就一层
* 代理注重控制,控制不易二次扩展；装饰注重增强，且通过上层装饰进行增强，易于扩展。装饰是特殊的代理，特殊的地方就在于其代理角色proxy 是一个抽象的Decorator，所以才易于扩展
4. 模板方法好绕。。。[代理模式](golearn/designpattern/template.go)