package trace

//Tracer はコードないでの出来事を記録できるオブジェクトを表すインターフェース。
type Tracer interface {
	Trace(...interface{}) //(...interface{}): 任意の方の引数を何個でも(ゼロでも可)受け取ることを意味する。
}
