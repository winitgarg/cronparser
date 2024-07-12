# Cron Parser

This project provides a Golang-based implementation for parsing cron expressions along with a command. The cron expression is provided in the format `"*/15 0 1,15 * 1-5 /usr/bin/find"`.

It uses Makefile to initiate the project. 

## Running the parser
Either you can build it from source code or use the binary directly to run the project. 

### Build the project from source code
You'll need the following installed on your macOS environment:
- Go (Golang) 1.16 or higher
- Git

#### Setup

1. **Install Go**:
   ```
   make install-go
   ```

2. **Build the project**:
    ```
    make build
    ```

3. **Run the project**:
   ```
   ./cronparser */15 0 1,15 * 1-5 /usr/bin/find"
   ```
### Running the binary directly

```
./cronparser */15 0 1,15 * 1-5 /usr/bin/find"
 ```

## Makefile usage
Below command should list out all the possible Makefile targets to build and run the project
```
make help
```

