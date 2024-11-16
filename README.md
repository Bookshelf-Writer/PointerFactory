![Fork GitHub Release](https://img.shields.io/github/v/release/Bookshelf-Writer/PointerFactory)
![Tests](https://github.com/Bookshelf-Writer/PointerFactory/actions/workflows/go-test.yml/badge.svg)

[![Go Report Card](https://goreportcard.com/badge/github.com/Bookshelf-Writer/PointerFactory)](https://goreportcard.com/report/github.com/Bookshelf-Writer/PointerFactory)

![GitHub repo file or directory count](https://img.shields.io/github/directory-file-count/Bookshelf-Writer/PointerFactory?color=orange)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/Bookshelf-Writer/PointerFactory?color=green)
![GitHub repo size](https://img.shields.io/github/repo-size/Bookshelf-Writer/PointerFactory)


# PointerFactory
Library for generating/validating unique UIDs within a cluster

---

### Generator

- Binding occurs to the date, the reference point is set during initialization
- The allowed number of parallel nodes to maintain uniqueness within the group is 65535
- Within a group, you can receive up to 4294967295 unique pointers per minute
- Multi-stage validation, returning an error at the first occurrence
- There is a checksum (consists of two characters to avoid collisions)
- The maximum allowable size of the number system is 36. The minimum is 2.

#### Usage example

```Go
userGroup := 'u'
regGroup := 'r'
commentGroup := 'g'

groups := []rune{userGroup, regGroup, commentGroup}
startTime := time.Date(2024, 1, 1, 1, 1, 0, 0, time.UTC)

uid, err := New(groups, 0, 36, startTime)   //Initialize the factory
if err != nil {
panic(err)
}

newUserID, _ := uid.New(userGroup)
newRegID, _ := uid.New(regGroup)
newCommentID, _ := uid.New(commentGroup)

fmt.Println(newUserID, newRegID) //Outputs: u09vnb0l9 r09vnb0i6
fmt.Println(uid.IsValid(newCommentID)) //Returns nil since the pointer is valid
uid.Close()
```

---

---

### Mirrors

- https://git.bookshelf-writer.fun/Bookshelf-Writer/PointerFactory