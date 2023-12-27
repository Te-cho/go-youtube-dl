package main


import (
	"io"
	"flag"
	"fmt"
	"net/url"
	"os"
	"log"
	"os/exec"
	"os/user"
	"path/filepath"
)


const CMD = "youtube-dl"
const BEST_VIDEO_FORMAT = "18"
const BEST_AUDIO_FORMAT = "140"


// Youtube-dl Downloader
type YoutubeDl struct {
	Path string
	Format string
}

func extractVideoID(videoURL string) (string, error) {
	parsedURL, err := url.Parse(videoURL)
	if err != nil {
		return "", err
	}

	queryParams := parsedURL.Query()
	videoID := queryParams.Get("v")

	return videoID, nil
}


func (youtubedl YoutubeDl) DownloadSubs(url string) error {
	videoID, error := extractVideoID(url)
	if error != nil {
		fmt.Fprintln(os.Stderr, error)
		return error
	}

	cmdArgs := []string{
		CMD,
		"--write-auto-sub",
		"--sub-lang",
		"en",
		"--output", youtubedl.Path + videoID + ".srt",
		"--skip-download",
		url,
	}

	err := run(cmdArgs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}

	return nil
}

func (youtubedl YoutubeDl) DownloadAudio(url string) (error) {
	// Escape and quote the path to handle spaces or special characters
	escapedPath := "\"" + youtubedl.Path + "%(title)s-%(id)s.%(ext)s\""

	cmdArgs := []string{
		CMD,
		"--format", youtubedl.Format,
		"--extract-audio",
		"--audio-format", "m4a",
		"--audio-quality", "192",
		"--output", escapedPath,
		"--write-thumbnail",
		"--embed-thumbnail",
		url,
	}

	err := run(cmdArgs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}

func (youtubedl YoutubeDl) DownloadVideo(videoURL string) error {
	sep := string(filepath.Separator)
	// Escape and quote the path to handle spaces or special characters
	escapedPath := "\"" + youtubedl.Path + sep + "%(title)s-%(id)s.%(ext)s\""

	cmdArgs := []string{
		CMD,
		"--format", youtubedl.Format,
		"--output", escapedPath,
		"--embed-thumbnail",
		videoURL,
	}

	return run(cmdArgs)
}

func run(cmdArgs []string) error {
	fmt.Println("Starting ...")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Executing command: ", cmd.String())

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdout.Close()
		io.Copy(os.Stdout, stdout)
	}()

	go func() {
		defer stderr.Close()
		io.Copy(os.Stderr, stderr)
	}()

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func main() {

	var (
		audioFlag bool
		format string
		output string
		subsFlag bool
		url string
		videoFlag bool
		ytdlFlag bool
	)

    usr, err := user.Current()
    if err != nil {
        panic(err)
    }
	sep := string(filepath.Separator)
	path := filepath.Join(sep, "home", usr.Username, "Music")

	// Parsing command line flags
	flag.StringVar(&url, "url", "", "url of the youtube /audio/video to download")
	flag.BoolVar(&subsFlag, "s", false, "Download subs from the youtube for the specified video")
	flag.BoolVar(&audioFlag, "a", false, "Download audio")
	flag.BoolVar(&videoFlag, "v", false, "Download video")
	flag.StringVar(&format, "f", "", "Specify video format")
	flag.StringVar(&output, "o", path + sep, "Path where the video will be written to")
	flag.BoolVar(&ytdlFlag, "y", false, "Find where youtube-dl is installed and its version")

	flag.Parse()

	if _, err := os.Stat(output); os.IsNotExist(err) {
		fmt.Printf("Error: The specified path '%s' does not exist.\n", output)
		os.Exit(1)
	}

	// If user did not provide youtube video url, show usage.
	if url == "" && !ytdlFlag {
		flag.Usage()
		os.Exit(1)
	}

	ytdl := YoutubeDl{Path: output + sep}

	switch {
		case ytdlFlag:
			cmdp := exec.Command("which", "youtube-dl")
			output, err := cmdp.CombinedOutput()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
				return
			}
		
			fmt.Printf("youtube-dl is located at:\n%s\n", output)

			cmdv := exec.Command("youtube-dl", "--version")
			version, error := cmdv.CombinedOutput()
			if error != nil {
				fmt.Printf("Error executing command: %v\n", error)
				return
			}
			fmt.Printf("Version:\n%s\n", version)

		case audioFlag:
			if format != "" {
				ytdl.Format = format
			} else {
				ytdl.Format = BEST_AUDIO_FORMAT
			}
			err := ytdl.DownloadAudio(url)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

		case videoFlag:
			if format != "" {
				ytdl.Format = format
			} else {
				ytdl.Format = BEST_VIDEO_FORMAT
			}
			err := ytdl.DownloadVideo(url)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

		case subsFlag:
			fmt.Println("Downloading subs...")
			err := ytdl.DownloadSubs(url)
			if err != nil {
				fmt.Println("Error downloading video.")
			}

		default:
			// Invalid flag provided
			fmt.Fprintln(os.Stderr, fmt.Sprintf("Invalid option"))
			os.Exit(1)
	}
}
