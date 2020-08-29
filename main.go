package main


import (
  //"fmt"
  "os"
  "github.com/docker/docker/client"
  "github.com/docker/docker/api/types"
  "context"
)

//refer https://github.com/moby/moby/blob/24e9ac15759cc23caf90383b5ea6de79af4363a0/client/image_build_test.go

func BuildDockerfile(repoName string, tagName string){
  cli, _ := client.NewEnvClient()

  dockerBuildContext, _ := os.Open("./work.tar.gz")
  defer dockerBuildContext.Close()

  buildOptions := types.ImageBuildOptions{
    //Dockerfile:   "Dockerfile", // optional, is the default
    Tags:   []string{repoName +":"+tagName},
    NoCache: true,
    SuppressOutput: true,
  }

  buildResponse, _ := cli.ImageBuild(context.Background(), dockerBuildContext, buildOptions)

  //fmt.Println(buildResponse.Body)

  defer buildResponse.Body.Close()
}

//func MakeTar(){}

/*
func WriteDockerfile(binname string){
  var file *os.File

  if Exists("../work/Dockerfile") == false{
    fmt.Println("Creating Dockerfile...")
    file, _ = os.Create("../work/Dockerfile")
  }else{
    file, _ = os.Open("../work/Dockerfile")
  }

  fmt.Println("Editting Dockerfile...")
  str :=  "FROM scratch\n" +
          "COPY " + binname + " /\n" +
          "CMD [\"/" + binname + "\"]"
  file.Write(([]byte)(str))

  defer file.Close()
}
*/


//is the file exeists?
func Exists(name string) bool {
    _, err := os.Stat(name)
    return !os.IsNotExist(err)
}




func main(){
	//binname := "hello"

	//Create Dockerfile
	//WriteDockerfile(binname)

	//docker build
	repoName := "hoge"
	tagName := "1"
	BuildDockerfile(repoName, tagName)
	//BuildDockerfile(binname)
}
