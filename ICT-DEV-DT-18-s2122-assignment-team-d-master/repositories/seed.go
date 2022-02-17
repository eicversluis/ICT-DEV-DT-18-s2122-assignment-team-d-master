package repositories

func Create(Data interface{}) {
	connection().Create(Data)
}
