package main

import (
    "github.com/FloatTech/ZeroBot-Plugin"
    "github.com/FloatTech/ZeroBot-Plugin/feature"
    "math/rand"
    "strconv"
    "time"
)

var (
    MyPlugin = ctx.NewPlugin("MyPlugin") // 插件名称
)

func init() {
    MyPlugin.OnCommand("!bet").SetBlock(true).Handle(func(ctx *ctx.Context) {
        args := ctx.State["args"].(map[string]string)
        timesStr := args["times"]
        times, _ := strconv.Atoi(timesStr)
        if times < 1 {
            ctx.Send("最少只能放1发欧")
            return
        }
        if times > 6 {
            ctx.Send("枪里最多只能塞6发欧，再多就溢出来了....")
            return
        }

        rand.Seed(time.Now().UnixNano())
        for i := 0; i < times; i++ {
            if rand.Intn(2) == 1 {
                ctx.Send("被爆头了！")
                return
            }
        }
        ctx.Send("运气不错，没被爆头！")
    }).SetRegex(`!bet (\d+)`)
}

func main() {
    // 这里可以添加插件初始化的代码
}
