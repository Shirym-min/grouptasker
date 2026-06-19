# gpx GroupTasker

A lightweight task runner for repetitive command workflows.

`gpx` allows you to define custom task aliases in a YAML configuration file and execute multiple commands with a single command.

## Features

- Run multiple commands with a single alias
- Store tasks in a simple YAML file
- Support argument placeholders (`{{1}}`, `{{2}}`, ...)
- Interactive task management
- Cross-platform support (macOS, Linux, Windows)

## Installation

### Homebrew

Coming Soon

### Build from Source

```bash
git clone https://github.com/Shirym-min/grouptasker.git
cd grouptasker
go build -o gpx
```

## Configuration

The configuration file is automatically created on first launch.

### macOS

```text
~/Library/Application Support/gpx/gpx.yaml
```

### Windows

Coming Soon

### Example

```yaml
tasks:
  build:
    - npm run build

  commit:
    - git add .
    - git commit -m "{{1}}"

  deploy:
    - git push origin main
```

## Usage

Run a task:

```bash
gpx build
```

Pass arguments:

```bash
gpx commit "Update README"
```

or Pass arguments in a dialogue format:

```bash
gpx commit
>> Input task 1 git commit -m "{{1}}" : Update README
```

Example expansion:

```yaml
commit:
  - git commit -m "{{1}}"
```

↓

```bash
git commit -m "Update README"
```

And you can pass multiple arguments

```bash
gpx echo2 "Hi" "I'm Shirym-min!"
```
or
```bash
gpx echo2
>> Input task 1 echo {{1}} : Hi
>> Input task 2 echo {{2}} : I'm Shirym-min!
```

↓

```bash
>>Hi
>>I'm Shirym-min!
```

```yaml
echo2:
  - echo {{1}}
  - echo {{2}}
```

## Task Management

List tasks:

```bash
gpx list
```

Add a task:

```bash
gpx config add
>> 
```

Example:

```bash
gpx config add
>>Task name : addandcommit
>>Command #1 : git add .
>>Command #2 : git commit -m "{{1}}"
>>Added Task
```

Delete a task:

```bash
gpx config delete <task-name>
```

Example:

```bash
gpx config delete commit
```



## License

MIT License
