func singleNumber(nums []int) int {
    cache := make([]int, 32)
    res := 0

    for _, num := range nums {
        for i := 0; i < 32; i++ {
            if num & (1 << i) != 0 {
                cache[i]++
            }
        }
    }

    for i := 31; i >= 0; i-- {
        res <<= 1
        res |= cache[i] % 3
    }

    return res
}
