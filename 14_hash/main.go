package main

import "fmt"

func main() {
	fmt.Println(hashint64(10003485834798))
	fmt.Println(hashstr("asv"))
	fmt.Println(hashstr("vsa"))
}

func hashint64(val int64) uint64 {
	// code here
	// Напишите свою собственную хеш-функцию,
	// которая будет возвращать хеш (типа uint64) от числа (типа int64),
	// используя остаток от деления на 1000.

	return uint64(val % 1000)
}

func hashstr(val string) uint64 {
	// code here
	// Напишите свою собственную хеш-функцию,
	// которая будет возвращать хеш (типа uint64) от строки (типа string),
	// используя остаток от деления на 1000.
	sum := 0
	for ind, el := range val {
		sum = +ind * int(el)
	}

	return uint64(sum % 1000)
}

// Реализуйте тип хеш-мап, который использует строки в качестве ключей и значений.
// Он должен реализовывать минимум три метода: Set, Get, Delete.

type hashmap struct {
	arr []string
}

func NewHashmap() *hashmap {
	return &hashmap{arr: make([]string, 1000)}
}

func (h *hashmap) Set(key, val string) {
	hashsum := hashstr(key)
	h.arr[hashsum] = val
}

func (h *hashmap) Get(key string) (value string, ok bool) {
	hashsum := hashstr(key)
	if h.arr[hashsum] != "" {
		return h.arr[hashsum], true
	}

	return "", false
}

func (h *hashmap) Delete(key string) {
	hashsum := hashstr(key)
	h.arr[hashsum] = ""
}
