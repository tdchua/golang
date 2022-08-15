package main

import (
    "context"
    "fmt"    
    "log"
    "os"
    "strconv"
    "github.com/shomali11/slacker"
)
func printCommandEvents(analyticsChannel <- chan *slacker.CommandEvent) {
    for event := range analyticsChannel {
        fmt.Println("Command Events")
        fmt.Println(event.Timestamp)
        fmt.Println(event.Command)
        fmt.Println(event.Parameters)
        fmt.Println(event.Event)
        fmt.Println()
    }
}


func main() {
    os.Setenv("SLACK_BOT_TOKEN", "xoxb-3962774274064-3938989626387-EdcqCnSI41hItHvXKHqrUBDN")
    os.Setenv("SLACK_APP_TOKEN", "xapp-1-A03UAP1M0CQ-3962784683344-b21aa40e4a4ce3face943b4d291f71599212cc40e524373dd2a333d78bc2267d")
    
    bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN")) 
    
    go printCommandEvents(bot.CommandEvents())
    
    bot.Command("my yob is <year>", &slacker.CommandDefinition {
        Description: "yob calculator",
        // Example:    "my yob is 2020",
        Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
            year := request.Param("year")
            yob, err := strconv.Atoi(year)
            if err != nil {
                println("error")
            }
            age := 2022 - yob
            r := fmt.Sprintf("age is %d", age)
            response.Reply(r)
        },
    })
    
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    err := bot.Listen(ctx)
    if err != nil {
        log.Fatal(err)
    }
}
