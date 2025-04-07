### Arrays

```go
    var arr1 []int // nil
	var arrEmpty = []int{} // zero size
	var arr2 = [3]int{1, 2, 3} // 3 size array
	var arr4 = [...]int{1, 2, 3} // 3 size array
```

### Slices

```go
    vlen := 3
	vcap := 8
	var sl1 []int = make([]int, vlen)
	var sl2 []int = make([]int, vlen, vcap)
	fmt.Println(slices.Equal(sl1, sl2))
	sl1 = append(sl1, 3)
	fmt.Println(sl1, cap(sl1), len(sl1))
	fmt.Println(sl2, cap(sl2), len(sl2))
```

### Strings

```go
    var char rune = 'A'
	var bytea byte = 'A'
	var emoji []byte = []byte("😃")
	fmt.Println(char, bytea)                 // 65 65
	fmt.Println(string(char), string(bytea)) // A A
	fmt.Println(string(emoji), emoji)        // 😃 [240 159 152 131]
```