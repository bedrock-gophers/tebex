Tebex: Tebex Library for Dragonfly

Tebex is a library designed for Dragonfly servers, facilitating integration with the Tebex platform for managing online stores. With Tebex, you can seamlessly connect your Dragonfly server with your Tebex store, enabling features such as in-game purchases, virtual currency, and more.
Getting Started
Creating a New Store Instance

To initialize a Tebex store instance in your Dragonfly server, follow these steps:

## Define a function to load the Tebex store connection.
```go
// loadStore initializes the Tebex store connection.
func loadStore(key string, log *logrus.Logger) *tebex.Client {
    store := tebex.NewClient(log, time.Second*5, key)
    name, domain, err := store.Information()
    if err != nil {
        log.Fatalf("tebex: %v", err)
        return nil
    }
    log.Infof("Connected to Tebex under %v (%v).", name, domain)
    return store
}
```
## Call the loadStore function with your Tebex API key and a logger instance.
```go
key := os.Getenv("TEBEX_KEY")
log := logrus.New()
store := loadStore(key, log)
```

# Running Pending Player Commands on Player Join

Call the ExecuteCommands method on the Tebex client instance when a player joins the server:
```go
func acceptFunc(store *tebex.Client) func(p *player.Player) {
    return func(p *player.Player) {
        store.ExecuteCommands(p)
    }
}
```