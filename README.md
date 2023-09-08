# tinyButton
A really simple button libary for TinyGo. I needed a simple button library for TinyGo projects, and i really liked the JC_Button library for aduino. So i decided to try translating it to Go (+ the #33 pullrequest that apparently fixes some noise issues).
So full credit goes to: [JChristensen](https://github.com/JChristensen) and his awesome and simple [JC_Button](https://github.com/JChristensen/JC_Button) library.

ps. idk how this licensing works so my bad if i violated any stuff. let me know and i will make this repo private :p

The code may be/is really bad. This was my first Go "project". (so i don't know even how you would install this :pp. You could use ```go get github.com/sudokit/tinyButton```, buut because go thinks ```machine``` is a std library, it will try to look for it as an std library but it isnt 🤷‍♂️. see [issue #4079 on golangs repo](https://github.com/golang/go/issues/43079) for related info)
(just found it but theres another awesome project for TinyGo buttons if you need a better and more efficent alternative: [bouncer](https://github.com/eyelight/bouncer). It just was a little too complicated for my needs :D)

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
