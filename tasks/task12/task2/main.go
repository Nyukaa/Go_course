//12.11.2. Configuration file handling
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Config struct {
    Server struct {
        Host string `json:"host"`
        Port int    `json:"port"`
    } `json:"server"`
    Database struct {
        Host     string `json:"host"`
        Port     int    `json:"port"`
        Username string `json:"username"`
        Password string `json:"password"`
        Database string `json:"database"`
    } `json:"database"`
    Logging struct {
        Level string `json:"level"`
        File  string `json:"file,omitempty"`
    } `json:"logging"`
}

// Load config from JSON file
func loadConfig(filename string) (*Config, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var config Config
    err = json.Unmarshal(data, &config)
    if err != nil {
        return nil, err
    }

    return &config, nil
}

// Save config to JSON file
func saveConfig(filename string, config *Config) error {
    data, err := json.MarshalIndent(config, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(filename, data, 0644)
}

func main() {
    // Create config
    config := Config{}
    config.Server.Host = "localhost"
    config.Server.Port = 8080
    config.Database.Host = "localhost"
    config.Database.Port = 5432
    config.Database.Database = "myapp"
    config.Logging.Level = "info"

    // Save to file
    err := saveConfig("config.json", &config)
    if err != nil {
        fmt.Println("Error saving config:", err)
        return
    }

    fmt.Println("Config saved successfully")

    // Load from file
    loaded, err := loadConfig("config.json")
    if err != nil {
        fmt.Println("Error loading config:", err)
        return
    }

    fmt.Printf("Server: %s:%d\n", loaded.Server.Host, loaded.Server.Port)
    fmt.Printf("Database: %s\n", loaded.Database.Database)
}
