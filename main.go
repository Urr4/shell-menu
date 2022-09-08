package main

import (
    "github.com/gen2brain/raylib-go/raylib"
    "fmt"
    )

const (
	screenWidth  = 1000
	screenHeight = 480
)

var (
    running = true

    backgroundColor = rl.NewColor(147, 211, 196, 255)
    camera rl.Camera2D

    cursorSprite rl.Texture2D

    cursor Cursor
    cursorBump int
    cursorBumper func()
    moveCursor func() int

    fontBm rl.Font

    menu Menu

)

func main(){
    for running {
        input()
        render()
        update()
    }
    rl.CloseWindow()
}

type Cursor struct {
    Speed int
    Src rl.Rectangle
    Dest rl.Rectangle
    ShouldMove bool
    Direction int
    Label Text
}

type Text struct {
    Dest rl.Vector2
    Value string
}

type Menu struct {
    Entries []Text
}

/////
func init(){
    isCursor := true
    rl.InitWindow(screenWidth, screenHeight, "Test")
    rl.SetExitKey(0)
    rl.SetTargetFPS(60)

    fontBm = rl.LoadFont("assets/Fonts/pixantiqua.fnt");


    //TODO init cursor relative to first entry in menu
    cursorTarget := rl.NewRectangle(screenWidth/10, screenHeight/7, 50, 50)
    cursorBump = 3
    cursorBumper = bump()
    moveCursor = cursorMover()

    arr := [3]Text{
        Text{rl.NewVector2(screenWidth/10, screenHeight/7), "Hello"},
        Text{rl.NewVector2(screenWidth/10, 2*screenHeight/7), "World"},
        Text{rl.NewVector2(screenWidth/10, 3*screenHeight/7), "Blub"},
    }
    menu = Menu{arr[:]}

    if isCursor {
        cursorSprite = rl.LoadTexture("assets/Ui/72.png")
        cursor = Cursor{2, rl.NewRectangle(0, 0, 77, 76), cursorTarget, false, 0, Text{rl.NewVector2(70, 31), "PEX"}}
    }else{
        cursorSprite = rl.LoadTexture("assets/Characters/beards.png")
        cursor = Cursor{2, rl.NewRectangle(275, 0, 275, 275), cursorTarget, false, 0, Text{rl.NewVector2(55, 70), "Stefan"}}
    }
    camera = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), 0.0, 1.0)

}

func input(){
    if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
    		cursor.ShouldMove = true
    		cursor.Direction = -1
    }
    if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
        cursor.ShouldMove = true
        cursor.Direction = 1
    }

    if rl.IsKeyReleased(rl.KeyW) || rl.IsKeyReleased(rl.KeyUp) {
        cursor.ShouldMove = false
    }
    if rl.IsKeyReleased(rl.KeyS) || rl.IsKeyReleased(rl.KeyDown) {
        cursor.ShouldMove = false
    }
}

func render(){
    rl.BeginDrawing()
	rl.ClearBackground(backgroundColor)
	rl.BeginMode2D(camera)

    rl.DrawTexturePro(cursorSprite, cursor.Src, cursor.Dest, rl.NewVector2(cursor.Dest.Width, cursor.Dest.Height), 0, rl.White)
    rl.DrawTextEx(fontBm, cursor.Label.Value, cursor.Label.Dest, float32(fontBm.BaseSize), 2, rl.Black);

    for _, v := range menu.Entries {
        rl.DrawTextEx(fontBm, v.Value, v.Dest, float32(fontBm.BaseSize), 2, rl.Black);
    }

	rl.EndMode2D()
	rl.EndDrawing()
}

func update(){
    running = !rl.WindowShouldClose()

    cursorBumper()

    if cursor.ShouldMove {
        frameCount := moveCursor()
        fmt.Println(frameCount)
    }

    //TODO check if cursor reached end of menu, if so, wrap to start from beginning

}

func cursorMover() func() int {
    fc := -1

    return func() int {
        fc++
        if fc%5==0{
            fc = 0
            switch cursor.Direction {
                case -1: {
                    cursor.Dest.Y -= screenHeight/7
                    cursor.Label.Dest.Y -= screenHeight/7
                    cursor.ShouldMove = !cursor.ShouldMove
                }
                case 1: {
                    cursor.Dest.Y += screenHeight/7
                    cursor.Label.Dest.Y += screenHeight/7
                    cursor.ShouldMove = !cursor.ShouldMove
                }
                default: panic("Unkown movement")
            }
        }
        return fc
    }
}

//TODO closure seems to work well for functions that should only trigger on a given keycount. use this for those cases
func bump() func() {
    bumpCount := -1

    return func(){
        bumpCount++
        if bumpCount%20==0 {
            bumpCount = 0
            cursor.Dest.Y += float32(cursorBump)
            cursor.Label.Dest.Y += float32(cursorBump)
            cursorBump *= -1
        }
    }
}