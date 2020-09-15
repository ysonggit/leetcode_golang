func isLongPressedName(name string, typed string) bool {
    i, j:=0, 0
    for i < len(name) && j < len(typed) {
        if name[i] != typed[j] {
            return false
        } else {
            cur, cnt:= name[i], i
            for j < len(typed) && typed[j] == cur {
                j++
                cnt++
            }
            for (i < len(name) && i < cnt) && name[i] == cur{
                i++
            }
        }
    }
    return i==len(name) && j==len(typed)
}
