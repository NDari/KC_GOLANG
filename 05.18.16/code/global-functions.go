// Bad
func main(){
    fmt.Println("Hey buddeh!")
    log.Println("I'm not your buddeh, friend!")
}
// Good
type ILogger interface {
    Info(message string)
}
type ConsoleLogger struct{}
func (cl *ConsoleLogger) Info(message string) {
    fmt.Println(message)
}
func main() {
    cl := new(ConsoleLogger)
    cl.Info("I'm not your friend, guy!")
}