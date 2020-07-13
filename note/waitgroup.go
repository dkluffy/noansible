package core

import (
    "sync"
)

type WaitGroupWrapper struct {
    sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(f func()) {
    w.Add(1)
    go func() {
        f()
        w.Done()
    }()
}

// func main(){
//     var w WaitGroupWrapper
//     w.Wrap(foo)  //四行变成1行
//     ... // do otherthing
//     w.wait() 
// }