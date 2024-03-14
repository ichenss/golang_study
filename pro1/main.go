package main

type Hero struct {
	name  string
	Age   int
	Level int
}

func (cur *Hero) GetName() string {
	return cur.name
}

func (cur *Hero) SetName(name string) {
	cur.name = name
}

func main() {
	hero1 := Hero{
		name:  "chen",
		Age:   25,
		Level: 100,
	}
	hero1.SetName("yuhang")
	println(hero1.GetName())
}
