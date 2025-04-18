package scenemanager

type Scene interface {
    GetName() string
    Update()
    Render()
    HandleInput()
}
