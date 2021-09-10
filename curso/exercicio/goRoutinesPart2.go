package exercicio

import (
	"fmt"
	"runtime"
	"sync"
)

var contador int

var mu sync.Mutex

const quantidadeDeGoroutines = 50

var wgroup sync.WaitGroup

func Ex47() {
	wgroup.Add(50)
	criarGoroutines(quantidadeDeGoroutines)
	wgroup.Wait()
	fmt.Println("total de goroutines:", quantidadeDeGoroutines, "\ntotal do contador:", contador)
}

func criarGoroutines(i int) {
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wgroup.Done()
		}()
	}
}
