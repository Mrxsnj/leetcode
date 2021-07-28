func add(a int, b int) int {
    for b != 0 {
        // carry
        c := (a & b) << 1
        // non-carry sum
        a ^= b
        b = c
    }

    return a
}
