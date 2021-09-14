// seen solution here https://leetcode.com/problems/happy-number/discuss/56917/My-solution-in-C(-O(1)-space-and-no-magic-math-property-involved-)
var digitSquareSum = function(n) {
    var sum = 0
    var tmp
    while (n) {
        tmp = n % 10
        sum += tmp * tmp
        n = Math.trunc(n / 10)
    }
    return sum;
}

var isHappy = function(n) {
    var slow, fast
    slow = fast = n
    do {
        slow = digitSquareSum(slow)
        fast = digitSquareSum(fast)
        fast = digitSquareSum(fast)
    } while(slow != fast)
    if (slow == 1) return 1
    else return 0
}

