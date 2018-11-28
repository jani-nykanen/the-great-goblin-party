![](https://img.itch.zone/aW1nLzE2NjQxNzYucG5n/original/eLZnxJ.png)

## [Play on Itch.io](https://jani-nykanen.itch.io/the-great-goblin-party)

------

## Building

**Linux:**

1) Make sure you have SDL2 and SDL_Mixer installed. 
2) Install Go bindings of those libraries available here: https://github.com/veandco/go-sdl2
3) Run *build.sh*, it should build everything and output a binary called "game"

**Windows:**

1) Do steps 1) & 2) Like above
2) Go to the src folder and type
`gobuild -o ../game.exe`
That might work, but I haven't tested it yet.

**Cross-building Windows binary on Linux:**
1) Install MinGW-w64 and Windows binaries of SDL2 and SDL_Mixer
2) Run build.win64.sh. It will build a 64-bit binary called "game.exe"

------

## License

I'm too lazy to find a proper license, so here's some conditions using the content in this repo:

**The code:**
This code was released for learning purposes. I hope one can learn Go programming by reading this code. Thus one is free to use any code written here in his or her own projects, although if one copies some code directly without any changes, it is recommended to give credit.

**Assets (graphics & audio):**
I recommend not to use them on your own projects, but if you really have to, make sure you project is not commercial and you give me credit for my work.

------

Copyright 2018 Jani Nykänen
