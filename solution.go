package solution
import "github.com/kyokomi/emoji"
func GetMessage() string {
    hello := "Hello"
    world := "🗺️"
    rendered := emoji.Sprint(hello, world)
    rendered = rendered + "!"
    return rendered
}