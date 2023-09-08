# tinyButton
**Theres also an existing wayy better and more efficent project [bouncer](https://github.com/eyelight/bouncer), but it was just a little bit too complex for me.**

A really simple button libary for TinyGo. I needed a simple button library for TinyGo projects, and i really liked the [JC_Button](https://github.com/JChristensen/JC_Button) library for aduino. So i decided to try translating it to Go (+ the #33 pullrequest that apparently fixes some noise issues).

The code quality is bad/non-existent and im aware that this code is probably pretty inefficient. This was my first Go "project". I wanted a simple button library for TinyGo but coulnd't find one. So i decided to translate a really simple but awesome button library for arduino to TinyGo.
So full credit goes to: [JChristensen](https://github.com/JChristensen) and his awesome and simple [JC_Button](https://github.com/JChristensen/JC_Button) library.

ps. idk how this licensing works so my bad if i violated any stuff. let me know and i will make this repo private :p

I don't know even how you would install this :pp. You could use ```go get github.com/sudokit/tinyButton```, buut because go thinks ```machine``` is a std library, it will try to look for it as an std library but it isnt 🤷‍♂️. see [issue #4079 on golangs repo](https://github.com/golang/go/issues/43079) for related info

example usage (just a simple click button):
```go
import (
    "machine"
    "github.com/sudokit/tinyButton"
)

func main() {
    button_pin := machine.D5
    button := tinyButton.NewButton(button_pin)
    button.Configure(25, true, true)

    for {
        button.Read()
        if button.IsPressed() {
            println("Button is pressed!)
        }
    }
}
```
