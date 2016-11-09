package main

import (
    "fmt"
)

func main() {
    fmt.Println(`
    Enter following commands to control the player:
    lib list -- View the existing music lib
    lib add <name><artist><source><type> -- Add a music to the music lib
    lib remove <name> -- Remove the specified music from the lib
    play <name> -- Play the specified music`);



}
