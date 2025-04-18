package scenemanager

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EditorScene struct {
    name string
	buttonNewProject    rl.Rectangle
	buttonImportProject rl.Rectangle
}

func (e EditorScene) GetName() string {
    return e.name;
}

func (e EditorScene) Update() {
    
}

func CreateEditorScene() *EditorScene {
    return &EditorScene{name: "Editor"}
}

func (e *EditorScene) Render() {
	rl.DrawText("Editor", 10, 10, 30, rl.Blue)

	e.buttonNewProject = rl.Rectangle{X: 100, Y: 100, Width: 200, Height: 50}
	e.buttonImportProject = rl.Rectangle{X: 100, Y: 170, Width: 200, Height: 50}

	rl.DrawRectangleRec(e.buttonNewProject, rl.LightGray)
	rl.DrawRectangleRec(e.buttonImportProject, rl.LightGray)
}

func (e *EditorScene) HandleInput() {
	mousePos := rl.GetMousePosition()

	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		if rl.CheckCollisionPointRec(mousePos, e.buttonNewProject) {
			fmt.Println("New Project button clicked!")
		}

		if rl.CheckCollisionPointRec(mousePos, e.buttonImportProject) {
			fmt.Println("Import Project button clicked!")
		}
	}
}
