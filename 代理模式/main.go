package main

func main()  {
	image:=NewImageProxy("test.jpg")
	image.display()
	// 第二次无需加载
	image.display()
}