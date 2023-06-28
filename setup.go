package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    fmt.Print("mttldev/sdk-env-setup\n")
    rootPath, _ := os.Getwd()
    rootDir := filepath.Dir(rootPath)

    // TODO: なんでパスを指定しなきゃいけない？
    addDir := "sdk-env-setup"
    fmt.Println(rootDir)

    fmt.Println("Setup starting...")

    fmt.Println("Downloading SDK...")
    // if err := DownloadFile("renpy-sdk.zip", sdk_url()); err != nil {
    // 	panic(err)
    // }
    fmt.Println("Extracting SDK...")
    if err := unZip(filepath.Join(rootDir, addDir, "renpy-sdk.zip"), filepath.Join(rootDir, addDir)); err != nil {
        panic(err)
    }
    fmt.Println("Moving SDK...")
    srcDir := filepath.Join(rootDir, addDir, "renpy-sdk", fmt.Sprintf("renpy-%s-sdk", sdk_version()))
    files, err := os.ReadDir(srcDir)
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        srcPath := filepath.Join(srcDir, file.Name())
        destPath := filepath.Join(rootDir, addDir, "renpy-sdk", file.Name())

        if err := os.Rename(srcPath, destPath); err != nil {
            panic(err)
        }
    }
    if err := os.Remove(srcDir); err != nil {
        panic(err)
    }

    sdk_path := filepath.Join(rootDir, addDir, "renpy-sdk")

    fmt.Println("Downloading Android Support...")
    if err := DownloadFile("rapt.zip", rapt_url()); err != nil {
        panic(err)
    }
    fmt.Println("Extracting Android Support...")
    if err := unZip(filepath.Join(rootDir, addDir, "rapt.zip"), sdk_path); err != nil {
        panic(err)
    }
    fmt.Println("Moving Android Support...")
    srcDir = filepath.Join(sdk_path, "rapt", "rapt")
    files, err = os.ReadDir(srcDir)
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        srcPath := filepath.Join(srcDir, file.Name())
        destPath := filepath.Join(sdk_path, "rapt", file.Name())

        if err := os.Rename(srcPath, destPath); err != nil {
            panic(err)
        }
    }
    if err := os.Remove(srcDir); err != nil {
        panic(err)
    }

    fmt.Println("Downloading iOS Support...")
    if err := DownloadFile("renios.zip", renios_url()); err != nil {
        panic(err)
    }
    fmt.Println("Extracting iOS Support...")
    if err := unZip(filepath.Join(rootDir, addDir, "renios.zip"), sdk_path); err != nil {
        panic(err)
    }
    fmt.Println("Moving iOS Support...")
    srcDir = filepath.Join(sdk_path, "renios", "renios")
    files, err = os.ReadDir(srcDir)
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        srcPath := filepath.Join(srcDir, file.Name())
        destPath := filepath.Join(sdk_path, "renios", file.Name())

        if err := os.Rename(srcPath, destPath); err != nil {
            panic(err)
        }
    }
    if err := os.Remove(srcDir); err != nil {
        panic(err)
    }

    fmt.Println("Downloading Web Platform Support...")
    if err := DownloadFile("web.zip", web_url()); err != nil {
        panic(err)
    }
    fmt.Println("Extracting Web Platform Support...")
    if err := unZip(filepath.Join(rootDir, addDir, "web.zip"), sdk_path); err != nil {
        panic(err)
    }
    fmt.Println("Moving Web Platform Support...")
    srcDir = filepath.Join(sdk_path, "web", "web")
    files, err = os.ReadDir(srcDir)
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        srcPath := filepath.Join(srcDir, file.Name())
        destPath := filepath.Join(sdk_path, "web", file.Name())

        if err := os.Rename(srcPath, destPath); err != nil {
            panic(err)
        }
    }
    if err := os.Remove(srcDir); err != nil {
        panic(err)
    }

    fmt.Println("Finish File Download.")
}
