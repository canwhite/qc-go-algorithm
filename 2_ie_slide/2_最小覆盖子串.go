package ieslide

import "math"

// 2. 最小覆盖子串:给定两个字符串s1和s2,找出s1中包含s2所有字符的最小子串。
// 使用滑动窗口,不断缩小窗口所有包含s2所有字符的子串。
func MinWindow(s1, s2 string) string {
	// 定义两个哈希表，need记录s2中每个字符的个数，
	///而需要统计的只是ASCII字符,一个byte(8位)就足够表示一个ASCII字符了
	//一个rune可以表示一个Unicode码点,它的大小是32位。
	//这里因为是ASCII，所以byte就够了
	need := make(map[byte]int)
	//window记录s1中窗口内每个字符的个数
	window := make(map[byte]int)
	for i := 0; i < len(s2); i++ {
		need[s2[i]]++
	}

	// 定义左右指针，分别指向窗口的左右边界
	left := 0
	right := 0
	// 定义valid变量，记录窗口内满足need条件的字符个数
	valid := 0
	// 定义start和len变量，记录最小覆盖子串的起始位置和长度
	start := 0
	length := math.MaxInt32

	for right < len(s1) {
		// c是将要移入窗口的字符
		c := s1[right]
		// 右移窗口
		right++
		// 如果c是s2中的字符，更新窗口内的数据
		if need[c] > 0 {
			window[c]++
			//window的主要作用就是在这里判断一个字符满足，然后更新valid
			if window[c] == need[c] {
				valid++
			}
		}

		// 判断左侧窗口是否要收缩
		// need的len和valid意思一样，如果有单个字母重复，是加到一个key上去的
		for valid == len(need) {
			// 更新最小覆盖子串的起始位置和长度
			// 更新发生在中间最合适，因为可以
			if right-left < length {
				start = left
				length = right - left
			}

			// 然后更新左侧
			// d是将要移出窗口的字符
			d := s1[left]
			// 左移窗口
			left++
			// 如果d是s2中的字符，更新窗口内的数据
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	// 如果没有找到符合条件的子串，返回空字符串，否则返回最小覆盖子串
	if length == math.MaxInt32 {
		return ""
	}
	return s1[start : start+length]
}
