package main

import (
    "archive/zip"
    "io"
    "os"
    "path/filepath"
    "regexp"
)

func unZip(src, dest string) error {
    r, err := zip.OpenReader(src)
    if err != nil {
        return err
    }
    defer r.Close()

    ext := filepath.Ext(src)
    rep := regexp.MustCompile(ext + "$")
    dir := filepath.Base(rep.ReplaceAllString(src, ""))

    destDir := filepath.Join(dest, dir)
    if err := os.MkdirAll(destDir, os.ModeDir); err != nil {
        return err
    }

    for _, f := range r.File {
        if f.Mode().IsDir() {
            continue
        }
        if err := saveUnZipFile(destDir, *f); err != nil {
            return err
        }
    }

    return nil
}

func saveUnZipFile(destDir string, f zip.File) error {
    destPath := filepath.Join(destDir, f.Name)
    if err := os.MkdirAll(filepath.Dir(destPath), f.Mode()); err != nil {
        return err
    }
    rc, err := f.Open()
    if err != nil {
        return err
    }
    defer rc.Close()
    destFile, err := os.Create(destPath)
    if err != nil {
        return err
    }
    defer destFile.Close()
    if _, err := io.Copy(destFile, rc); err != nil {
        return err
    }

    return nil
}
