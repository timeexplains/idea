package funcmode

type Fibo struct{
	a int64
	b int64
	c int64
}

func add(obj Fibo) (res Fibo){
	res.a = obj.a
	res.c = obj.b
	res.c = obj.a + obj.b
	return
}