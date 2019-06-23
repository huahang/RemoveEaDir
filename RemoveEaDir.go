package main

import (
  "container/list"
  "fmt"
  "os"
  "path/filepath"
)

func checkError(err error) error {
  if err != nil {
    fmt.Printf("[Error] Hit an error! " + err.Error() + "\n")
  }
  return err
}

func main() {
  args := os.Args
  if len(args) != 2 {
    fmt.Printf("Usage: RemoveEaDir [path to scan]\n")
    return
  }
  l := list.New()
  _ = filepath.Walk(args[1], func(path string, info os.FileInfo, err error) error {
    if err != nil {
      return checkError(err)
    }
    if info.Mode().IsDir() && info.Name() == "@eaDir" {
      l.PushBack(path)
    }
    return nil
  })
  for p := l.Front(); p != nil; p = p.Next() {
    fileInfo, err := os.Stat(p.Value.(string))
    if err != nil {
     _ = checkError(err)
     return
    }
    if fileInfo.IsDir() {
      fmt.Printf("[Removing] %s\n", p.Value.(string))
      _ = os.RemoveAll(string(p.Value.(string)))
    }
  }
}
