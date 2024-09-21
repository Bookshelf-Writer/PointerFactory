![Fork GitHub Release](https://img.shields.io/github/v/release/Bookshelf-Writer/PointerFactory)
![Tests](https://github.com/Bookshelf-Writer/PointerFactory/actions/workflows/go-test.yml/badge.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/Bookshelf-Writer/PointerFactory)](https://goreportcard.com/report/github.com/Bookshelf-Writer/PointerFactory)

![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/Bookshelf-Writer/PointerFactory?color=orange)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Bookshelf-Writer/PointerFactory?color=green)
![GitHub repo size](https://img.shields.io/github/repo-size/Bookshelf-Writer/PointerFactory)


# PointerFactory
Библиотека для генерации/валидации уникальных UID в рамках кластера

---

### Генератор

- Привязка происходит к дате, точка отсчета задается при инициализации
- Допустимое количество параллельных нод для сохранения уникальности в рамках группы - 65535
- В рамках группы, в минуту можно получить до 4294967295 уникальных указателей
- Валидация многоступенчатая, возврат ошибки при первом же вхождении
- Имеется контрольная сумма (состоит из двух символов во избежания коллизий)
- Максимально допустимый размер системы счисления 36. Минимальный - 2.

#### Пример использования

```Go
userGroup := 'u'
regGroup := 'r'
commentGroup := 'g'

groups := []rune{userGroup, regGroup, commentGroup}
startTime := time.Date(2024, 1, 1, 1, 1, 0, 0, time.UTC)

uid, err := New(groups, 0, 36, startTime)   //Инициализируем фабрику
if err != nil {
panic(err)
}

//Ждем пока фабрика запустится
for !uid.IsActive() {time.Sleep(10 * time.Millisecond)} 

newUserID, _ := uid.New(userGroup)
newRegID, _ := uid.New(regGroup)
newCommentID, _ := uid.New(commentGroup)

fmt.Println(newUserID, newRegID) //Выведет: u07qpgy0ya r07qpgy0v7
fmt.Println(uid.IsValid(newCommentID)) //Вернет nil так как указатель полностью валиден
uid.Close()
```

---

---

### Mirrors

- https://git.bookshelf-writer.fun/Bookshelf-Writer/PointerFactory