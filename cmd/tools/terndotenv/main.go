package main

import (
  "os/exec"

  "github.com/joho/godotenv"
)

func main() {
  if err := godotenv.Load(); err != nil {
    panic(err)
  }

  cmd := exec.Command(
    "tern",
    "/home/john/go/bin/tern",
    "migrate",
    "--migrations",
    "./internal/store/pgstore/migrations",
    "--config",
    "./internal/store/pgstore/migrations/tern.conf",
    )
  if err := cmd.Run(); err != nil {
    panic(err)
  }
}
