=================================================================
## 011. container-with-most-water
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

# template

=================================================================

## 064. minimum_path_sum
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
