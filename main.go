package main

func main() {
    cha, chb := make(chan int), make(chan int)

    go func() {
        cha <- 1 // cha通道的数据没有被其他goroutine读取走，堵塞当前goroutine
        chb <- 0
    }()
	<- cha
    <- chb // chb 等待数据的写
}
