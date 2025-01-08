package main

func main(){
	todos := Todos{}
	todos.add("Learn Go")
	todos.add("Write a blog post")
	todos.add("Read a book")
	todos.toggle(0)
	todos.print()
}