switch ($args[0]) {
    "release" {
        go build -ldflags="-s -w -H=windowsgui" -trimpath -o myapp.exe main.go
        echo "Generated myapp.exe in release mode"
    }
    default {
        go build -o myapp.exe main.go
        echo "Generated myapp.exe in development mode"
    }
}


