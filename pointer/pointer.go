package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Age  int
	}
	// *Personポインタ型
	var p *Person

	// &でアドレスが分かる
	p = &Person{
		Name: "太郎",
		Age:  20,
	}
	fmt.Printf("変数pに格納されているアドレス :%p", p)

	a, b, c := 1, 1, 1

	// main関数内で a に 2 をかける
	a *= 2
	fmt.Printf("main() aの値: %d\n\n", a) // 2

	// multiply2関数内で b に 2 をかける. 引数には b の値を渡す.
	multiply2(b)
	fmt.Printf("main() bの値: %d, bのアドレス: %p\n\n", b, &b)

	// multiply2pointer関数内で c に 2 をかける. 引数には c のアドレスを渡す.
	multiply2pointer(&c)
	fmt.Printf("main() cの値: %d, cのアドレス: %p\n\n", c, &c)
}

func multiply2(b int) {
	b *= 2
	// 違うbを計算している
	fmt.Printf("multiply2() bの値: %d, bのアドレス: %p\n", b, &b)
}

func multiply2pointer(c *int) {
	*c *= 2 // c のアドレスに格納されている値: 1 に 2 をかける.
	// 呼び出し元と同じc
	fmt.Printf("multiply2pointer() cの値: %d, cのアドレス: %p\n", *c, c)

	sub()
}

// output >
// main() aの値: 2

// multiply2() bの値: 2, bのアドレス: 0xc000016120
// main() bの値: 1, bのアドレス: 0xc0000160f8

// multiply2pointer() cの値: 2, cのアドレス: 0xc000016100
// main() cの値: 2, cのアドレス: 0xc000016100


func sub() {
	type Person struct {
		Name string
		Age  int
	}
	// 値として、pに代入
	p := Person{
			Name: "太郎",
			Age:  20,
	}

	fmt.Printf("最初のp :%+v\n", p)

	p2 := p
	p2.Name = "二郎"
	p2.Age = 21
	// pではなく
	fmt.Printf("p2で二郎に書き換えを行なったはずのp :%+v\n", p)
	// 【そのまま渡すだけじゃ変化なし】

	// &pで*Person(Personのポインタ型)を生成する
	// p3はpのアドレスが格納されている状態になる
	p3 := &p
	p3.Name = "二郎"
	p3.Age = 21
	// 【ポインタを渡すことで元のデータも変わる】

	fmt.Printf("p3で二郎に書き換えを行なったp :%+v\n", p)
}