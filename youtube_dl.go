package youtube_dl

import "os/exec"

// Youtube-dl SRT Downloader
type YoutubeDl struct {
	Path string
}

func (youtubedl YoutubeDl) DownloadVideo(id string) error{
	filename := "\"" + youtubedl.Path + "/" + id + ".srt\""
	commandParams := " --write-auto-sub --skip-download --sub-lang en -o " + filename + " -- " + id
	commandName := "youtube-dl"
	command := commandName + " " + commandParams
	cmd := exec.Command("bash", "-c", command)
	err := cmd.Run() // waits until the commands runs and finishes
	return err
}

