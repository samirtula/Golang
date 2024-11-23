package main

import (
	"fmt"
	s "strings"
    "unicode"
)

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)
	fmt.Println(len("1234"))

	s2 := "€£3"
	for x, y := range s2 {
		fmt.Printf("%#U starts at byte position %d\n", y, x)
	}
	fmt.Printf("s2 length: %d\n", len(s2))

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)
	fmt.Printf("s3 length: %d\n", len(s3))
	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()

	var f = fmt.Printf
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))
	f("%s\n", s.Title("tHis wiLL be A title!"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAlis"))
	f("EqualFold: %v\n", s.EqualFold("Mihalis", "MIHAli"))

	f("Prefix: %v\n", s.HasPrefix("Mihalis", "Mi"))
	f("Prefix: %v\n", s.HasPrefix("Mihalis", "mi"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "is"))
	f("Suffix: %v\n", s.HasSuffix("Mihalis", "IS"))
	f("Index: %v\n", s.Index("Mihalis", "ha"))
	f("Index: %v\n", s.Index("Mihalis", "Ha"))
	f("Count: %v\n", s.Count("Mihalis", "i"))

	f("Count: %v\n", s.Count("Mihalis", "I"))
	f("Repeat: %s\n", s.Repeat("ab", 5))
	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s", s.TrimLeft(" \tThis is a\t line. \n", "\n\t "))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t "))


    f("Compare: %v\n", s.Compare("Mihalis", "MIHALIS"))
    f("Compare: %v\n", s.Compare("Mihalis", "Mihalis"))
    f("Compare: %v\n", s.Compare("MIHALIS", "MIHalis"))
    f("Fields: %v\n", s.Fields("This is a string!"))
    f("Fields: %v\n", s.Fields("Thisis\na\tstring!"))
    f("%s\n", s.Split("abcd efg", ""))

    f("%s\n", s.Replace("abcd efg", "", "_", -1))
    f("%s\n", s.Replace("abcd efg", "", "_", 4))
    f("%s\n", s.Replace("abcd efg", "", "_", 2))

    lines := []string{"Line 1", "Line 2", "Line 3"}
    f("Join: %s\n", s.Join(lines, "+++"))
    f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))
    
    trimFunction := func(c rune) bool {
        return !unicode.IsLetter(c)
    }
    f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}
