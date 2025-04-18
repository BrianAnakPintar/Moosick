package scenemanager

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type mainMenuScene struct {
    name string
	buttonNewProject    rl.Rectangle
	buttonImportProject rl.Rectangle
	buttonQuit          rl.Rectangle
}

func (mm mainMenuScene) GetName() string {
    return mm.name;
}

func (mm mainMenuScene) Update() {
    
}

func CreateMainMenu() *mainMenuScene {
    return &mainMenuScene{name: "MainMenu"}
}

func (mm *mainMenuScene) Render() {
	rl.DrawText("Main Menu", 10, 10, 30, rl.Blue)

	mm.buttonNewProject = rl.Rectangle{X: 100, Y: 100, Width: 200, Height: 50}
	mm.buttonImportProject = rl.Rectangle{X: 100, Y: 170, Width: 200, Height: 50}
	mm.buttonQuit = rl.Rectangle{X: 100, Y: 240, Width: 200, Height: 50}

	rl.DrawRectangleRec(mm.buttonNewProject, rl.LightGray)
	rl.DrawRectangleRec(mm.buttonImportProject, rl.LightGray)
	rl.DrawRectangleRec(mm.buttonQuit, rl.LightGray)

	rl.DrawText("New Project", int32(mm.buttonNewProject.X+30), int32(mm.buttonNewProject.Y+15), 20, rl.Black)
	rl.DrawText("Import Project", int32(mm.buttonImportProject.X+30), int32(mm.buttonImportProject.Y+15), 20, rl.Black)
	rl.DrawText("Quit", int32(mm.buttonQuit.X+80), int32(mm.buttonQuit.Y+15), 20, rl.Black)
}

func (mm *mainMenuScene) HandleInput() {
	mousePos := rl.GetMousePosition()

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if rl.CheckCollisionPointRec(mousePos, mm.buttonNewProject) {
            GetInstance().SwitchSceneByName("Editor")
		}

		if rl.CheckCollisionPointRec(mousePos, mm.buttonImportProject) {
			fmt.Println("Import Project button clicked!")
		}

		if rl.CheckCollisionPointRec(mousePos, mm.buttonQuit) {
			GetInstance().quitRequest = true
		}
	}
}
