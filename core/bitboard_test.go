package core

import "fmt"

func ExampleBitboard_Clear() {
	b := NewBitboard(A1, H8)
	b.Clear(H8)
	fmt.Println(b.Debug())
	// Output:
	// ........
	// ........
	// ........
	// ........
	// ........
	// ........
	// ........
	// X.......
}

func ExampleBitboard_Set() {
	b := NewBitboard(A1, H8)
	b.Set(A2)
	fmt.Println(b.Debug())
	// Output:
	// .......X
	// ........
	// ........
	// ........
	// ........
	// ........
	// X.......
	// X.......
}

func ExampleBitboard_Get() {
	b := NewBitboard(A1, H8)
	fmt.Println(b.Get(A1))
	fmt.Println(b.Get(A2))
	// Output:
	// true
	// false
}

func ExampleBitboard_Count() {
	b := NewBitboard(A1, H8)
	fmt.Println(b.Count())
	// Output:
	// 2
}

func ExampleBitboard_FlipV() {
	b := NewBitboard(C2, D2, E2, F2)
	b.FlipV()
	fmt.Println(b.Debug())
	// Output:
	// ........
	// ..XXXX..
	// ........
	// ........
	// ........
	// ........
	// ........
	// ........
}

func ExampleBitboard_With() {
	p := NewBitboard(A1, H8)
	q := NewBitboard(A2, H7)
	p.With(q)
	fmt.Println(p.Debug())
	// Output:
	// .......X
	// .......X
	// ........
	// ........
	// ........
	// ........
	// X.......
	// X.......
}

func ExampleBitboard_First() {
	b := NewBitboard(E5, D4, C3)
	fmt.Println(b.First())
	// Output:
	// C3
}

func ExampleBitboard_Intersects() {
	a := NewBitboard(A1, H8)
	b := NewBitboard(A2, H7)
	c := NewBitboard(A1, H7)
	fmt.Println(a.Intersects(b))
	fmt.Println(a.Intersects(c))
	// Output:
	// false
	// true
}
