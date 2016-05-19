func naive() {
    for ii := 0; ii < 10000; ii++ {
        database.Magic()
    }
}

func fancy() {
    var wait sync.WaitGroup
    for ii := 0; ii < 10000; ii++ {
        wait.Add(1)
        go func() {
            database.Magic()
            wait.Done()
        }()
    }
    wait.Wait()
}