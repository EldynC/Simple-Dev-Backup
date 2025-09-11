# Go File Copy Tool

A simple and efficient command-line file copying utility written in Go that intelligently excludes common development artifacts and system files.

## Features

- **Smart File Filtering**: Automatically excludes common development directories and files
- **Recursive Directory Copying**: Copies entire directory structures with subdirectories
- **Path Resolution**: Handles relative paths (`.`, `./folder`) and absolute paths
- **Verbose Output**: Shows which files are being processed and skipped
- **Cross-Platform**: Works on macOS, Linux, and Windows

## Excluded Files and Directories

The tool automatically skips these common development artifacts:

### Development Dependencies
- `node_modules` - Node.js dependencies
- `vendor` - Go dependencies (vendor mode)
- `.venv` - Python virtual environments

### Build Outputs
- `dist` - Distribution/build output
- `build` - Build output directories
- `out` - Build output directories
- `.next` - Next.js build
- `.nuxt` - Nuxt.js build
- `.svelte-kit` - SvelteKit build

### IDE and Editor Files
- `.vscode` - VS Code settings
- `.idea` - IntelliJ/WebStorm settings
- `.git` - Git repository data

### System Files
- `.DS_Store` - macOS system files
- `Thumbs.db` - Windows thumbnail cache

### Cache and Temporary Files
- `.cache` - Cache directories
- `tmp`, `temp` - Temporary files
- `.tmp`, `.temp` - Temporary files
- `__pycache__` - Python cache
- `.pytest_cache` - Python test cache
- `.nyc_output` - Node.js coverage

### Environment Files
- `.env` - Environment files
- `.env.local` - Local environment files

### Test Coverage
- `coverage` - Test coverage reports

## Installation

### Prerequisites
- Go 1.25.1 or later

### Build from Source
```bash
git clone https://github.com/EldynC/Simple-Dev-Backup.git
cd Simple-Dev-Backup
go build -o filecopy hello.go
```

## Usage

```bash
./filecopy <source> <target>
```

### Examples

```bash
# Copy current directory to target folder
./filecopy . ./backup

# Copy specific directory
./filecopy ./source ./destination

# Copy with absolute paths
./filecopy /path/to/source /path/to/destination
```

### Command Line Arguments

- `<source>`: Source directory to copy from
  - `.` - Current directory
  - `./folder` - Relative path from current directory
  - `/absolute/path` - Absolute path
- `<target>`: Destination directory to copy to

## How It Works

1. **Path Resolution**: Converts relative paths to absolute paths
2. **Directory Walking**: Uses `filepath.Walk` to traverse the source directory
3. **Smart Filtering**: Checks each file/directory against the exclusion list
4. **Recursive Copying**: Creates directories and copies files maintaining the original structure
5. **Verbose Output**: Shows which files are being processed and which are being skipped

## Code Structure

```
filecopy.go
├── main()                    # Entry point, handles command line arguments
├── getPath()                 # Resolves relative and absolute paths
├── getSourceDirectories()    # Orchestrates the copying process
├── copyDir()                 # Recursively walks and copies directories
├── copyFile()                # Copies individual files
└── check()                   # Error handling utility
```

## Performance

- **O(1) Exclusion Lookup**: Uses Go maps for fast file/directory exclusion checking
- **Efficient File Copying**: Uses `ReadFrom` for optimal file copying performance
- **Memory Efficient**: Processes files one at a time without loading everything into memory

## Error Handling

The tool includes comprehensive error handling:
- File system errors during directory traversal
- File copying errors
- Path resolution errors
- Command execution errors

## Limitations

- Currently doesn't handle symbolic links
- No progress bar for large directory structures
- No dry-run mode
- No custom exclusion patterns (hardcoded list)


## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request
