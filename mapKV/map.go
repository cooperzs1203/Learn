package mapKV

import "log"

/*
如果不初始化 map，那么就会创建一个 nil map。nil map 不能用来存放键值对
var map_variable map[key_data_type]value_data_type

*/

var siMap map[string]int

func mapMain() {
	initMap()
	queryMap()
	deleteKV()
}

func initMap() {
	siMap = make(map[string]int)
	siMap["ny"] = 1
	siMap["pr"] = 2
	siMap["bj"] = 3

	log.Println(siMap)
}

func queryMap() {
	value , ok := siMap["pr"]
	if !ok {
		log.Println("Key Not Exist")
		return
	}

	log.Println(value)

	value2 , ok := siMap["ddd"]
	if !ok {
		log.Println("Key Not Exist")
		return
	}

	log.Println(value2)
}

func deleteKV() {
	delete(siMap , "bj")
	log.Println(siMap)
}