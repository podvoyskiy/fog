### Fuzzy Command History Search

**Linux/Bash only**

### Installation
1. Download binary:
    ```bash
    wget https://github.com/podvoyskiy/fog/releases/latest/download/fog -O ~/.local/bin/fog
    ```

2. Make executable:
   ```sh
   chmod +x ~/.local/bin/fog
   ```

3. Add to `~/.bashrc`:
    ```bash
    f() { if [ $# -gt 0 ]; then ~/.local/bin/fog "$@"; else eval "$(~/.local/bin/fog)"; fi }
    ```

4. Reload bashrc:
    ```bash
    source ~/.bashrc
    ```

### Usage
```bash
f           # Interactive search
f --help    # Show options
```