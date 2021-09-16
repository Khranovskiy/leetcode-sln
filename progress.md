
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

# template

## name
### timing
start:

end:

### sources, urls
### tests
```
```
### Constraints
### Brute force solution 
```
```
### complexity
space - O( )
runtime - O( )
modification input data = yes|no
### code in pseudo language
```
```
### code 
```python
```
### links
### Затраченное время
