func main() {
    g, ctx := errgroup.WithContext(context.Background())
    server := http.NewServer()

    // register server lifecycle
    g.Go(func() error {
        go func() {
            <- ctx.Done()   // handle exit
            fmt.Println("http server goint DOWN")
            server.Shutdown()
        }
        fmt.Println("http server start DONE")
        return server.Start()
    })

    // handle signal
    g.Go(func() error {
        sig := []os.Signal{syscall.SIGTERM, syscall.SIGINT}
        ch := make(chan os.Signal, len(sig))
        signal.Notify(ch, sig...)

        select {
        case <- ctx.Done():
            return errors.New("ctx Done()")
        case <- ch:
            server.Shutdown()
            return errors.New("get signal channel")
        }
    })

    // wait for all server\event exit
    if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
        fmt.Println("get error in Wait(),", err)
    }
}
