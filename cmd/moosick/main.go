package main

import (
	sm "github.com/briananakpintar/moosick/internal/scenemanager"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func registerScenes() {
    sceneManager := sm.GetInstance()

    mainMenu := sm.CreateMainMenu()
    sceneManager.RegisterScene(mainMenu)

    editor := sm.CreateEditorScene()
    sceneManager.RegisterScene(editor)
}

func main() {
	rl.InitWindow(800, 450, "Moosick")
	defer rl.CloseWindow()
    
	rl.SetTargetFPS(60)
    sceneManager := sm.GetInstance() 
    registerScenes()
    sceneManager.SwitchSceneByName("MainMenu") 

	for !rl.WindowShouldClose() && !sceneManager.ShouldQuit() {
        currScene := sceneManager.GetCurrScene()
        if (currScene == nil) {
            return
        }

        currScene.Update()
        currScene.HandleInput()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
        currScene.Render()
		rl.EndDrawing()
	}
}
