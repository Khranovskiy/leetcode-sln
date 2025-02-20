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



=================================================================
# 933. Number of Recent Calls
### Откуда взял
https://algorithmics-blog.github.io/blog/number_of_recent_calls/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/number-of-recent-calls/description/
https://algorithmics-blog.github.io/blog/number_of_recent_calls/
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
type RecentCounter struct {
    calls []int
}


func Constructor() RecentCounter {
    return RecentCounter{
		calls: []int{},
	}
}


func (this *RecentCounter) Ping(t int) int {
    this.calls = append(this.calls, t)

	for this.calls[0] < t-3000 {
		this.calls = this.calls[1:]
	}

	return len(this.calls)
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
```

### links
### Затраченное время
### Оставшиеся вопросы



=================================================================
# 1372. Longest ZigZag Path in a Binary Tree
### Откуда взял
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/description/
https://algorithmics-blog.github.io/blog/longest_zigzag_path_in_a_binary_tree/
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
func longestZigZag(root *TreeNode) int {
    _, _, m := dfs(root)
	return m
}

func dfs(root *TreeNode) (int, int, int) {
    var leftPathLength, rightPathLength, leftMaxLength, rightMaxLength int
	if root == nil {
		return leftPathLength, rightPathLength, 0
	}
    if root.Left != nil{
        var childRightZigZag int
        _, childRightZigZag, leftMaxLength = dfs(root.Left)
        leftPathLength = childRightZigZag + 1
    }
    if root.Right != nil{
        var childLeftZigZag int
        childLeftZigZag, _, rightMaxLength = dfs(root.Right)
        rightPathLength = childLeftZigZag + 1
    }
    rlMax := max(leftMaxLength, rightMaxLength)
    rlPathMax := max(leftPathLength, rightPathLength)
    maxLength := max(rlMax, rlPathMax)

	return leftPathLength, rightPathLength, maxLength
}
```

### links
### Затраченное время
### Оставшиеся вопросы


https://leetcode.com/problems/move-zeroes/submissions/1510967764/



=================================================================
# 1657. Determine if Two Strings Are Close
### Откуда взял
https://leetcode.com/problems/determine-if-two-strings-are-close/description/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/determine-if-two-strings-are-close/
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
func closeStrings(word1 string, word2 string) bool {
    // operation 1: swap two existing characters abcde -> aecdb
    // operation 2: tranf every occurrrence of 1 existing
    //              character and do the same with the other char
    // aacabb (a->b) bbcbaa
    //        (b->a)

    // Example 1 abc, bca:
    //   op1   b<->c  acb
    //   op1   a<->b  bca

    // Example 2 a aa - false

    // Example 3 cabbba, abbccc
    // op1 b<->a       caabbb
    // op2 (c>b,b>c)   baaccc
    // op2 (b>a,a>b)   abbccc

    if len(word1) != len(word2) {
		return false
	}
	counter1 := make(map[rune]int)
	counter2 := make(map[rune]int)

	for _, char := range word1 {
		counter1[char]++
	}

	for _, char := range word2 {
		counter2[char]++
	}
	tempCounter := make(map[int]int)

	for k, v := range counter1 {
		v2, exist := counter2[k]
		if !exist {
			return false
		}

		tempCounter[v]++
		tempCounter[v2]--
	}

    for _, count := range tempCounter {
		if count != 0 {
			return false
		}
	}

	return true
}
```

### links
### Затраченное время
### Оставшиеся вопросы


=================================================================
# 2390. Removing Stars From a String
### Откуда взял
https://algorithmics-blog.github.io/blog/remove_stars/
### timing
start:

end:

### sources, urls
https://leetcode.com/problems/removing-stars-from-a-string/description/
https://algorithmics-blog.github.io/blog/remove_stars/
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
func removeStars(s string) string {
	stack := make([]rune, 0, len(s))
	for _, v := range []rune(s) {
		if v == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, v)
		}
	}
	return string(stack)
}
```

### links
### Затраченное время
### Оставшиеся вопросы


https://leetcode.com/problems/is-subsequence/submissions/1512959610/
https://algorithmics-blog.github.io/blog/is_subsequence/

https://leetcode.com/problems/string-compression/submissions/1512992916/
https://algorithmics-blog.github.io/blog/string_compression/


https://leetcode.com/problems/max-number-of-k-sum-pairs/
https://algorithmics-blog.github.io/blog/max_number_of_k_sum_pair/

https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/
https://algorithmics-blog.github.io/blog/maximum_number_of_vowels_in_a_substring_of_given_length/

https://leetcode.com/problems/maximum-average-subarray-i/
https://algorithmics-blog.github.io/blog/maximum_average_subarray_i/

https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
https://algorithmics-blog.github.io/blog/minimum_number_of_arrows_to_burst_balloons/


https://leetcode.com/problems/maximum-frequency-stack/submissions/1518438653/
https://algorithmics-blog.github.io/blog/maximum_frequency_stack/

https://leetcode.com/problems/minimum-flips-to-make-a-or-b-equal-to-c/submissions/1519408905/
https://algorithmics-blog.github.io/blog/minimum_flips_to_make_a_or_b_equal_c/

https://leetcode.com/problems/equal-row-and-column-pairs/submissions/1519418113/
https://algorithmics-blog.github.io/blog/equal_row_and_column_pairs/


https://leetcode.com/problems/maximum-depth-of-binary-tree/submissions/1519425002/
https://algorithmics-blog.github.io/blog/maximum_depth_of_binary_tree/

https://leetcode.com/problems/leaf-similar-trees/description/
https://algorithmics-blog.github.io/blog/leaf_similar_trees/

skip hard (Acceptance Rate 35.4%)
https://algorithmics-blog.github.io/blog/longest_valid_parentheses/


8. String to Integer (atoi) medium (Acceptance Rate 18.5%)
https://leetcode.com/problems/string-to-integer-atoi/description/
https://algorithmics-blog.github.io/blog/atoi/

интересная easy
https://leetcode.com/problems/number-of-1-bits/description/
https://algorithmics-blog.github.io/blog/number_of_1_bits/

https://leetcode.com/problems/counting-bits/description/
https://algorithmics-blog.github.io/blog/counting_bits/

(Acceptance Rate 28.7%)
https://leetcode.com/problems/design-linked-list/description/
https://algorithmics-blog.github.io/blog/design_linked_list/

skip
https://algorithmics-blog.github.io/blog/design_linked_list/

классная
https://algorithmics-blog.github.io/blog/maximum_twin_sum_of_a_linked_list/
https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/submissions/1519917254/

```go
func pairSum(head *ListNode) int {
    return pairSumStack(head)
}

func pairSumSlice(h *ListNode) int {
    l := []int{h.Val}
    for curr := h.Next; curr != nil; curr = curr.Next {
        l = append(l, curr.Val)
    }
    res := 0
    for ii := 0; ii < len(l) / 2; ii++ {
        jj := len(l) - ii - 1
        res = max(res, l[ii]+l[jj])
    }
    return res
}
func pairSumStack(h *ListNode) int {
    var reversedHead *ListNode

    curListElement, oddListElement := h, h
    for oddListElement != nil && oddListElement.Next != nil {
        oddListElement = oddListElement.Next.Next

        nextDirect := curListElement.Next

        curListElement.Next = reversedHead
        reversedHead = curListElement

        curListElement = nextDirect
    }
    res := 0
    for ; curListElement != nil; curListElement = curListElement.Next {
        res = max(res, curListElement.Val +
                  reversedHead.Val )
        reversedHead = reversedHead.Next
    }
    return res
}
```

пропущу
https://algorithmics-blog.github.io/blog/search_suggestions_system/


https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/
https://algorithmics-blog.github.io/blog/kids_with_the_greatest_number_of_candies/

meduim, (Array, Hash Table, String, Sorting)
Вместо string ключа отсортированной строки используется [26]int{}
https://leetcode.com/problems/group-anagrams/description/

https://leetcode.com/problems/top-k-frequent-elements/description/


(locked)
https://leetcode.com/problems/encode-and-decode-strings/
https://neetcode.io/problems/string-encode-and-decode


128. Longest Consecutive Sequence
https://leetcode.com/problems/longest-consecutive-sequence/
https://neetcode.io/problems/longest-consecutive-sequence

242. Valid Anagram
https://leetcode.com/problems/valid-anagram/description/
https://leetcode.com/problems/valid-anagram/solutions/2347062/golang-optimized-with-followup-o-n/


**todo** open topic  4 - Prefix Sums https://neetcode.io/courses/advanced-algorithms/4
- e Range Sum Query - Immutable https://leetcode.com/problems/range-sum-query-immutable/
- m Range Sum Query 2D Immutable https://leetcode.com/problems/range-sum-query-2d-immutable/
- e Find Pivot Index https://leetcode.com/problems/find-pivot-index/
- m Product of Array Except Self  https://leetcode.com/problems/product-of-array-except-self/
- m Subarray Sum Equals K https://leetcode.com/problems/subarray-sum-equals-k/

m 560. Subarray Sum Equals K (running prefix Sum, hash table) (44.7%)
Интересная
https://leetcode.com/problems/subarray-sum-equals-k/description/
https://neetcode.io/courses/advanced-algorithms/4

взрыв бошка (Matrix, Prefix Sum) medium
https://leetcode.com/problems/range-sum-query-2d-immutable/
https://neetcode.io/courses/advanced-algorithms/4


724. Find Pivot Index (easy)
но выглядет сложнее (2 ошибочных submission)
https://leetcode.com/problems/find-pivot-index/description/

unicode.IsLetter(ch)
unicode.IsDigit(ch)
unicode.ToLower(a)
https://leetcode.com/problems/valid-palindrome/description/

167. Two Sum II - Input Array Is Sorted (med)
https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/

https://leetcode.com/problems/3sum/description/

Интересная
https://leetcode.com/problems/trapping-rain-water/
https://neetcode.io/problems/trapping-rain-water


https://leetcode.com/problems/valid-parentheses/

medium
https://leetcode.com/problems/min-stack/

medium
https://leetcode.com/problems/evaluate-reverse-polish-notation/

stack, Dynamic Programming, Backtracking (medium)
понравилась
https://leetcode.com/problems/generate-parentheses/description/

sort, решил через стек (medium)
https://leetcode.com/problems/car-fleet/

Binary search
https://leetcode.com/problems/binary-search/

good article
https://leetcode.com/discuss/study-guide/2371234/An-opinionated-guide-to-binary-search-(comprehensive-resource-with-a-bulletproof-template)/1532153#template

good article
https://www.code-recipe.com/post/binary-search

good article
https://leetcode.com/problems/binary-search/solutions/2794222/binary-search/

binary search
https://leetcode.com/problems/search-a-2d-matrix/description/

useful tips how convert 1D index to r,c (matrix placed row-wise) or column-wise
https://leetcode.com/problems/search-a-2d-matrix/description/comments/2050872

binary search tip
I think when we write e=mid-1 then we can write s<=e BUT when we write e=mid then we need to write s<e
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/comments/1966772


153. Find Minimum in Rotated Sorted Array (middle)
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

Blind 75
https://leetcode.com/discuss/general-discussion/460599/blind-75-leetcode-questions

https://leetcode.com/problems/search-in-rotated-sorted-array

100th !
https://leetcode.com/problems/time-based-key-value-store/

интересная, Sliding window
https://leetcode.com/problems/longest-repeating-character-replacement


https://leetcode.com/problems/longest-substring-without-repeating-characters/


https://leetcode.com/problems/reorder-list/description/
```go
func reorderList(head *ListNode)  {
	c := head
	cFast := c
	middle := c
	for cFast != nil && cFast.Next != nil {
		c = c.Next
		middle = c
		cFast = cFast.Next.Next
	}

	secondHalf := reverseList(middle.Next)
	middle.Next = nil

	// merge
	first, second := head, secondHalf
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " -> ")
		head = head.Next
	}
	fmt.Println("nil")
}
```

интересная
https://leetcode.com/problems/remove-nth-node-from-end-of-list
```go
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		p1, p2 = head, head
		prev   *ListNode
	)
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	for p1 != nil {
		p1 = p1.Next
		prev = p2
		p2 = p2.Next
	}
	if prev == nil {
		return head.Next
	}
	prev.Next = p2.Next
	return head
}
```

Blind 75 Animated Playlist
https://www.youtube.com/playlist?list=PLHm8nzcbp3_19DiTlDg8QYvR-hN5jzPCp

Every Leetcode Pattern You Need To Know
https://www.blog.codeinmotion.io/p/leetcode-patterns


hashmap метод, можно оптимальнее, просто пройтись один раз и занести в мапу
второй раз уже проставить указатели на новые
Есть решение с перемешиванем старых и новых копией узлов
https://leetcode.com/problems/copy-list-with-random-pointer
```go
func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node)
	var get func(orig *Node) *Node = nil

	get = func(orig *Node) *Node {
		if orig == nil {
			return nil
		}
		if clone, ok := m[orig]; ok {
			return clone
		}
		clone := &Node{}
		m[orig] = clone
		clone.Val = orig.Val
		clone.Next = get(orig.Next)
		clone.Random = get(orig.Random)

		fmt.Println(m[orig])
		return clone
	}
	newHead := get(head)
	return newHead
}
```

```go
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	curr := head
	for curr != nil {
		newNode := &Node{
			Val:  curr.Val,
			Next: curr.Next,
		}
		curr.Next = newNode
		curr = newNode.Next
	}

	curr = head
	for curr != nil {
		if curr.Random != nil {
			newCurr := curr.Next
			newCurr.Random = curr.Random.Next
		}
		// jump to the next original
		curr = curr.Next.Next
	}

	curr = head
	newHead := head.Next
	copyCurr := newHead
	for curr != nil {
		// set to the next original
		curr.Next = curr.Next.Next
		if copyCurr.Next != nil {
			copyCurr.Next = copyCurr.Next.Next
		}
		curr = curr.Next
		copyCurr = copyCurr.Next
	}

	return newHead
}
```

https://leetcode.com/problems/add-two-numbers


287. Find the Duplicate Number
Array, Two Pointers (medium)
Floyd's Tortoise and Hare Algorithm: Finding a Cycle in a Linked List

https://dev.to/alisabaj/floyd-s-tortoise-and-hare-algorithm-finding-a-cycle-in-a-linked-list-39af
https://leetcode.com/problems/find-the-duplicate-number/
```go
func findDuplicate(nums []int) int {
    slow, fast := nums[0], nums[0]

    // Phase 1: Finding the intersection point
    for {
        slow = nums[slow]
        fast = nums[nums[fast]]
        if slow == fast {
            break
        }
    }

    // Phase 2: Finding the entrance to the cycle
    slow = nums[0]
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }

    return slow
}
```

container/list used for doubly linked list
https://leetcode.com/problems/lru-cache

Min Heap with container/heap package
https://leetcode.com/problems/merge-k-sorted-lists


Привет!
Моя любимая эта:
https://leetcode.com/problems/summary-ranges/
только я даю что массив не отсортирован и если чувак решил быстро, то спрашиваю а как решать, если массив не отсортирован и результат тоже не должен быть отсортирован, т.е. для 0,1,2,4,5,7 можно засчитать [4 -> 5, 7, 0 -> 2].
Задача проверяет умеет ли человек думать инвариантами или он постоянно рассматривает какие-то "тесты" в голове и на их основе дописывает. Мне нужны те, что на инвариантах умеют думать.
Дальше спрашиваю любую из:
https://leetcode.com/problems/one-edit-distance/
https://leetcode.com/problems/group-anagrams/
https://leetcode.com/problems/line-reflection/

disjoint_set
https://leetcode.com/problems/number-of-provinces
https://algorithmics-blog.github.io/blog/number_of_provinces/
https://algorithmics-blog.github.io/blog/disjoint_set/


Trees, easy
https://leetcode.com/problems/invert-binary-tree

Trees, easy
https://leetcode.com/problems/diameter-of-binary-tree/
https://leetcode.com/problems/diameter-of-binary-tree/solutions/3241250/simplest-golang-solution-using-tree-height

Trees, easy тяжело зашла
https://leetcode.com/problems/balanced-binary-tree

Trees, easy
https://leetcode.com/problems/same-tree

Trees, easy
https://leetcode.com/problems/subtree-of-another-tree/


Trees, medium (короткое решение)
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree

trees medium (but easy),
второе решение интереснее
https://leetcode.com/problems/binary-tree-level-order-traversal
