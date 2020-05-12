# go-waller
Changes your wallpaper to a specified one. Built with love in Golang.


##External package(s)
- go-waller
    - run "go get -v github.com/gabielfemi/go-waller"
    
##Usage
- git clone https://github.com/gabielfemi/go-waller
- open the repository in an editor or IDE of your choice e.g. VSCode, Atom, GoLand, e.t.c.
- open "set.go" file in the root directory,
- change ``folderPath`` to the folder where your pictures are stored and ``pictures`` to the various pics you want to include (separate them by a comma)
- run ``go run main.go`` then see your background changed. 


Extra: You can build an executable file with ``go build main.go``

Thanks for reading :)

##Imports
````````
import (
    "fmt"
    "syscall"
    "unsafe"

    "github.com/gabielfemi/go-filer"
)
````````
##Main code
````
func main() {
	folderPath := "C:\\Users\\Real Stuff\\Pictures"
	pictures := [...]string {"juli_and_gab.jpg","kali-linux-wallpaper-v1.png", "kali-linux-wallpaper-v3.png", "linux_debian.jpg", "kali-linux-wallpaper-v8.png", "kali-linux-wallpaper-v4.png", "Programmer-Developer-Wallpaper.jpg"}

	for i:= range pictures {
		fileName := pictures[i]
		completeFilePath := folderPath + "\\" + fileName
		time.Sleep(5000 * time.Millisecond)
		go_waller.SetWallPaper(completeFilePath)


	}

	main()

}

````
