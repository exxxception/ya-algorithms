**Задача A**

Дан список. Определите, является ли он монотонно возрастающим
(то есть верно ли, что каждый элемент этого списка больше предыдущего).
Выведите YES, если массив монотонно возрастает и NO в противном случае.

Решение:
```go
func growingSequence(arr []int) {
    for i := 0; i < len(arr)-1; i++ {
        if arr[i] >= arr[i+1] {
            fmt.Println("NO")
            return
        }
    }
    fmt.Println("YES")
}
```

**Задача C**

Напишите программу, которая находит в массиве 
элемент, самый близкий по величине к  данному числу.

Решение:
```go
func closestNumber(arr []int, x int) {
    if len(arr) == 0 {
        return
    }

    closestNum := arr[0]
    distance := math.Abs(float64(x - closestNum))

    for i := 1; i < len(arr); i++ {
        newDist := math.Abs(float64(x - arr[i]))
        if newDist < distance {
            distance = newDist
            closestNum = arr[i]
        }
    }
	
    fmt.Println(closestNum)
}
```

**Задача D**

Дан список чисел. Определите, сколько в этом списке элементов, которые больше двух
своих соседей и выведите количество таких элементов.

Решение:
```go
func neighbors(arr []int) {
    count := 0
    for i := 1; i < len(arr)-1; i++ {
        if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
            count++
        }
    }

    fmt.Println(count)
}
```
