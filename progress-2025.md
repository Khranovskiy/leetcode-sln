# 011. container-with-most-water
### Откуда взял
https://algorithmics-blog.github.io/blog/container_with_most_water/

### timing
start:

end:

### sources, urls
https://leetcode.com/problems/container-with-most-water/
https://algorithmics-blog.github.io/blog/container_with_most_water/

### Constraints
```
x
```

### Main idea

### Brute force solution

```go
func maxArea(height []int) int {
  res := 0
  for i := 0; i < len(height)-1; i++ {
      for j := i + 1; j < len(height); j++ {
          minHeight := height[i]
          if height[j] < minHeight {
              minHeight = height[j]
          }

          square := minHeight*(j-i)
          if square > res {
              res = square
          }
      }
  }

  return res
}
```

### test cases (normal and edge cases)
```
[1,8,6,2,5,4,8,3,7]
49
```

### code in pseudo language

```
x
```

### complexity

space - O(1)
runtime - O(n)
modification input data = no

### code

```typescript
function maxArea(height: number[]): number {
    let res = 0
    let i = 0
    let j = height.length - 1

    while (i < j) {
        const length = j - i
        let minHeight = height[i]

        if (height[j] < minHeight) {
            minHeight = height[j]
        }

        const square = length * minHeight

        if (square > res) {
            res = square
        }

        if (height[i] < height[j]) {
            i++
            continue
        }

        if (height[j] < height[i]) {
            j--
            continue
        }

        i++
        j--
    }

    return res
};
```


```go
func maxArea(h []int) int {
    var l, r = 0, len(h) - 1
    var res = int(0)

    for l < r {
        length := r - l
        minHeight := h[l]

        if h[r] < h[l]{
          minHeight = h[r]
        }

        square := length * minHeight

        if square > res {
            res = square
        }
        if h[l] < h[r]{
            l = l + 1
            continue
        }
        if h[r] < h[l]{
            r  = r - 1
            continue
        }
        r = r - 1
        l = l + 1
    }
    return res
}
```

### links
### Затраченное время
### Оставшиеся вопросы

# 064. minimum_path_sum
### Откуда взял
https://algorithmics-blog.github.io/blog/minimum_path_sum/

### timing
start:

end:

### sources, urls
https://algorithmics-blog.github.io/blog/minimum_path_sum/
https://leetcode.com/problems/minimum-path-sum/description/

### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
grid := [[1,3,1],[1,5,1],[4,2,1]]
```
output 7

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func minPathSum(g [][]int) int {
    var i,j int
    for i = 0; i < len(g); i++ {
        for j = 0; j < len(g[i]); j++{
            //  x - - -
            //  - - - -
            //  - - - -
            if i == 0 && j == 0 {
                continue
            }
            //  p c x x
            //  - - - -
            //  - - - -
            if i == 0 {
                p := g[i][j-1]
                g[i][j] = p + g[i][j]
                continue
            }
            //  p - - -
            //  c - - -
            //  x - - -
            if j == 0 {
                p := g[i-1][j]
                g[i][j] = p + g[i][j]
                continue
            }
            //  -  p1 - -
            //  p2 c  x x
            //  -  x  x x
            p2, p1 := g[i][j-1], g[i-1][j]
            minPrices := p2
            if p1 < minPrices {
                minPrices = p1
            }
            g[i][j] = minPrices + g[i][j]
        }
    }
    // i == len(g), j = len(g[i])
    return g[i-1][j-1]
}
```

### links
### Затраченное время
### Оставшиеся вопросы



# 289. Game of Life
### Откуда взял
https://algorithmics-blog.github.io/blog/game_of_life/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/game-of-life/description/
https://algorithmics-blog.github.io/blog/game_of_life/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
const (
    dead = 0
    live = 1
    toLive = 2
    toDead = 3
)
func gameOfLife(b [][]int)  {
    for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
            c := getliveNeighbors(i, j, b)

            if b[i][j] == dead {
                if c == 3 {
                    b[i][j] = toLive
                }
                continue
            }

            if c < 2 || c > 3 {
                b[i][j] = toDead
                continue
            }

            b[i][j] = live
        }
    }
    for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			if b[i][j] == toDead {
				b[i][j] = dead
			}

			if b[i][j] == toLive {
				b[i][j] = live
			}
		}
	}
}

func getliveNeighbors(i, j int, board [][]int) int {
    up, down, in := 0, 0, 0

	if i > 0 {
		up = getValue(board[i-1][j])

		if j > 0 {
			up += getValue(board[i-1][j-1])
		}

		if j < len(board[i])-1 {
			up += getValue(board[i-1][j+1])
		}
	}

	if i < len(board)-1 {
		down = getValue(board[i+1][j])

		if j > 0 {
			up += getValue(board[i+1][j-1])
		}

		if j < len(board[i])-1 {
			up += getValue(board[i+1][j+1])
		}
	}

	if j > 0 {
		in += getValue(board[i][j-1])
	}

	if j < len(board[i])-1 {
		in += getValue(board[i][j+1])
	}

	return up + down + in
}

func getValue(val int) int {
	if val == 2 {
		return 0
	}

	if val == 3 {
		return 1
	}

	return val
}
```

### links
### Затраченное время
### Оставшиеся вопросы

# 200. Number of Islands
### Откуда взял
https://algorithmics-blog.github.io/blog/number_of_islands/
https://leetcode.com/problems/number-of-islands/description/

### timing
start:

end:

### sources, urls
https://leetcode.com/problems/number-of-islands/description/?source=submission-ac
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func numIslands(g [][]byte) int {
    count := 0
    for i:=0; i < len(g); i++{
        for j:=0;j<len(g[i]);j++{
            if g[i][j] == '1' {
                count++
                recursiveDelete(g, i, j)
            }
        }
    }
    return count
}
func recursiveDelete(g [][]byte, i, j int) {
    g[i][j] = '0'
    if j < len(g[i]) - 1 && g[i][j+1] == '1' {
        recursiveDelete(g, i, j + 1)
    }
    if i < len(g) - 1 && g[i+1][j] == '1' {
        recursiveDelete(g, i + 1, j)
    }
    if j > 0 && g[i][j-1] == '1' {
        recursiveDelete(g, i, j - 1)
    }
    if i > 0 && g[i-1][j] == '1' {
        recursiveDelete(g, i - 1, j)
    }
}
```

### links
### Затраченное время
### Оставшиеся вопросы

# 198. House Robber
### Откуда взял
https://algorithmics-blog.github.io/blog/house_robber/
### timing
start:

end:

12 minutes

### sources, urls
https://leetcode.com/problems/house-robber/description/
https://algorithmics-blog.github.io/blog/house_robber/
### Constraints
```
x
```

### Main idea

The robber can arrive at the last house either from the third-to-last house or the fourth-to-last house. So, knowing the maximum amount looted for the third-to-last and fourth-to-last houses, it is easy to calculate the amount for the last house by choosing the maximum of the two and adding the value of the current house. The same principle applies to the third-to-last and fourth-to-last houses (and to the second-to-last, in case it is more profitable for the robber to stop at the second-to-last house).

```
nums = [2, 7, 9, 3, 1]
           p  p     с
```
two types of prev: fourth-to-last (pos-3) - 7 and third-to-last (pos-2) - 9;

### test cases (normal and edge cases)
```
nums = [2, 0, 0, 0, 1]
nums = [2, 0]
nums = [2]
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
        max := nums[0]
        if nums[1] > max{
            max = nums[1]
        }
        return max
	}
    nums[2] = nums[2] + nums[0]

	for i := 3; i < len(nums); i++ {
		max := nums[i-2]
		if  nums[i - 3] > max {
			max = nums[i-3]
		}
		nums[i] = nums[i] + max
	}
	return max(nums[len(nums)-1], nums[len(nums)-2])
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
```

### links
### Затраченное время
12 minutes
### Оставшиеся вопросы


# 189. Rotate Array

### Откуда взял
https://algorithmics-blog.github.io/blog/rotate_array/

### timing
start:

end:

### sources, urls
https://leetcode.com/problems/rotate-array/description/

### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O(1)
runtime - O(n)
modification input data = yes

### code

```python
x
```

```go
func rotate(n []int, k int) {
    length := len(n)
    if length < 2 {
        return
    }
    count := k % length

	revInpl(n, 0, length - 1)

    revInpl(n, 0, count - 1)
    revInpl(n, count, length - 1)
}

func revInpl(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
```

### links
### Затраченное время
10 min
### Оставшиеся вопросы


## 208. Implement Trie (Prefix Tree)
### Откуда взял
https://algorithmics-blog.github.io/blog/implement_trie_prefix_tree/

### timing
start:

end:

### sources, urls
https://leetcode.com/problems/implement-trie-prefix-tree/
https://algorithmics-blog.github.io/blog/implement_trie_prefix_tree/

### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O(1)
runtime - O(n)
modification input data = no

### code

```python
x
```


```go
type Trie struct {
    children   map[rune]*Trie
	isFullWord bool
}


func Constructor() Trie {
    return Trie{
        children: nil,//make(map[rune]*Trie),
    }
}


func (t *Trie) Insert(word string)  {
    t.insert([]rune(word))
}


func (t *Trie) Search(word string) bool {
    node, exist := t.traverse([]rune(word))
	if !exist {
		return false
	}
	return node.isFullWord
}


func (t *Trie) StartsWith(prefix string) bool {
    _, exist := t.traverse([]rune(prefix))

	return exist
}

func (t *Trie) insert(word []rune) {
    if len(word) == 0 {
        t.isFullWord = true
        return
    }
    child, exist := t.children[word[0]]
    if !exist {
        child = &Trie{
            children: nil, //make(map[rune]*Trie),
        }
        if t.children == nil{
            t.children = make(map[rune]*Trie)
        }
        t.children[word[0]] = child
    }
    child.insert(word[1:])
}

func (t *Trie) traverse(prefix []rune) (*Trie, bool){
    if len(prefix) == 0 {
        return t, true
    }
    child, exist := t.children[prefix[0]]
    if !exist{
        return nil, false
    }
    return child.traverse(prefix[1:])
}
```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================

# 36. Valid Sudoku
### Откуда взял
### timing
start:

end:

### sources, urls

### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func isValidSudoku(board [][]byte) bool {
	var (
		rows    [9][9]bool
		columns [9][9]bool
		quads   = make(map[string][9]bool, 9)
	)
	n := len(board)
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if board[row][col] == '.' {
				continue
			}
			strNum := string(board[row][col])
			num, _ := strconv.Atoi(strNum)
			pos := num - 1
			key := fmt.Sprintf("%d:%d", row/3, col/3)

			if rows[row][pos] || columns[col][pos] || quads[key][pos] {
				return false
			}
			rows[row][pos] = true
			columns[col][pos] = true

			temp := quads[key]
			temp[pos] = true
			quads[key] = temp
		}
	}
	return true
}
```

### links
### Затраченное время
### Оставшиеся вопросы
=================================================================
# 238. Product of Array Except Self
### Откуда взял
https://algorithmics-blog.github.io/blog/product_of_array_except_self/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/product-of-array-except-self/description/
https://algorithmics-blog.github.io/blog/product_of_array_except_self/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```go
func productExceptSelf2(nums []int) []int {
	leftProduct := make([]int, len(nums))
	rightProduct := make([]int, len(nums))

	leftProduct[0] = 1
	for i := 1; i < len(nums); i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	rightProduct[len(nums)-1] = 1
	for j := len(nums) - 2; j >= 0; j-- {
		rightProduct[j] = rightProduct[j+1] * nums[j+1]
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = leftProduct[i] * rightProduct[i]
	}
	return nums
}
```

### complexity

space - O(1)
runtime - O(n)
modification input data = no

### code

```python
x
```


```go
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	r := 1
	for j := len(nums) - 1; j >= 0; j-- {
		res[j] = r * res[j]
		r *= nums[j]
	}
	return res
}
```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================
# 412. Fizz Buzz
### Откуда взял
https://algorithmics-blog.github.io/blog/fizz_buzz/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/fizz-buzz/description/
https://algorithmics-blog.github.io/blog/fizz_buzz/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```go
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			res[i-1] = "FizzBuzz"
		} else if i%3 == 0 {
			res[i-1] = "Fizz"
		} else if i%5 == 0 {
			res[i-1] = "Buzz"
		} else {
			res[i-1] = strconv.Itoa(i)
		}
	}
    return res
}
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
        if i%3 == 0 {
			res[i-1] = "Fizz"
		}
        if i%5 == 0 {
			res[i-1] = res[i-1] + "Buzz"
		}
        if len(res[i-1]) == 0 {
			res[i-1] = strconv.Itoa(i)
		}
	}
    return res
}
```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================
# 2119. A Number After a Double Reversal
### Откуда взял
https://algorithmics-blog.github.io/blog/number_double_reversion/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/a-number-after-a-double-reversal/description/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func isSameAfterReversals(num int) bool {
    if num == 0 {
        return true
    }
    return !(num % 10 == 0 )
}
```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================
# 206. Reverse Linked List
### Откуда взял
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/reverse-linked-list/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func reverseList(head *ListNode) *ListNode {
    var (
        curr = head
        prev *ListNode = nil
    )
    for curr != nil {
        next := curr.Next
        curr.Next = prev

        prev = curr
        curr = next
    }
    return prev
}
```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================
# 7. Reverse Integer
### Откуда взял
https://algorithmics-blog.github.io/blog/reverse_integer/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/reverse-integer/
https://algorithmics-blog.github.io/blog/reverse_integer/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func reverse(left int) int {
	res := 0

	for left != 0 {
		right := left % 10
		left = left / 10

		if res > math.MaxInt32/10 || res == math.MaxInt32/10 && right > 7 {
			return 0
		}
		if res < math.MinInt32/10 || res == math.MinInt32/10 && right < -7 {
			return 0
		}
		res = res*10 + right
	}
	return res
}
```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================
# 735. Asteroid Collision
### Откуда взял
https://algorithmics-blog.github.io/blog/asteroid_collision/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/asteroid-collision/description/
https://algorithmics-blog.github.io/blog/asteroid_collision/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))

	for i := 0; i < len(asteroids); i++ {
		stack = appendToStack(stack, asteroids[i])
	}
	return stack
}
func appendToStack(stack []int, current int) []int {
	if len(stack) == 0 {
		return append(stack, current)
	}
	top := stack[len(stack)-1]

	if top*current > 0 || top < 0 {
		// put back to the stack
		// add the current element to the stack.
		return append(stack, current)
	}
	// top positive and curr element negative
	// curr asteroid has a smaller mass
	if top > -1*current {
		// return top back to the stack and move to the next element.
		return stack
	}

	// If the top asteroid has a smaller mass, we extract the next element from the stack without returning the top element to the stack and repeat the checks with a new pair.
	if top < -1*current {
		return appendToStack(stack[:len(stack)-1], current)
	}
	// If the elements have the same mass
	// move to the next element
	// without returning the top element to the stack.
	return stack[:len(stack)-1]
}
```

### links
### Затраченное время
20 min
### Оставшиеся вопросы

=================================================================
# 739. Daily Temperatures
### Откуда взял
https://algorithmics-blog.github.io/blog/daily_temperatures/

### timing
start:

end:

### sources, urls
https://leetcode.com/problems/daily-temperatures/description/
https://algorithmics-blog.github.io/blog/daily_temperatures/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	res := make([]int, len(temperatures)) // filled with zeros
	for j, dailyTemp := range temperatures {
		for len(stack) > 0 && dailyTemp > temperatures[stack[len(stack)-1]] {
			i := stack[len(stack)-1] // индекс предыдущего дня попрохладне
			stack = stack[:len(stack)-1]
			res[i] = j - i // j or tempIndex индекс первого теплого дня
		}
		stack = append(stack, j)
	}
	return res
}

// [73,74,75,71,69,72,76,73]

// [9 8 7 6 5 4 3 2 10]
// [8 7 6 5 4 3 2 1 0]

// [ 1 2 3 4 5 6 7 8]

// j = 0, dayTemp 1, stack []: stack is [0]
// j = 1, dayTemp 2, stack [0]: 2 > temp[ stack head aka 0] 2 > 1, stack [], res[head] = j - head
```

### links
### Затраченное время
бесконечность

### Оставшиеся вопросы
Similar problems:
- 901. Online Stock Span (M) https://leetcode.com/problems/online-stock-span/
- 496. Next Greater Element I (E) https://leetcode.com/problems/next-greater-element-i/


=================================================================
# 605. Can Place Flowers
### Откуда взял
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/can-place-flowers/description/
https://algorithmics-blog.github.io/blog/can_place_flowers/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) == 0 {
		return false
	}

	if len(flowerbed) == 1 {
		if n > 1 {
			return false
		}

		return flowerbed[0] == 0
	}
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			i++
			continue
		}

		if canBePlaced(flowerbed, i) {
			flowerbed[i] = 1
			n--
			i++
		}

		if n == 0 {
			return true
		}
	}

	return n == 0
}

func canBePlaced(flowerbed []int, i int) bool {
	if i == 0 {
		return flowerbed[i+1] == 0
	}

	if i == len(flowerbed)-1 {
		return flowerbed[i-1] == 0
	}
	return flowerbed[i-1] == 0 && flowerbed[i+1] == 0
}
```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================
# 345. Reverse Vowels of a String
### Откуда взял
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/reverse-vowels-of-a-string/
https://algorithmics-blog.github.io/blog/reverse_vowels_of_string/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func reverseVowels(s string) string {
	lp, rp := 0, len(s)-1
	runes := []rune(s)
	for lp < rp {
		if !isVowel(runes[lp]) {
			lp++
		}
		if !isVowel(runes[rp]) {
			rp--
		}
		if isVowel(runes[lp]) && isVowel(runes[rp]) {
			runes[lp], runes[rp] = runes[rp], runes[lp]
			lp++
			rp--
		}
	}
	return string(runes)
}
func isVowel(ch rune) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================


# 435. Non-overlapping Intervals
### Откуда взял
https://algorithmics-blog.github.io/blog/non_overlapping_intervals/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/non-overlapping-intervals/

### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( logn) .. 0(n)  -- thanks to sorting
runtime - O(nlogn) -- thanks to sorting
modification input data = yes|no

### code

```python
x
```


```go
func eraseOverlapIntervals(intervals [][]int) int {
	// [[1,2],[2,3],[3,4],[1,3]]
    sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
    [[1,2],[2,3],[1,3],[3,4]]
	res := 0
	k := intervals[0][1] // 2
	for idx := 1; idx < len(intervals); idx++ {
        // [1,3]
		end := intervals[idx][1]
		start := intervals[idx][0]

        //  1 >= 2
		if start >= k {
			k = end
		} else {
			res++ // 1
		}
	}
	return res
}
```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================

# 66. Plus One
### Откуда взял
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/plus-one/
https://algorithmics-blog.github.io/blog/plus_one/

### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func plusOne(digits []int) []int {
    addition := 0
    for i:=len(digits)-1; i>=0;i--{
        if i == len(digits)-1{
            addition = 1
        }
        val := digits[i] + addition
        if val > 9 {
            digits[i] = val % 10
            addition = 1
        } else {
            digits[i] = val
            addition = 0
        }
    }
    res := make([]int, 0, len(digits)+1)
    if addition != 0 {
        res = append(res, addition)
    }
    res = append(res, digits...)
    return res
}
```

### links
### Затраченное время
### Оставшиеся вопросы


=================================================================
# 151. Reverse Words in a String
### Откуда взял
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/reverse-words-in-a-string/

### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go

func reverseWords(s string) string {
	// word separated at least one space
	// return string of the words in reverse order concatenated by a signle space

	sRunes := trimSpaces([]rune(s))
	lWb := 0
	rWb := 0
	for i := 0; i < len(sRunes); i++ {
		// skip leading spaces
		if sRunes[i] == ' ' && lWb == rWb {
			lWb++
			rWb++
		}
		if sRunes[i] == ' ' && lWb != rWb {
			reverseWord(sRunes, lWb, rWb)
			rWb += 1
			lWb = rWb
			continue

		}
		if sRunes[i] != ' ' && rWb < len(sRunes) {
			if sRunes[lWb] == ' ' {
				lWb = i
			}
			rWb = i
		}
	}
	if lWb != rWb {
		reverseWord(sRunes, lWb, rWb)
	}
	reverseWord(sRunes, 0, len(sRunes)-1)

	return strings.TrimSpace(string(sRunes))
}
func reverseWord(s []rune, i, j int) {
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func trimSpaces(sRunes []rune) []rune {
	isStart := true
	j := 0
	for i := 0; i < len(sRunes); i++ {
		if sRunes[i] == ' ' && isStart {
			continue
		}

		sRunes[j] = sRunes[i]
		j++
		isStart = sRunes[i] == ' '
	}

	if sRunes[j-1] == ' ' {
		j--
	}

	return sRunes[:j]
}
```

### links
### Затраченное время
### Оставшиеся вопросы


=================================================================
# 394. Decode String
### Откуда взял
https://algorithmics-blog.github.io/blog/decode_string/
### timing
start:

end:

### sources, urls
https://algorithmics-blog.github.io/blog/decode_string/

### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func decodeString(s string) string {
    res := ""
    for i := 0; i < len(s); i++ {
        if s[i] >= 'a' && s[i] <= 'z' {
			res += string(s[i])
			continue
		}

        count := 0
		for s[i] != '[' {
			count = count*10 + int(s[i]-'0')
			i++
		}
        bracket := 1
		j := i + 1
		for bracket > 0 {
			if s[j] == '[' {
				bracket++
			}
			if s[j] == ']' {
				bracket--
			}
			j++
		}
        subStr := s[i+1:j-1]
        inner := decodeString(subStr)
		res += strings.Repeat(inner, count)
		i = j - 1
	}

	return res
}
```

### links
### Затраченное время
### Оставшиеся вопросы

=================================================================
# 649. Dota2 Senate
### Откуда взял
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/dota2-senate/submissions/1507766234/
https://algorithmics-blog.github.io/blog/dota2_senate/

### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func predictPartyVictory(senate string) string {
	rCount, dCount := 0, 0
	queue := make([]byte, 0, 2*len(senate))
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			rCount++
		} else {
			dCount++
		}

		queue = append(queue, senate[i])
	}

	for rCount > 0 && dCount > 0 {
		firstElem := queue[0]
		killCount := 1
		queue = append(queue, firstElem)
		queue = queue[1:]

		for killCount > 0 && rCount > 0 && dCount > 0 {
			if queue[0] == firstElem {
				queue = append(queue, queue[0])
				queue = queue[1:]
				killCount++
			} else {
				if queue[0] == 'R' {
					rCount--
				} else {
					dCount--
				}
				killCount--
				queue = queue[1:]
			}
		}
		fmt.Println(string(queue))
	}

	if rCount > 0 {
		return "Radiant"
	}

	return "Dire"
}

```

### links
### Затраченное время
### Оставшиеся вопросы


=================================================================
# 374. Guess Number Higher or Lower
### Откуда взял
https://algorithmics-blog.github.io/blog/guess_number_higher_or_lower/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/guess-number-higher-or-lower/description/
https://algorithmics-blog.github.io/blog/guess_number_higher_or_lower/
### Constraints
```
x
```

### Main idea

### test cases (normal and edge cases)
```
x
```

### code in pseudo language

```
x
```

### complexity

space - O( )
runtime - O()
modification input data = yes|no

### code

```python
x
```


```go
func guessNumber(n int) int {
	l, r := 0, n
	for l <= r {
		mid := l + (r-l)/2
		g := guess(mid)
		if g == 0 {
			return mid
		} else if g == 1 { // mid is lower than the picked
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	 -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *           otherwise return 0
 * func guess(num int) int;
 */
```

### links
### Затраченное время
### Оставшиеся вопросы
