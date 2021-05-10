package main
import (

        "net/http"
        "fmt"
        "os/exec"
)

func handler(w http.ResponseWriter, r *http.Request){

        fmt.Fprintln(w, "download done", r.URL.RawQuery)
        cmd := exec.Command("youtubedr", "download", "-q hd720", r.URL.RawQuery)
        err := cmd.Run()
        if err != nil{

                panic(err)
        }
        fmt.Println(w, cmd.Args)
}

func main(){

        http.HandleFunc("/", handler)
        http.ListenAndServe(":8080", nil)
}
