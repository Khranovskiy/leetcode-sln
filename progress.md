
https://leetcode.com/problems/contains-duplicate/

217-contains-duplicate.py
217-contains-duplicate.go

https://play.golang.org/p/2JxWUXuAnDH


## 268. Missing Number

# sources, urls
https://leetcode.com/problems/missing-number/submissions/
268-missing-number
https://www.youtube.com/watch?v=nTt3929ik8U&list=PLyHj6yBbnkUgC6T9RpPxEBUv3My9lPd7m&index=3

```py
class Solution:
    def missingNumber(self, nums: List[int]) -> int:
        n = len(nums)

        return n * (n+1) // 2 - sum(nums)
```

```golang
func sum(array []int) int {
 r := 0  
 for _, v := range array {  
  r += v  
 }  
 return r  
}

func missingNumber(nums []int) int {
    n := len(nums)

    return int(n * (n+1)) / 2 - sum(nums)
}
```

## 448. Find All Numbers Disappeared in an Array

# sources, urls
https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

### Brute force solution with extra space

```
setNums = set(nums)
result = int[]
for i in [1,n] {
    if i not in setNums {
        result.add(i)
    }
}
return result
```
# complexity
space - O(n)
runtime - O(n)
modification input data = no

# tests
```
 t([1,1]) => [2]
 t([1]) => []
 t([4,3,2,7,8,2,3,1]) => [5,6]
 t([]) => [] -- invalid test case by constratints
```

# Solution without extra space
# итерироваться по массиву
position - [0, n -1]
value = nums[position]
делать swap? типа отсортировать?
[1,2,3,4,5,6,7,8] all integers in range 1,n
[4,3,2,7,8,2,3,1] 
[7,3,2,4,8,2,3,1] 4-7
[3,3,2,4,8,2,7,1] 7-3
[2,3,3,4,8,2,7,1] 3-2
[3,2,3,4,8,2,7,1] 2-3
не куда не приводит

# а если так 
cursor = [0, n - 1]
position = nums[cursor] - 1 // correct position
nums[cursor] <-> nums[position]
если значение разное менять, если одинаковое двигать cursor

cursor = 1

[3,2,3,4,8,2,7,1] 2=2 двигаем курсор
cursor = 2 он же индекс
[0,1,2,3,4,5,6,7] index
[3,2,3,4,8,2,7,1] value ~ nums[cursor] ~ 3
nums[cursor] == nums[position]
position - там где должен быть элемент в [1,2,3,4,5,6,7,8]
двигаем cursor
cursor <- 3
[3,2,3,4,8,2,7,1]
4 ка - она на своем месте
nums[cursor:3] = nums[position:3]
cursor <- 4
[3,2,3,4,8,2,7,1]
nums[cursor:4] ~ 8 != nums[position:7] ~ 1
swap

дальше новая итерация
nums[cursor:4] это value:1
position = value - 1
[3,2,3,4,1,2,7,8]
nums[cursor:4] ~ 1 != nums[position:0] ~ 3
swap

новая итерация
[1,2,3,4,3,2,7,8]
nums[cursor:4] это 3 
position = nums[cursor] - 1 -> 2
nums[cursor:4] ~ 3 == nums[position:2] ~ 3
значение равны делаем сдвиг cursor

cursor <- 5


новая итерация
[1,2,3,4,3,2,7,8]
nums[cursor:5] это 2
position = nums[cursor] - 1 -> 1
nums[cursor:5] ~ 2 == nums[position:1] -> 2
делаем сдвиг cursor

cursor <- 6
новая итерация
[1,2,3,4,3,2,7,8]
nums[cursor:6] это 7
position 6
nums[cursor:6] ~ 7 == nums[position:6] -> 7
он на своем месте делаем сдвиг cursor

cursor <- 7
новая итерация
[1,2,3,4,3,2,7,8]
nums[cursor:7] это 8
position 7
nums[cursor:7] ~ 8 == nums[position:7] -> 8
сдвиг cursor 
cursor <- 8

cursor == n
получившися массив
[1,2,3,4,3,2,7,8]
нужно пройтись по массиву и найти значения (индекс+1), стоящие не на своем месте
или же пройтись по значениям от 1 до n включительно и проверить что значение есть в массиве


## complexity
space - O(1)
runtime - O(n)
modification input data = yes

## code in pseudo language
cursor = 0
n = len(nums)
while cursor < n
    value = nums[cursor]
    position = value - 1
    
    if nums[cursor] != nums[position]
        # swap
        nums[cursor], nums[position] = nums[position], nums[cursor]
    else
        cursor++

result = []
for i in range(n) # [0..len(nums)-1]
    val = i + 1
    if val != nums[i]
        result.add(val)

## code in python variant with using additional space

```python
def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
    setNums = set(nums)
    result = []
    for i in range(1, len(nums) + 1): # [1..n]
        if i not in setNums:
            result.append(i)

    return result
```

## code in python without useing any additional space 
```python
def findDisappearedNumbers2(self, nums: List[int]) -> List[int]:
    cursor = 0
    n = len(nums)
    while cursor < n:
        value = nums[cursor]
        position = value - 1

        if nums[cursor] != nums[position]: 
            nums[cursor], nums[position] = nums[position], nums[cursor]
        else:
            cursor += 1

    result = []
    for i in range(n):
        val = i + 1
        if val != nums[i]:
            result.append(val)

    return result
```

## links
https://www.digitalocean.com/community/tutorials/how-to-do-math-in-go-with-operators

https://docs.python.org/3/library/stdtypes.html#set

https://docs.python.org/3/howto/sorting.html

## Затраченное время
160

## 136-single-number
### start
2021-09-16T15:47:46+03:00
Thu Sep 16 15:47:24 2021
### end
2021-09-16T15:55:50+03:00
Thu Sep 16 15:55:46 2021

### sources, urls
https://leetcode.com/problems/single-number/
### tests
```
f([2,2,1]) = 1
f([4,1,2,1,2]) = 4
f([1]) = 1
```
### Constraints
1. -3 * 10^4 <= nums[i] <= 3 * 10^4 -- 30000 
possible int16 (Range: -32768 through 32767.) will be sufficient
### Brute force solution 
```

```
### complexity
space - O(1)
runtime - O(n)
modification input data = no
## code in pseudo language
```
single = 0
for n in nums
    single = single XOR n

return single
```
## code 
```python
class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        single = 0
        for n in nums:
            single ^= n

        return single
```
## links
## Затраченное время
8


## 70-climbing-stairs
### timing
start:
Thu Sep 16 16:01:20 2021
2021-09-16T16:01:24+03:00
end:
Thu Sep 16 16:29:34 2021
2021-09-16T16:29:39+03:00
### sources, urls
https://leetcode.com/problems/climbing-stairs/
### tests
```
f(3) = 3
f(2) = 2
f(1) = 1
f(4) = 5
f(5) = 8
```
### Constraints
- 1 <= n <= 45
### Brute force solution 
```
 f(n) = f(n-1) + f(n-2)
```
### complexity
space - O(n)
runtime - O(n)
modification input data = no
## code in pseudo language
```
func climbStairs(n)
    dp = [0] * (n + 1) 
    //n =  1, 2, 3, 4, 5
    // [0, 1, 2, 3, 5, 8]
    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    for i in range(3, n+1)
        dp[i] = dp[i-2] + dp[i - 1]
    return dp[n]
```
ooops 
1. Забыл проверить на n = 2, n = 1, добавил условие if n < 3 return n

### code 
```python
def climbStairs(self, n: int) -> int:
    if n < 3:
        return n
    dp = [0] * (n + 1)
    dp[0] = 0
    dp[1] = 1
    dp[2] = 2
    for i in range(3, n + 1):
        dp[i] = dp[i-2] + dp[i-1]
    return dp[n]
```
### code without additional memory
```python
def climbStairs(self, n: int) -> int:
    if n == 1:
        return 1
    n1 = 1
    n2 = 2
    for i in range(3, n + 1):
        n1, n2 = n2, n1 + n2
    return n2
```
## links
https://www.geeksforgeeks.org/python-initialize-empty-array-of-given-length/
## Затраченное время
28


## 121-best-time-to-buy-and-sell-stock
### timing
start:
Thu Sep 16 16:58:17 2021
2021-09-16T16:58:21+03:00

2021-09-16T18:30:57+03:00

2021-09-16T18:40:32+03:00
end:
2021-09-16T17:05:06+03:00

2021-09-16T18:39:37+03:00

2021-09-16T18:55:02+03:00
### sources, urls
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
### Constraints
- 1 <= prices.length <= 10^5
- 0 <= prices[i] <= 10^4

### tests
```
f([7,1,5,3,6,4]) = 5
f([7,6,4,3,1]) = 0
f([1]) = 0
f([0]) = 0
```
### code in pseudo language
```
func maxProfit(prices) int:
    maxProfit = 0
    bestLowPrice = prices[0]

    for price in prices: # i [1.. n-1]
        maxProfit = max(maxProfit, price - bestLowPrice)
        bestLowPrice = min(bestLowPrice, price)

    return maxProfit
```
### complexity
space - O(1)
runtime - O(n)
modification input data = no

### code 
```python
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        maxProfit = 0
        bestLowPrice = prices[0]

        for price in prices[1:]:
            maxProfit = max(maxProfit, price - bestLowPrice)
            bestLowPrice = min(bestLowPrice, price)
        return maxProfit
```

```golang
func maxProfit(prices []int) int {
    maxProfit := 0
    bestLowPrice := prices[0]

    for _, price := range prices[1:] {
        maxProfit = max(maxProfit, price - bestLowPrice)
        bestLowPrice = min(bestLowPrice, price)
    }

    return maxProfit
}
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
func maxOf(vars ...int) int {
    max := vars[0]
    for _, i := range vars {
        if max < i {
            max = i
        }
    }
    return max
}
```

### links
https://python-reference.readthedocs.io/en/latest/docs/brackets/slicing.html
https://go.dev/blog/slices-intro

### Затраченное время
20
+15 минут на go

## 53-maximum-subarray
### timing
start:
?
end:

### sources, urls
https://leetcode.com/problems/maximum-subarray/
### Constraints
```
1 <= nums.length <= 3 * 10^4
-10^5 <= nums[i] <= 10^5
```
### tests
```
f([-2,1,-3,4,-1,2,1,-5,4]) = 6
f([5,4,-1,7,8]) = 23
f([1]) = 1
```
### complexity
space - O(1)
runtime - O(n)
modification input data = no

### Main idea
Берём cur_sum в него постоянно складываем числа, но если получилось что если прибавать к cur_sum текущее число и получилось хуже, чем если бы начали отчет от текущего числа, где мы сейчас, лучше мы с него и начнём, на каждой итерации переписываем max_sum
### code in pseudo language
```
```

### code

```python
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        max_sum = nums[0]
        cur_sum = nums[0]
        
        for i in range(1, len(nums)):
            num = nums[i]
            cur_sum = max(cur_sum + num, num)
            max_sum = max(max_sum, cur_sum)
            
        return max_sum
```
```golang
package main

import "fmt"

func main() {
    fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
func maxSubArray(nums []int) int {
    max_sum := nums[0]
    cur_sum := nums[0]

    for _, num := range nums[1:] {
        cur_sum = max(cur_sum+num, num)
        max_sum = max(cur_sum, max_sum)
    }
    return max_sum
}
func max(a, b int) int {
    if a >= b {
        return a
    }
    return b
}
```
### links
https://gobyexample.com/arrays
### Затраченное время


## 303 range-sum-query-immutable
### timing
start:
2021-09-16T20:42:03+03:00
end:
2021-09-16T21:05:20+03:00
### sources, urls
https://leetcode.com/problems/range-sum-query-immutable/

### Constraints
```
1 <= nums.length <= 10^4
-10^5 <= nums[i] <= 10^5
0 <= left <= right < nums.length
At most 10^4 calls will be made to sumRange
```

### tests
```
//[-2, 0, 3, -5, 2, -1], [0,2] =  1
//[-2, 0, 3, -5, 2, -1], [2,5] = -1 
//[-2, 0, 3, -5, 2, -1], [0,5] = -3
```

### Main idea
inp  [-2, 0, 3, -5, 2, -1]
sums [-2, 0, 1, -4, -2, -3]
sums[x] - сумма всех элементов до элемента с индексом x включая элемент
включая! - если он нужен то надо предыдущий брать

//нам нужен элемент nums[left], поэтому sums[left] - sums[left-1] = будет как раз элемент nums[left]
return this.sums[right] - this.sums[left - 1]

### complexity
space - O(n)
runtime - O(1) on query operation
modification input data = no
### code in pseudo language
```
```
### code 
```python
class NumArray:
    def __init__(self, nums: List[int]):
        sums = []
        current_sum = 0
        for num in nums:
            current_sum += num
            sums.append(current_sum)
            
        self.sums = sums
        print(sums)

    def sumRange(self, left: int, right: int) -> int:
        if left == 0:
            return self.sums[right]
        return self.sums[right] - self.sums[left - 1]

```

```golang
type NumArray struct {
    sums []int
}

func Constructor(nums []int) NumArray {
    sums := make([]int, len(nums), len(nums))
    
    current_sum := 0
    for i, num := range nums {
        current_sum += num
        sums[i] = current_sum
    }
    // sums[0] = nums[0]
    // for i := 1; i < len(nums); i++ {
    //     sums[i] = sums[i-1] + nums[i]
    // }
    return NumArray{ sums }
}

func (this *NumArray) SumRange(left int, right int) int {
    if left == 0 {
        return this.sums[right]
    }
    return this.sums[right] - this.sums[left - 1]
}
```
### links
### Затраченное время
23



## 338-counting-bits
### timing
start:
Thu Sep 16 21:50:50 2021
2021-09-16T21:50:43+03:00

end:
Thu Sep 16 21:54:19 2021
2021-09-16T21:54:22+03:00

Thu Sep 16 22:46:05 2021
2021-09-16T22:46:08+03:00

### sources, urls
https://leetcode.com/problems/counting-bits/
### Constraints
```
0 <= n <= 10^5
```
### tests
```
f(0) = [0]
f(2) = [0, 1,  1]
f(3) = [0, 1,  1,  2]
f(4) = [0, 1, 1, 2, 1]
f(5) = [0, 1, 1, 2, 1, 2]
f(6) = [0, 1, 1, 2, 1, 2, 2]
f(7) = [0, 1, 1, 2, 1, 2, 2, 3]
```
### Main idea
ans[i] is the number of 1s in the binary representation of i
0 <= i <= n  
f(5) = [0, 1, 1, 2, 1, 2]

if i is odd we should add 1 to count of 1s in bin repr
if i is even - add 0

Dynamic programming

Hint 2
Divide the numbers in ranges like [2-3], [4-7], [8-15] and so on. And try to generate new range from previous.
```
0  0 0 0 0  
1  0 0 0 1
2  0 0 1 0
3  0 0 1 1   0 1 1  3
4  0 1 0 0 
5  0 1 0 1 
6  0 1 1 0   0 1 1 ~ 3 = 6/2 ~ 6 >> 1
7  0 1 1 1 
8  1 0 0 0 
```
memo[x] = memo[x >> 1] + x % 2
memo[6] = memo[3] + 6 % 2
### complexity
space - O( n + 1)
runtime - O(n)
modification input data = no
### code in pseudo language
```
ans = [n+1]
ans[0] = 0
for i = 1 .. n // n included
    ans[i] = ans[i >> 1] + i % 2
return ans
```
### code 
```python
def countBits(self, n: int) -> List[int]:
    ans = [0] * (n+1)
    ans[0] = 0
    for i in range(1, n + 1):
        ans[i] = ans[i >> 1] + i % 2
    return ans
```

```golang
func countBits(n int) []int {
    ans := make([]int, n+1)
    ans[0] = 0
    for i := 1; i <= n; i++ {
        ans[i] = ans[i >> 1] + i % 2
    }
    return ans
}
```

### links
https://yourbasic.org/golang/for-loop/
### Затраченное время


## 141 linked-list-cycle
### timing
start:
Thu Sep 16 23:01:20 2021
2021-09-16T23:01:22+03:00

end:
Fri Sep 17 03:26:57 2021

### sources, urls
https://leetcode.com/problems/linked-list-cycle/
### Constraints
```
The number of the nodes in the list is in the range [0, 10^4].
-10^5 <= Node.val <= 10^5
pos is -1 or a valid index in the linked-list.
```
### tests
```
// head = [3, 2, 0, -4], pos = 1 -> yes
// head = [1, 2], pos = 0 -> yes
// head = [1], pos = -1 -> no
// head = [1, 1, 1, 1]. pos = -1 -> false
```
### Main idea
Дан односвязный список, нужно определить есть ли цикл или нет.
Решим используя алгоритм Floyd's tortoise (floyd's cycle-finding algorithm). Есть 2 указателя, медленный и быстрый. Если представить гонку 2х машин, и они едут по кругу, одна быстрее другой и они пересекутся в конечном итоге.

### complexity
space - O(1)
**runtime - O() ?????**
modification input data = no
### code in pseudo language
```
```
### code 
```python
def hasCycle(self, head: ListNode) -> bool:
    if not head:
        return False
    fast = slow = head
    while fast.next and fast.next.next:
        fast = fast.next.next
        slow = slow.next
        if fast == slow:
            return True

    return False
```

```golang
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    fast, slow := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if slow == fast {
            return true
        }
    }
    return false
}
```
### links
https://en.wikipedia.org/wiki/Cycle_detection
### Затраченное время
90
### Вопросы
Как посчитать сложность алгоритма?

## 876-middle-of-the-linked-list
### timing
start:
Fri Sep 17 03:27:12 2021
2021-09-17T03:27:14+03:00
end:

### sources, urls
https://leetcode.com/problems/middle-of-the-linked-list/
### Constraints
```
The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
```
### tests
```
head = [1,2,3,4,5]
Output: [3,4,5]

head = [1,2,3,4,5,6]
[4,5,6]
```
### Main idea
Fast & Slow Pointers

### complexity
space - O( )
runtime - O()
modification input data = yes|no
### code in pseudo language
```
func middleNode(head) {
    fast, slow = head
    while fast != nil && fast.next != nil {
        fast = fast.next.next
        slow = slow.next
    }
    return slow
}
```
### code 
```python
class Solution:
    def middleNode(self, head: Optional[ListNode]) -> Optional[ListNode]:
        fast = slow = head
        while fast and fast.next:
            fast = fast.next.next
            slow = slow.next         
        return slow
```
```golang
func middleNode(head *ListNode) *ListNode {
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow
}
```
### links
### Затраченное время
### Оставшиеся вопросы


## 206 reverse-linked-list
### timing
start:
Fri Sep 17 10:08:56 2021
2021-09-17T10:08:59+03:00
end:
Fri Sep 17 10:56:18 2021
2021-09-17T10:56:20+03:00

### sources, urls
https://leetcode.com/problems/reverse-linked-list/
### Constraints
```
The number of nodes in the list is the range [0, 5000].
-5000 <= Node.val <= 5000
```
### Main idea
Given the head of a singly linked list, reverse the list, and return the reversed list.

Modify list inplace, without creating a copy

Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?
```
[x] --> [x] --> [x]
         |
```
Lets solve iteratively.
On some step we have 
* current element
* pointer to the next
* saved element from prev step

previous -> current -> next
previous <- current <- next

### complexity
space - O(1)
runtime - O(n)
modification input data = yes
### tests
```
f([1,2,3,4,5]]) = [5,4,3,2,1] (actually it is not array)
f([1,2]) = [2,1]
f([]) = []
```
### code in pseudo language
```
func reverseList(head) {
    previous = nil
    current = head

    while current != nil {
        next = current.next
        current.next = previous
        previous = current
        current = next
    }
    return previous
}
    1 -> 2 -> 3 -> 4 -> 5
```
### code 
```python
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        prev = None
        current = head
        while current:
            next = current.next
            current.next = prev
            prev = current
            current = next
        return prev
```
```golang
func reverseList(head *ListNode) *ListNode {
    var previous *ListNode = nil
    current := head
    for current != nil {
        next := current.Next
        current.Next = previous
        previous = current
        current = next 
    }
    return previous
}
```
### links
### Затраченное время
48
### Оставшиеся вопросы


## 234-palindrome-linked-list
### timing
start:
Fri Sep 17 11:58:02 2021
2021-09-17T11:58:07+03:00

end:
Fri Sep 17 13:08:23 2021
2021-09-17T13:08:25+03:00

### sources, urls
https://leetcode.com/problems/palindrome-linked-list/
### Constraints
```
The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9
```
### Main idea
[ 1  2  2  1]
[ 1  2  
        1  2->nil]
### tests
```
```
### code in pseudo language
```
```
### complexity
space - O(1)
runtime - O(n)
modification input data = yes
### code 
```python
class Solution:
    def isPalindrome(self, head: Optional[ListNode]) -> bool:
        if not head:
            return None
        
        middle = self.middleNode(head)
        reverse = self.reverseList(middle)

        while reverse:
            if head.val != reverse.val:
                return False

            reverse = reverse.next
            head = head.next

        return True
```

```golang
func isPalindrome(head *ListNode) bool {
    if head == nil{
        return false
    }
    middle := middleNode(head)
    reversed := reverseList(middle)
    
    for reversed != nil {
        if head.Val != reversed.Val{
            return false
        }
        head = head.Next
        reversed = reversed.Next
    }
    return true
}
```

### links
### Затраченное время
70 (отвлекался)
### Оставшиеся вопросы



## 203 remove-linked-list-elements
### timing
start:
Fri Sep 17 15:38:21 2021
2021-09-17T15:38:25+03:00
end:
Fri Sep 17 16:34:29 2021
2021-09-17T16:34:32+03:00

### sources, urls
https://leetcode.com/problems/remove-linked-list-elements/
### Constraints
```
```
### Main idea
Single-linked list
to remove current element, we have keep previous

to delete first/head element -> simple return head.next
to delete second ...
tips: try to not use prev link

анализируем следующий элемент
текущий нужен для того чтобы у него обновить next
если нужно удалить следующий элемент

process head: if head value meets condition, shift head to the next element, then
```
for loop on every element
    if current.next == nil 
        break
    current.next = подходит под критерий удаления
        current.next = current.next.next
```

```
result = ListNode(0)
result.next = head
...
return result.next
```
техника фальшивого ответа, чтобы не связываться с prev, проверять head на null. 
просто идём по списку.
К head добавили перед ним элемент.
current = result
сейчас указываем на элемент до списка.
Когда будем вырезать, может получиться что у нас пустой список останется в случае f([7,7,7,7], 7), благодаря этому трюку мы этот случай корретно обработаем.

### tests
```
f([1,2,6,3,4,5,6], 6) = [1,2,3,4,5]
f([], 1) = []
f([7,7,7,7], 7) = []
```
### code in pseudo language
```
func removeElement(head, val){
    result = ListNode(0)
    result.next = head

    current = result
    while current.next:
        if current.next.val == val:
            current.next = current.next.next
        else:
            current = current.next

    return result.next
}
```
### complexity
space - O(1)
runtime - O(n)
modification input data = yes|no
### code 
```python
class Solution:
    def removeElements(self, head: Optional[ListNode], val: int) -> Optional[ListNode]:
        result = ListNode(0)
        result.next = head
        
        current = result
        while current.next:

            if current.next.val == val:
                current.next = current.next.next
            else:
                current = current.next
        
        return result.next
```

```golang
func removeElements(head *ListNode, val int) *ListNode {
    result := &ListNode{Val:0}
    result.Next = head
    
    current := result
    
    for current.Next != nil {
        if current.Next.Val == val {
            current.Next = current.Next.Next
        } else {
            current = current.Next   
        }
    }
    return result.Next
}
```

### links
### Затраченное время
60
### Оставшиеся вопросы
не проверил на пустой массив
сделал кучу ошибок
в golang версии забываю как инстанцировать структуру





## 083 remove-duplicates-from-sorted-list
### timing
start:
Fri Sep 17 16:37:22 2021
2021-09-17T16:37:25+03:00
end:
Fri Sep 17 16:56:18 2021
2021-09-17T16:56:21+03:00

### sources, urls
https://leetcode.com/problems/remove-duplicates-from-sorted-list/
### Constraints
```
 * The number of nodes in the list is in the range [0, 300].
 * -100 <= Node.val <= 100
 * The list is guaranteed to be sorted in ascending order.
```
### Main idea
Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

Добавим фейковый элемент в начало списка, чтобы он указывал на head

current.Val будем сравнивать с current.Next.Val
если равен то нужно пропустить current.Next элемент
как бы его удалить из списка.

### tests
```
f([1,1,2]) = [1,2]
f([1,1,2,3,3]) = [1,2,3]
f([]) = []
f([1]) = [1]
f([1,1,1]) = [1]
```
### code in pseudo language
```
func deleteDuplicates(head){
    current = head

    while current != nil && current.Next != nil:
        if current.Val == current.Next.Val:
            current.Next = current.Next.Next
        else:
            current = current.Next

    return result.next
}
```
### complexity
space - O(1)
runtime - O(n)
modification input data = yes
### code 
```python
def deleteDuplicates(self, head: Optional[ListNode]) -> Optional[ListNode]:
    current = head

    while current and current.next:
        if current.val == current.next.val:
            current.next = current.next.next
        else:
            current = current.next

    return head
```

```golang
func deleteDuplicates(head *ListNode) *ListNode {
    current := head
    for current != nil && current.Next != nil {
        if current.Val == current.Next.Val{
            current.Next = current.Next.Next
        }else {
            current = current.Next
        }
    }
    return head
}
```

### links
### Затраченное время
19
### Оставшиеся вопросы



## 021-merge-two-sorted-lists
### timing
start:
Fri Sep 17 17:00:00 2021
2021-09-17T17:00:02+03:00
end:
Fri Sep 17 17:15:24 2021
2021-09-17T17:15:26+03:00
### sources, urls
https://leetcode.com/problems/merge-two-sorted-lists/
### Constraints
```
```
### Main idea

### tests
```
l1 = [1,2,4], l2 = [1,3,4]
f(l1,l2) = [1,1,2,3,4,4]
f([], []) = []
f([], [0]) = [0]
```
### code in pseudo language
```
```
### complexity
space - O( )
runtime - O()
modification input data = yes|no
### code 
```python
class Solution:
    def mergeTwoLists(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        merged = ListNode(0)
        result = merged
        
        while l1 and l2:
            if l1.val < l2.val:
                merged.next = ListNode(l1.val)
                l1 = l1.next
            else:
                merged.next = ListNode(l2.val)
                l2 = l2.next
            merged = merged.next
        
        while l1:
            merged.next = ListNode(l1.val)
            l1 = l1.next
            merged = merged.next
            
        while l2:
            merged.next = ListNode(l2.val)
            l2 = l2.next
            merged = merged.next
            
        return result.next
```

```golang
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    merged := &ListNode{Val:0}
    result := merged
    
    for l1 != nil && l2 != nil{
        if l1.Val < l2.Val {
            merged.Next = &ListNode{Val:l1.Val}
            l1 = l1.Next
        } else {
            merged.Next = &ListNode{Val:l2.Val}
            l2 = l2.Next
        }
        merged = merged.Next
    }
    
    for l1 != nil {
        merged.Next = &ListNode{Val:l1.Val}
        l1 = l1.Next
        merged = merged.Next
    }
    
    for l2 != nil {
        merged.Next = &ListNode{Val:l2.Val}
        l2 = l2.Next
        merged = merged.Next
    }
    
    return result.Next
}
```

### links
### Затраченное время
15
+ 5 min on golang version
### Оставшиеся вопросы


## 920-meeting-rooms
### timing
start:
Fri Sep 17 17:34:39 2021
2021-09-17T17:34:41+03:00
end:

### sources, urls
https://leetcode.com/problems/meeting-rooms/
https://www.lintcode.com/problem/920/
### Constraints
```
si < ei
```
### Main idea

### tests
```
canAttendMeetings
f([(9,10),(1,5),(4,6)]) = false
f([(0,8),(8,10)]) = false
f([(0,30),(5,10),(15,20)]) = false
f([(5,8),(9,15)]) = true
```
### code in pseudo language
```
func canAttendMeetings(intervals){
    intervals = sorted(intervals, key=lambda x: x.start)

    for i in range(len(intervals) - 1):
        if intervals[i].end > intervals[i + 1].start:
            return False
    return True
}
```
### complexity
space - O( )
runtime - O()
modification input data = yes|no
### code 
```python
"""
Definition of Interval.
class Interval(object):
    def __init__(self, start, end):
        self.start = start
        self.end = end
"""
class Solution:
    def canAttendMeetings(self, intervals):
        intervals = sorted(intervals, key=lambda x: x.start)

        for i in range(len(intervals) - 1):
            if intervals[i].end > intervals[i + 1].start:
                return False
        return True
```

```golang
func canAttendMeetings (intervals []*Interval) bool {
    sort.SliceStable(intervals, func(i, j int) bool {
        return intervals[i].Start < intervals[j].Start
    })
    for i := 0; i < len(intervals) - 1; i++{
        if intervals[i].End > intervals[i + 1].Start{
            return false
        }
    }
    return true
}
```

### links
https://freshman.tech/snippets/go/sorting-in-go/
https://zetcode.com/golang/sort/
https://hyperskill.org/learn/step/15705
### Затраченное время
### Оставшиеся вопросы


## 704 binary-search
### timing
start:
Sat Sep 18 02:22:18 2021
2021-09-18T02:22:20+03:00
end:
Sat Sep 18 03:13:09 2021
2021-09-18T03:13:07+03:00

### sources, urls
https://leetcode.com/problems/binary-search/
### Constraints
```
 * 1 <= nums.length <= 10^4
 * -10^4 < nums[i], target < 10^4
 * All the integers in nums are unique.
 * nums is sorted in ascending order
```
### Main idea
Given an array of integers nums which is sorted in ascending order, and an integer target, write a function to search target in nums. If target exists, then return its index. Otherwise, return -1.


### tests
```
f([-1,0,3,5,9,12], 9) = 4
f([-1,0,3,5,9,12], 2) = -1
f([5], 5) = 0
f([-1,0,3,5,9,12], 2) = -1
```
### code in pseudo language
```
```
### complexity
space - O(1)
runtime - O(logn)
modification input data = yes|no
### code 
```python
class Solution:
    def searchReccursive(self, nums: List[int], target: int) -> int:
        if not nums:
            return -1
        
        middle = len(nums) // 2
        
        if nums[middle] == target:
            return middle
        if target < nums[middle]:
            return self.search(nums[:middle], target) # middle excluded
        
        index = self.search(nums[middle + 1:], target)
        if index == -1:
            return index
        return middle + 1 + index

    def search(self, nums: List[int], target: int) -> int:
        low, high = 0, len(nums)-1
    
        while low <= high:
            # median = (low + high) >>> 1
            median = (high - low) // 2 + low 
            
            if nums[median] < target:
                low = median + 1
            elif nums[median] == target:
                return median
            else:
                high = median - 1
        return -1
```

```golang
func search(nums []int, target int) int {
    low, high := 0, len(nums) - 1

    for low <= high {
        median := (low + high) >> 1
        // median = (high - low) / 2 + low 
        middle := nums[median]
        if middle == target {
            return median
        } 
        if middle < target {
            low = median + 1
        } else { 
            high = median - 1
        }
    }
    return -1
}
```

### links
https://stackoverflow.com/questions/6735259/calculating-mid-in-binary-search
https://ai.googleblog.com/2006/06/extra-extra-read-all-about-it-nearly.html
https://stackoverflow.com/questions/509211/understanding-slice-notation
### Затраченное время
50
### Оставшиеся вопросы


## 744 find-smallest-letter-greater-than-target
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/find-smallest-letter-greater-than-target
### Constraints
```
```
### Main idea

### tests
```
```
### code in pseudo language
```
```
### complexity
space - O( )
runtime - O()
modification input data = yes|no
### code 
```python
class Solution:
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        left, right = 0, len(letters)-1
    
        while left <= right:
            median = (right - left) // 2 + left 
            
            if letters[median] > target:
                right = median - 1
            else:
                left = median + 1
                
        return letters[left % len(letters)]

```

```golang
func nextGreatestLetter(nums []byte, target byte) byte {
   low, high := 0, len(nums) - 1

    for low <= high {
        median := (low + high) >> 1
        middle := nums[median]
    
        if middle <= target {
            low = median + 1
        } else { 
            high = median - 1
        }
    }
    return nums[low % len(nums)]
}

```

### links
### Затраченное время
### Оставшиеся вопросы



## 852-peak-index-in-a-mountain-array
### timing
start:
Sun Sep 19 00:47:29 2021
2021-09-19T00:47:31+03:00

end:
Sun Sep 19 01:15:26 2021
2021-09-19T01:15:29+03:00

### sources, urls
https://leetcode.com/problems/peak-index-in-a-mountain-array/
### Constraints
```
3 <= arr.length <= 10^4
0 <= arr[i] <= 10^6
arr is guaranteed to be a mountain array.
```
### Main idea
Binary Search
Follow up: Finding the O(n) is straightforward, could you find an O(log(n)) solution?

so if left-current-right
if left > current -> peek possible in left side

Т.к. по условию задачи гарантированно массив из трех элементов и что гарантированно есть пик, можно найти его по условию
arr[cur-1] < arr[cur] AND arr[cur+1] < arr[cur]

Можно пройтись по массиву с 1 элемента до предпоследнего - будет O(n) решение, а можно начать с median элемента, и используя метод бинарного поиска найти пик. 
Вначале проверить мединный элемент на условие, если да - индекс медианного элемента - ответ. Дальше нужно выбрать рассматривать левеее или правее. Так как входной массив гарантированно имеет один пик, то можно сравнить медианное значение с тем что леве и тем что правее - и обновить low,high индексы, сделать новый диапазон для поиска и так сужать диапазон для поиска и найти пик.
### tests
```
f([0,1,0]) = 1
f([0,2,1,0]) = 1
f([0,10,5,2]) = 1
f([3,4,5,1]) = 2
f([24,69,100,99,79,78,67,36,26,19]) = 2
```
### code in pseudo language
```
```
### complexity
space - O(1)
runtime - O(log(n))
modification input data = yes|no
### code 
```python
class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        low, high = 1, len(arr) - 2
        
        while low <= high:
            median = (low + high) >> 1
            
            if arr[median - 1] < arr[median] and arr[median + 1] < arr[median]:
                return median
            
            if arr[median - 1] > arr[median]:
                high = median - 1
            elif arr[median + 1] > arr[median]:
                low = median + 1

        return -1
```

```golang
func peakIndexInMountainArray(nums []int) int {
    low, high := 1, len(nums) - 2

    for low <= high {
        median := (low + high) >> 1
        
        if nums[median-1] < nums[median] && nums[median+1] < nums[median] {
            return median
        } 
        if nums[median - 1] > nums[median] {
            high = median - 1
        } else if nums[median + 1] > nums[median] { 
            low = median + 1
        }
    }
    return -1
}
```

### links
### Затраченное время
28
### Оставшиеся вопросы



## 001-two-sum (Two Pointers)
### timing
start:
Sun Sep 19 01:28:03 2021
2021-09-19T01:28:05+03:00

end:
Sun Sep 19 02:21:34 2021
2021-09-19T02:21:37+03:00

Sun Sep 19 02:38:06 2021
2021-09-19T02:38:11+03:00

### sources, urls
https://leetcode.com/problems/two-sum/
### Constraints
```
```
### Main idea

### tests
```
```
### code in pseudo language
```
```
### complexity
space - O( )
runtime - O()
modification input data = yes|no
### code 
```python
class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        otherPart = {}
        for i, num in enumerate(nums):
            diff = target - num
            if diff in otherPart:
                return [otherPart[diff], i]
            
            otherPart[num] = i
        
        return []
```

```golang
func twoSum(nums []int, target int) []int {
    other := make(map[int]int)
    
    for index, num := range nums {
        diff := target - num
        
        if otherIndex, ok := other[diff]; ok {
            return []int{otherIndex, index}
        }
        other[num] = index
    }
    return []int{}
}
```

### links
https://mypy.readthedocs.io/en/stable/cheat_sheet_py3.html
https://www.pythonpool.com/python-dynamic-array/
https://betterprogramming.pub/stop-using-range-in-your-python-for-loops-53c04593f936

https://go.dev/blog/maps

### Затраченное время
70
### Оставшиеся вопросы


## 977 squares-of-a-sorted-array (Two pointers)
### timing
start:
Sun Sep 19 02:46:43 2021
2021-09-19T02:46:46+03:00
end:

### sources, urls
https://leetcode.com/problems/squares-of-a-sorted-array/
### Constraints
```
```
### Main idea
Two pointers, easy

### tests
```
```
### code in pseudo language
```
```
### complexity
space - O( )
runtime - O()
modification input data = yes|no
### code 
```python
```

```golang

```

### links
### Затраченное время
### Оставшиеся вопросы


## 977 squares-of-a-sorted-array (Two Pointers)
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/squares-of-a-sorted-array/
### Constraints
```
```
### Main idea

bruteforce sln:
```python
def sortedSquares2(self, nums: List[int]) -> List[int]:
    result = list(map(lambda x: x*x, nums))
    return sorted(result)
```
### tests
```
```
### code in pseudo language
```
```
### complexity
space - O( )
runtime - O()
modification input data = yes|no
### code 
```python
def sortedSquares(self, nums: List[int]) -> List[int]:
    result = [0] * len(nums)
    low, high = 0, len(nums) - 1
    ri = high
    bigger = 0

    while ri >= 0:
        leftVal = nums[low] ** 2
        rightVal = nums[high] ** 2

        if leftVal > rightVal:
            bigger = leftVal
            low += 1
        elif rightVal >= leftVal:
            bigger = rightVal
            high -= 1

        result[ri] = bigger
        ri -= 1
    return result
```

```golang

```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================


=================================================================
# template

## name
### timing
start:

end:

### sources, urls
### Constraints
```
```
### Main idea

### tests
```
```
### code in pseudo language
```
```
### complexity
space - O( )
runtime - O()
modification input data = yes|no
### code 
```python
```

```golang

```

### links
### Затраченное время
### Оставшиеся вопросы