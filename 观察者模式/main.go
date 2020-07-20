package main

func main() {
	//	主题
	wd := NewWeatherData()
	//	注册一个观察者
	_ = NewCurrentConditionDisplay(wd)
	//	更新数据
	wd.updateMeasurement(80, 12, 12)
	wd.updateMeasurement(1.1, 1, 1)
}
