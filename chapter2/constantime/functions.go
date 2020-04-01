package constantime

// ThreeWords Big O(1)
func ThreeWords() string {
	threewords := [3]string{"foo", "bar", "baz"}
	return threewords[1]
}

// TenWords Big O(1)
func TenWords() string {
	tenwords := [10]string{"foo", "bar", "baz", "qux", "grault", "waldo", "plugh", "zyzzy", "thud", "spam"}
	return tenwords[6]
}
