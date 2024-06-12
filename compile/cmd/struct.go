package cmd

type flagInfoObj struct {
	name  string
	usage string
}

type flagPointBoolObj struct {
	info  flagInfoObj
	value *bool
}

type flagPointStringObj struct {
	info  flagInfoObj
	value *string
}

type flagPointUintObj struct {
	info  flagInfoObj
	value *uint
}
type flagPointUint64Obj struct {
	info  flagInfoObj
	value *uint64
}

//////////////////////////////////////////////////////////////////////////////

type flagObj struct {
	ver        flagPointBoolObj // Вывод информации о сборке
	jsonFormat flagPointBoolObj // Ответы в json

	printTypes flagPointBoolObj   // Вывести все доступные типы указателей
	validType  flagPointStringObj // Проверить валидный ли это тип

	createPointer flagPointBoolObj   // Режим создания указателя
	setType       flagPointStringObj // Задать тип
	setServer     flagPointUintObj   // Задать сервер
	setPointer    flagPointUint64Obj // Задать указатель

	setDateTime flagPointStringObj // Создать указатель по дате
	setCount    flagPointUintObj   // Количество указателей в эту дату (по умолчанию 1)

	validString flagPointStringObj // Ввод указателя (проверка на корректность)
	printInfo   flagPointBoolObj   // Подробная информация об указателе
}
