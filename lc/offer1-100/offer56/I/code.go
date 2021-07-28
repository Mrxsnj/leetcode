func singleNumbers(nums []int) []int {
    x := 0
    y := 0
    z := 0
    p := 1

    for _, num := range nums {
        // z = x ^ y
        z ^= num
    }

    //  first lower non-zero bit
    for z & p == 0 {
        p <<= 1
    }

    for _, num := range nums {
        if num & p == 0 {
            x ^= num
        } else {
            y ^= num
        }
    }

    return []int{x, y}
}
