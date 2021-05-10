package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/NautiluX/webcamindicator/icon"
	"github.com/getlantern/systray"

	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " /dev/video<x>")
		return
	}
	systray.Run(onReady, onExit)
}

func onReady() {
	go checkWebcam()
	systray.SetIcon(icon.CamWhite)
	mQuit := systray.AddMenuItem("Quit", "Exit Webcam Indicator")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func checkWebcam() {
	mIcons := systray.AddMenuItem("", "")
	camDevice := os.Args[1]

	for {
		//fmt.Printf("checking webcam use...\n")
		cmd := exec.Command("bash", "-c", "lsof "+camDevice+"  2>/dev/null |  tail -n -1 | grep -v obs |awk '{print $1}'")
		b, err := cmd.Output()
		if err != nil {
			fmt.Printf("error getting programs: %v\n", err)
			time.Sleep(1 * time.Second)
			continue
		}
		res := strings.Split(strings.Trim(string(b), "\n "), "\n")
		if len(res) > 0 && res[0] != "" {
			fmt.Printf("following %d programs use the webcam: %v\n", len(res), res)
			systray.SetIcon(icon.CamRed)
			mIcons.SetTitle("Webcam " + camDevice + " in use by: " + strings.Join(res, ", "))
			time.Sleep(1 * time.Second)
			continue
		}
		mIcons.SetTitle("Webcam " + camDevice + " not in use.")
		systray.SetIcon(icon.CamWhite)
		time.Sleep(1 * time.Second)
	}
}

func onExit() {
	// clean up here
}
