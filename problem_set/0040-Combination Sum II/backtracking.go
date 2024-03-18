func combinationSum2(candidates []int, target int) [][]int {
    n := len(candidates)
    sort.Ints(candidates) // sort the array
    res := make([][]int, 0)
    var find func(start int, subs []int, num int) //declare function find

    // define function find
    find = func(start int, subs []int, num int) {
        //base conditions
        // if num is zero, then we found a combination
        if num == 0 {
            tmp := make([]int, len(subs))
            copy(tmp, subs) // copy subs to tmp
            res = append(res, tmp) // append tmp to res
            return
        }
        // if start is equal to n or num < 0 then return
        if start == n || num < 0 {
            return
        }
        // loop through the candidates
        for i, j := start, start+1; i < n; i, j = j, j+1 {
            // check for duplicates
            for j < n && candidates[i] == candidates[j] {
                j++
            }
            var sum int
            c := subs
            // attempt each candidate by decreasing the target by the candidate
            for k := i; k < j; k++ {
                sum += candidates[k]
                if sum > num {
                    break
                }
                c = append(c, candidates[k])
                find(j, c, num-sum) // recursively find the next candidate
            }
        }
    }
    find(0, nil, target)
    return res
}