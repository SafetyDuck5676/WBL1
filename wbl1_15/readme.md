15. К каким негативным последствиям может привести данный фрагмент кода, и как
это исправить? Приведите корректный пример реализации.


```go
var justString string
func someFunc() {
v := createHugeString(1 << 10)
justString = v[:100]
}
func main() {
someFunc()
}
```
Ответ: Данный фрагмент кода может привести к утечке памяти. При вызове функции createHugeString создается большая строка, занимающая 2^10 (или 1024) байт в памяти. Затем в переменную justString присваивается подстрока этой большой строки, содержащая только первые 100 символов. Однако, сама большая строка не удаляется из памяти, она продолжает занимать место.
