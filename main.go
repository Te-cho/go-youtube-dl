package main


import (
	"io"
	"flag"
	"fmt"
	"os"
	"log"
	"os/exec"
	"os/user"
	"path/filepath"
)


const BASE_URL = "https://www.youtube.com/watch?v="
const CMD = "youtube-dl"
const BEST_VIDEO_FORMAT = "18"
const BEST_AUDIO_FORMAT = "140"


// Youtube-dl Downloader
type YoutubeDl struct {
	Path string
	Format string
}

func run(commandParams string) error {
	command :=  CMD + commandParams
	cmd := exec.Command("bash", "-c", command)

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

func (youtubedl YoutubeDl) DownloadSubs(id string) error {
	filename := "\"" + youtubedl.Path + "/" + id + ".srt\""
	commandParams := " --write-auto-sub --skip-download --sub-lang en -o " + filename + " -- " + id
	err := run(commandParams)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}

func (youtubedl YoutubeDl) DownloadMedia(id string) (error) {
	fmt.Println("Downloading ...")
	// Escape and quote the path to handle spaces or special characters
	escapedPath := "\"" + youtubedl.Path + "%(title)s-%(id)s.%(ext)s\""
	commandParams := " --format " + youtubedl.Format + " -o " + escapedPath + " " + BASE_URL + id
	err := run(commandParams)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}


func main() {
    usr, err := user.Current()
    if err != nil {
        panic(err)
    }

	path := filepath.Join("home", usr.Username, "Music")

	// Parsing command line flags
	yid := flag.String("yid", "", "ID of the youtube video to download subs")
	// url := flag.String("url", "", "url of the youtube /audio/video to download")
	s := flag.Bool("s", false, "Download subs from the youtube for the specified video id")
	a := flag.Bool("a", false, "download audio.")
	v := flag.Bool("v", false, "download video.")
	f := flag.String("f", "", "Specific media format to download.")
	p := flag.String("p", path, "Path where you want to download.")
	y := flag.Bool("y", false, "Find where youtube-dl is installed.")

	flag.Parse()

	if _, err := os.Stat(*p); os.IsNotExist(err) {
		fmt.Printf("Error: The specified path '%s' does not exist.\n", *p)
		os.Exit(1)
	}

	// If user did not provide id, show usage
	if *yid == "" && !*y {
		flag.Usage()
		os.Exit(1)
	}

	ytdl := YoutubeDl{Path: *p}

	switch {
		case *y:
			cmd := exec.Command("which", "youtube-dl")
			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
				return
			}
		
			fmt.Printf("Output:\n%s\n", output)
			path, err := exec.LookPath(CMD)
			if err != nil {
				fmt.Printf("Error executing command: %v\n", err)
			}
			fmt.Printf("path:\n%s\n", path)

		case *a:
			if *f != "" {
				ytdl.Format = *f
			} else {
				ytdl.Format = BEST_AUDIO_FORMAT
			}
			err := ytdl.DownloadMedia(*yid)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

		case *v:
			if *f != "" {
				ytdl.Format = *f
			} else {
				ytdl.Format = BEST_VIDEO_FORMAT
			}
			err := ytdl.DownloadMedia(*yid)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

		case *s:
			fmt.Println("Downloading subs...")
			err := ytdl.DownloadSubs(*yid)
			if err != nil {
				fmt.Println("Error downloading video.")
			}

		default:
			// Invalid flag provided
			fmt.Fprintln(os.Stderr, fmt.Sprintf("Invalid option"))
			os.Exit(1)
	}
}
