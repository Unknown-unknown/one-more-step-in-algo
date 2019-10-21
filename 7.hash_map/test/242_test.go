package test

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func Test_242(t *testing.T) {
	s1, s2 := "anagram", "nagaram"
	res1 := isAnagramBruteForce1(s1, s2)
	res2 := isAnagramBruteForce2(s1, s2)
	res3 := isAnagramHashMap(s1, s2)
	fmt.Printf("%v, %v, %v\n", res1, res2, res3)

	compareMap()
}

func isAnagramBruteForce1(s string, t string) bool {
	return sortedString(s) == sortedString(t)
}

func sortedString(s string) string {
	bs := strings.Split(s, "")
	sort.Strings(bs)
	return strings.Join(bs, "")
}

func isAnagramBruteForce2(s string, t string) bool {
	return sortedString2(s) == sortedString2(t)
}

func sortedString2(s string) string {
	bs := []byte(s)
	sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
	return string(bs)
}

func isAnagramHashMap(s string, t string) bool {
	cache1 := make(map[rune]int)
	cache2 := make(map[rune]int)
	for _, r := range s {
		cache1[r]++
	}
	for _, r := range t {
		cache2[r]++
	}
	return reflect.DeepEqual(cache1, cache2)
}

func compareMap() {
	map_1 := map[int]string{

		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
	}
	map_2 := map[int]string{

		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
		206: "Sumit",
	}

	map_3 := map[int]string{

		200: "Anita",
		201: "Neha",
		203: "Suman",
		204: "Robin",
		205: "Rohit",
	}
	map_4 := map[string]int{

		"Anita": 200,
		"Neha":  201,
		"Suman": 203,
		"Robin": 204,
		"Rohit": 205,
	}

	// Comparing maps
	// Using DeepEqual() function
	res1 := reflect.DeepEqual(map_1, map_2)
	res2 := reflect.DeepEqual(map_1, map_3)
	res3 := reflect.DeepEqual(map_1, map_4)
	res4 := reflect.DeepEqual(map_2, map_3)
	res5 := reflect.DeepEqual(map_3, map_4)
	res6 := reflect.DeepEqual(map_4, map_4)
	res7 := reflect.DeepEqual(map_2, map_4)

	// Displaying result
	fmt.Println("Is Map 1 is equal to Map 2: ", res1)
	fmt.Println("Is Map 1 is equal to Map 3: ", res2)
	fmt.Println("Is Map 1 is equal to Map 4: ", res3)
	fmt.Println("Is Map 2 is equal to Map 3: ", res4)
	fmt.Println("Is Map 3 is equal to Map 4: ", res5)
	fmt.Println("Is Map 4 is equal to Map 4: ", res6)
	fmt.Println("Is Map 2 is equal to Map 4: ", res7)
}
