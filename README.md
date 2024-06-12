![Fork GitHub Release](https://img.shields.io/github/v/release/Bookshelf-Writer/PointerFactory)
![Tests](https://github.com/Bookshelf-Writer/PointerFactory/actions/workflows/go-test.yml/badge.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/Bookshelf-Writer/PointerFactory)](https://goreportcard.com/report/github.com/Bookshelf-Writer/PointerFactory)

![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/Bookshelf-Writer/PointerFactory?color=orange)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Bookshelf-Writer/PointerFactory?color=green)
![GitHub repo size](https://img.shields.io/github/repo-size/Bookshelf-Writer/PointerFactory)


# PointerFactory
Фабрика работы с UID в кластере

---

### Пример указателя

Строчный (внешний): `bzyrcn1hssighrb`

Числовой (внутренний): `129499999999999999999`

|                | **Тип** | **Сервер**   | **Указатель**     | **CRC** |
|----------------|---------|--------------|-------------------|---------|
| **внешний**    | b       | zy           | rcn1hssighr       | b       |
| **внутренний** |         | 1294         | 99999999999999999 |         |

#### Допустимые значения сервера: от `1` до `1294`

#### Допустимые значения указателя: от `1` до `99999999999999999`

#### Алгоритм CRC: [file](/checksum.go)

---

### Генератор

- Уникальность указателей сохраняется при создании до 9999 указателей в минуту на тип.
- При превышении возможны коллизии с будущими указателями.
- Поддерживается асинхронный режим создания. (уникальность при асинхронности сохраняется)

#### Пример использования

```Go
crt := PointerFactory.InitCreator(12) // Задаем сервер
id := crt.New(PointerFactory.TypeAuthor) //При создании указываем тип генерируемого указателя

fmt.Println(id.String(), id.StringINT(), id.Uint()) //Вывод трех возможных 'однострочных' вариантов
```
```Bash
>>> a0c18dlxjph 122683370001 122683370001
```

---

### Результат оптимизаций

- CRS - 2,49
- Линейная генерация - 4,84
- Асинхрон по потокам - 1,37
- Асинхрон с единым потоком - 1,8

---

---

### Mirrors

- https://git.bookshelf-writer.fun/Bookshelf-Writer/PointerFactory