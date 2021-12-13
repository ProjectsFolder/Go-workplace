package services

import (
    "fmt"
    "os"
    "path/filepath"
    "sync"
    "time"
    "workplace/internal/config"
)

type Logger struct {
    mutex sync.Mutex
    path  string
}

func CreateLogger(config *config.Configuration) *Logger {
    return &Logger{path: config.LogPath}
}

func (l *Logger) Log(content string) error {
    l.mutex.Lock()
    defer l.mutex.Unlock()

    if err := os.MkdirAll(filepath.Dir(l.path), 0777); err != nil {
        return err
    }

    file, err := os.OpenFile(l.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
    if err != nil {
        return err
    }
    
    now := time.Now()
    timeFormatted := now.Format(time.RFC3339)
    if _, err = file.WriteString(fmt.Sprintf("[%s] %s\n", timeFormatted, content)); err != nil {
        return err
    }

    if err = file.Close(); err != nil {
        return err
    }

    return nil
}
