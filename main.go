package main

import (
	"fmt"
	"html"
	"runtime"
	"time"

	"github.com/Toyz/BlizzVersionViewer/btapi"
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	strip "github.com/grokify/html-strip-tags-go"
	imgui "github.com/inkyblackness/imgui-go"
)

var games []btapi.Game
var regionInfo []btapi.RegionInfo
var patchNote *btapi.PatchNote

var selectedChannel btapi.Channel
var loadingregion bool

func getGames() {
	localGames, err := btapi.AllGames()
	if err != nil {
		// we will probally panic
		panic(1)
	}

	games = localGames

	selectedChannel = games[0].Channels[0]
	go handleViewingGame(selectedChannel)
}

func handleViewingGame(channel btapi.Channel) {
	localGames, err := channel.Versions()
	if err != nil {
		panic(1)
	}

	localNotes, err := channel.PatchNotes(1, 1)

	if localNotes.PatchNotes != nil && len(localNotes.PatchNotes) == 1 {
		patchNote = &localNotes.PatchNotes[0]
	} else {
		patchNote = nil
	}

	regionInfo = localGames
	loadingregion = false
}

func main() {
	runtime.LockOSThread()

	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	loadingregion = true
	go getGames()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, 1)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	window, err := glfw.CreateWindow(1280, 720, "Blizz Version Viewer", nil, nil)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	window.MakeContextCurrent()
	glfw.SwapInterval(1)
	gl.Init()

	context := imgui.CreateContext(nil)
	defer context.Destroy()

	// imgui.CurrentIO().SetFontGlobalScale(2.0)
	imgui.CurrentIO().Fonts().AddFontFromFileTTF("./res/Roboto-Regular.ttf", 18.0)
	// imgui.CurrentIO().Fonts().SetTextureID()

	impl := imguiGlfw3Init(window)
	defer impl.Shutdown()

	for !window.ShouldClose() {
		glfw.PollEvents()
		impl.NewFrame()

		imgui.SetNextWindowSize(imgui.Vec2{float32(1280), float32(720)})
		imgui.SetNextWindowPos(imgui.Vec2{float32(0), float32(0)})
		imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 0)
		{
			imgui.BeginV("Select Game", nil, imgui.WindowFlagsMenuBar|imgui.WindowFlagsNoNav|imgui.WindowFlagsNoCollapse|imgui.WindowFlagsNoResize|imgui.WindowFlagsNoTitleBar|imgui.WindowFlagsNoMove)

			if imgui.BeginMenuBar() {
				if imgui.BeginMenu("File") {
					if imgui.MenuItem("Quit") {
						window.SetShouldClose(true)
					}

					imgui.EndMenu()
				}

				imgui.EndMenuBar()
			}

			imgui.BeginChildV("GameList", imgui.Vec2{500, 685}, true, 0)

			if games != nil {
				for _, game := range games {
					for _, channel := range game.Channels {
						imgui.BeginGroup()

						if imgui.Button(fmt.Sprintf("View Info##%s", channel.Code)) && !loadingregion {
							loadingregion = true
							selectedChannel = channel
							go handleViewingGame(channel)
						}

						imgui.SameLineV(0, 15)
						imgui.CurrentStyle().SetColor(imgui.StyleColorText, yellow)
						imgui.Text(channel.Name)
						imgui.CurrentStyle().SetColor(imgui.StyleColorText, white)
						imgui.EndGroup()
					}
				}
			}
			imgui.EndChild()

			if regionInfo != nil {
				imgui.SameLineV(0, 30)
				imgui.BeginGroup()
				imgui.PushItemWidth(685)
				imgui.Text(fmt.Sprintf("%60v", selectedChannel.Name))
				imgui.BeginChildV("GameViewing", imgui.Vec2{733, 325}, true, 0)

				imgui.Text(fmt.Sprintf("%-25v%-19v%s", "Region", "Version", "Build Config"))
				imgui.Separator()
				imgui.Text("")

				for idx, region := range regionInfo {
					imgui.Text(fmt.Sprintf("%-25v%-19v%s", region.Regionname, region.Versionsname, region.Buildconfig))
					imgui.SameLineV(0, 70)
					//imgui.Button("View Build Info")
					if idx < len(regionInfo)-1 {
						imgui.Separator()
					}
				}

				imgui.EndChild()

				imgui.BeginChildV("PatchNotesInfo", imgui.Vec2{733, 339}, true, 0)
				if patchNote == nil {
					imgui.Text("This game has no patch notes...")
				} else {
					imgui.PushTextWrapPosV(710)
					imgui.Text(html.UnescapeString(strip.StripTags(patchNote.Detail)))
				}

				imgui.EndChild()

				imgui.EndGroup()
			}

			imgui.End()
			imgui.PopStyleVar()
		}

		displayWidth, displayHeight := window.GetFramebufferSize()
		gl.Viewport(0, 0, int32(displayWidth), int32(displayHeight))
		gl.ClearColor(black.X, black.Y, black.Z, black.W)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		imgui.Render()
		impl.Render(imgui.RenderedDrawData())

		window.SwapBuffers()
		<-time.After(time.Millisecond * 25)
	}
}
