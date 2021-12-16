package services

import (
    "fmt"
    "log"
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

func (l *Logger) LogAsync(content string) {
    go func () {
        l.mutex.Lock()
        defer l.mutex.Unlock()
    
        if err := os.MkdirAll(filepath.Dir(l.path), 0777); err != nil {
            log.Println("Unable to create dir log:", err)

            return
        }
    
        file, err := os.OpenFile(l.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
        if err != nil {
            log.Println("Unable to open log:", err)

            return
        }
    
        now := time.Now()
        timeFormatted := now.Format(time.RFC3339)
        if _, err = file.WriteString(fmt.Sprintf("[%s] %s\n", timeFormatted, content)); err != nil {
            log.Println("Unable to write log:", err)

            return
        }
    
        if err = file.Close(); err != nil {
            log.Println("Unable to close log:", err)

            return
        }
    }()
}
