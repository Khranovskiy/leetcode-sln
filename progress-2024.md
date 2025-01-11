https://leetcode.com/problems/merge-sorted-array/
https://leetcode.com/problems/merge-sorted-array/post-solution/?submissionId=964310581
```csharp
public class Solution {
    public void Merge(int[] a, int m, int[] b, int n) {
        int it1=0;
        int it2=0;

        int[] inpA = new int[m];
        Array.Copy(a, inpA, m);
        var result = a;

        while (it1<m && it2<n)
        {
            if(inpA[it1]<b[it2]){
                result[it1+it2]=inpA[it1];
                it1++;
            }
            else{
                result[it1+it2]=b[it2];
                it2++;
            }
        }

        while(it1<m){
            result[it1+it2]=inpA[it1];
            it1++;
        }

        while(it2<n){
            result[it1+it2]=b[it2];
            it2++;
        }

        return ;
    }
}
```


top-k-frequent-elements
https://leetcode.com/problems/top-k-frequent-elements/description/
https://leetcode.com/problems/top-k-frequent-elements/post-solution/?submissionId=1037672520

```javascript
const topKFrequent2 = (numbers, k) => {
    const hashmap = {};
    numbers.forEach( (number ) => {
        hashmap[number] = hashmap[number] ? hashmap[number] + 1 : 1
    })
    const uniqNumbers = [];

     for (const number in hashmap) {
        uniqNumbers.push(Number(number));
    }

    const sortedUniqNumbers = [...uniqNumbers].sort((a, b) => hashmap[b] - hashmap[a]);

    return sortedUniqNumbers.slice(0, k);
};


const topKFrequent = (nums, k) => {
    // Create an empty hash map to store the frequency of each element in the array
    const hm = {};

    // Create an empty array to store the elements based on their frequency.
    const freq = Array.from({length:nums.length+1}, ()=>0 );

    // Iterate through the input array and add the frequency of each element to the hash map
    for( const num of nums){
        hm[num] = (hm[num]|| 0) + 1
    }
    // Iterate through the hash map and add the elements to the frequency array based on their frequency
    for ( const key in hm){
        freq[hm[key]] = (freq[hm[key]] || []).concat(key);
    }
    // Create an empty array to store the top k frequent elements
    let ans = [];

    // Iterate through the frequency array from the highest frequency to the lowest
    for ( let j = freq.length - 1; j >= 0; j--){
        for (let i = 0; i < freq[j].length && k > 0; i++){
            ans.push(Number(freq[j][i]));
            k--;
        }
    }
    return ans;
};
```
