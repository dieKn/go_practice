package main

import (
    "fmt"
    "time"
)

func main() {
    const totalExecuteNum = 6   //合計実行数
    const maxConcurrencyNum = 3 //同時実行数
    sig := make(chan string, maxConcurrencyNum)
    res := make(chan string, totalExecuteNum)
    defer close(sig)
    defer close(res)
    fmt.Printf("start concurrency execute %s \n", time.Now())
    for i := 0; i < totalExecuteNum; i++ {
        go wait6Sec(sig, res, fmt.Sprintf("no%d", i))
    }
    for {
        //全部が終わるまで待つ
        if len(res) >= totalExecuteNum {
            break
        }
    }

    fmt.Printf("end concurrency execute %s \n", time.Now())
}

func wait6Sec(sig chan string, res chan string, name string) {
    sig <- fmt.Sprintf("sig %s", name)
    time.Sleep(6 * time.Second)
    fmt.Printf("%s:end wait 6sec \n", name)
    res <- fmt.Sprintf("sig %s", name)
    <-sig
}
