package main

func main() {
	m := map[int]int{1: 5, 2: 3, 4: 6, 6: 8}

	keys := MapKeys(m)

	Print(keys)

	Print("Hello ", "world", "!")

	lst := List[int]{}
	lst.Push(1)
	lst.Push(2)
	lst.Push(3)
	Print(lst.GetAll())
	Print(Sum(2, 5))
}
