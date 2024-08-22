package hashMap

// 哈希表和我们常说的Map(键值映射)不是一个东西
// Map本质是是一个接口，而实现这个接口的类有很多，比如HashMap、TreeMap、LinkedHashMap等等
// HashMap是Map的一个实现类，而Map本身并不是一个具体的数据结构类
// 换句话说，你可以说 HashMap的get, put, remove方法的复杂度都是O(1)的，但你不能说Map接口的复杂度都是O(1)。因为如果换成其他的实现类，比如TreeMap，那么这些方法的复杂度就变成 O(logN) 了

// 哈希表基本原理: 加强版的数组
// 数组可以通过索引(非负整数)在O(1)的时间复杂度内查找到对应元素
// 哈希表是类似的，可以通过key在O(1)的时间复杂度内查找到这个key对应的value,key的类型可以是数字、字符串等多种类型

// 具体实现：哈希表的底层实现就是一个数组，它先把这个key通过一个哈希函数(hash)转化成数组里面的索引，然后增删改查操作和数组基本相同

// 几个关键概念及原理：
// 关键概念1: key是唯一的，value可以重复
// 关键概念2: 哈希函数：他的作用是把任意长度的输入(key)转化成固定长度的输出(索引)
// 哈希函数实现的几个难点及解决方案：
// 1. 如何把`key`转化成整数(通过hash函数)
// 2. 如何保证索引合法，哈希函数得到的整数是一个int类型，可能会出现负数的情况，但索引的下标是非负整数；两种方式可以做到，一种是取模，一种是补码之后的位运算，因为取模运算的性能开销比较大，所以基本上都是会使用位运算的模式
// 关键概念3: 哈希冲突：两个不同的key通过哈希函数得到了相同的索引，该怎么办
// 1. 哈希冲突无法避免，只能在算法层面妥善的处理哈希冲突的情况，原因是hash函数相当于是把一个无穷大的空间映射到一个有限的索引空间，所以必然会有不同的key映射到同一个索引上
// 2. 哈希冲突的解决方案1：拉链法(垂直纵向扩展) 数组+链表
// 3. 哈希冲突的解决方案2：线性探查法(水平横向扩展) 数组
// 关键概念4: 扩容与负载因子
// 拉链法和线性探查法虽然能解决哈希冲突的问题，但是他们会导致性能下降；
// 拉链法，根据哈希函数计算出来的索引去查询，但是查出来是一个链表，还需要遍历链表才能找到想要的value，这个过程的时间复杂度是O(K)，K是这个链表的长度
// 线性探查法，根据哈希函数计算出来的索引去查询，去这个索引位置查看，发现存储的不是要找的key，但由于线性探查法解决哈希冲突的方式，并不能确定这个key真的不存在，你必须顺着这个索引往后找，直到找到一个空的位置或者找到这个key为止，这个过程的时间复杂度也是O(K)，K为连续探查的次数
// 基于哈希值分布均匀的控制，负载因子应运而生
// 负载因子是一个哈希表装满的程度的度量，一般来说，负载因子越大，说明哈希表里面的key-value对越多，哈希冲突的概率就越大
// 负载因子的计算公式也很简单，就是 size / table.length。其中 size 是哈希表里面的 key-value 对的数量，table.length 是哈希表底层数组的容量
// 因为链表可以无限延伸，所以拉链法的负载因子可以无限大，用线性探查法实现的哈希表，负载因为不会超过1