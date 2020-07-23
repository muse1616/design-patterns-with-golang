package main


//懒汉模式 非线程安全 当正在创建时，有线程来访问此时ins = nil就会再创建，单例类就会有多个实例了。
type Singleton struct {
	id int
}
var ins *Singleton

func GetIns()*Singleton {
	// 单例未被创建
	if ins == nil{
		ins = &Singleton{}
	}
	return ins
}