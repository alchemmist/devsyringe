<h2><img src="./media/logo.svg" alt="Favicon Preview" width="130" align="center"> Devsyringe</h2>

[![Github](https://img.shields.io/badge/alchemmist%2Fdevsyringe-blue?logo=github&label=github&color=blue)](https://github.com/alchemmist/devsyringe)
![Last commit](https://img.shields.io/github/last-commit/alchemmist/devsyringe?style=flat) ![Stars](https://img.shields.io/github/stars/alchemmist/devsyringe?style=flat)
![Forks](https://img.shields.io/github/forks/alchemmist/devsyringe?style=flat)
![License](https://img.shields.io/github/license/alchemmist/devsyringe?style=flat)
![Contributors](https://img.shields.io/github/contributors/alchemmist/devsyringe?style=flat)
![Go](https://img.shields.io/badge/1.24-default?label=Go)
[![Build](https://github.com/alchemmist/devsyringe/actions/workflows/build.yaml/badge.svg?branch=main)](https://github.com/alchemmist/devsyringe/actions/workflows/build.yaml)

CLI for inject dynamic values into code/config files from external commands using a declarative YAML config.

## Why?

In many projects, you have parameters like web domains, API tokens, titles, passwords, or UUIDs.
When working in a programming environment, we often use `.env` files and helper libraries:

```python
load_dotenv(find_dotenv())
host = get_env("HOST")
```

This works well when you have a compiler, package manager, or libraries.
But sometimes you need to work with simple HTML files, static configs, or scripts --- where such tools are not available.

Devsyringe solves this problem: it can fetch values from any source and inject them into any target files, making your parameters dynamic without templates.

## Demo

<p align="center">
    <img src="./media/demo.gif" alt="Demo GIF" width="700" style="border-radius: 15px; box-shadow: 0px 0px 40px rgba(0, 0, 0, 0.3)">
</p>

In this demo we use this compose file:

```yaml
serums:
  https_host:
    source: ssh -R 80:localhost:3000 serveo.net
    mask: https?://[^\s]*?\.serveo\.net\b
    targets:
      .env:
        path: ./.env
        clues: ["HOST"]
```

Devsryinge run command, cut vlue from ouput with mask and replace value in target `.env` file under the mask to new. Very simple!

## Features

- Commands for preocess cofnig and controll processes, which need to be got injectable values. See [Usage](#usage).

- <p>Use <code>dsy [command] --help</code> for more information about a command.</p>
- Use `dsy` without commands or flags for run TUI with table of processes.
  - In TUI you can see table of processes, which devsyringe run. In table you can see status of process (`stoped` or `active`), command for running, `PID` and title.
  - Here you can controll processes with hotkeys. Press `?` for see help. You can stop process, delete process, see full ouput of process and so on.

**TUI**:

<p align="center">
    <img src="./media/tui-demo.gif" alt="Demo GIF" width="700" style="border-radius: 15px; box-shadow: 0px 0px 40px rgba(0, 0, 0, 0.3)">
</p>

<h2 id="usage">Usage</h2>
Available commands for process config and controll processes:
<table>
  <thead>
    <tr>
      <th>Command</th>
      <th>Description</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td><code>completion</code></td>
      <td>Generate the autocompletion script for the specified shell</td>
    </tr>
    <tr>
      <td><code>delete</code></td>
      <td>If not stopped, stop. Then delete process with [title] from list and delete all logs.</td>
    </tr>
    <tr>
      <td><code>help</code></td>
      <td>Help about any command</td>
    </tr>
    <tr>
      <td><code>inject</code></td>
      <td>Start an injection based on some config</td>
    </tr>
    <tr>
      <td><code>list</code></td>
      <td>Show dynamic list of running processes</td>
    </tr>
    <tr>
      <td><code>logs</code></td>
      <td>Show logs from process with [title]</td>
    </tr>
    <tr>
      <td><code>stop</code></td>
      <td>Stop process with [title], but save logs and keep in list</td>
    </tr>
  </tbody>
</table>

## Configuration

Devsyringe uses a **YAML file** (usually named `syringe.yaml`) to describe what commands to run,
how to extract values from their output, and where to inject them.

### Top-level structure

```yaml
serums:
  <title>:
    source: <command>
    mask: <regex>
    max-timeout: <seconds>
    targets:
      <alias>:
        path: <file-path>
        clues: [<clue1>, <clue2>, ...]
```

`serums` — the root section. Each child key under serums defines a serum (a dynamic value you want to manage).

`<title>` — unique name of the serum (e.g. localtunnel, https_host, ping-test).

| Option       | Type      | Description                                                                 | Example |
|--------------|-----------|-----------------------------------------------------------------------------|---------|
| `source`     | string    | Command to run. Its output will be scanned to extract values.               | `lt --port 80` |
| `mask`       | regex     | Regular expression to capture the desired part of the output. If omitted, the whole output is used. | `https://[a-z0-9\-]+\.loca\.lt` |
| `max-timeout`| int (sec) | Maximum time to wait for `source` output before failing.                    | `5` |
| `targets`    | map       | Where to inject the extracted value(s). Each target has its own settings.   | see below |


**Targets:**
Each serum may inject values into one or multiple files.

```yaml
targets:
  <alias>:
    path: <file-path>
    clues: [<clue1>, <clue2>, ...]
```

`<alias>` — arbitrary name for readability (e.g. env, js, config).

`path` — relative or absolute path to the file to update.

`clues` — list of words/markers that help Devsyringe find the right place in the file to replace.

### Examples

**Exampel 1** -- Localtunnel URL injection:

```yaml
serums:
  localtunnel:
    source: lt --port 80
    mask: https://[a-z0-9\-]+\.loca\.lt
    targets:
      .env:
        path: test/.env
        clues: ["HOST"]
      js:
        path: test/static.js
        clues: ["const", "url"]
```

**Example 2** -- SSH serveo tunnel with timeout:

```yaml
serums:
  https_host:
    source: ssh -R 80:localhost:3000 serveo.net
    mask: https?://[^\s]*?\.serveo\.net\b
    max-timeout: 5
    targets:
      .env:
        path: ./.env
        clues: ["HOST"]
      js:
        path: ./static/scanner.js
        clues: ["const", "url"]
```

## Contribution

If you run into any issues or have any suggestions, open a Pull Request that includes the updates and I'll review/comment/merge it as soon as I can. If you don't have enough time or don't know how to fix the issue, submit an Issue and provide as much detail as you can.

For any questions, mail at: <a href="mailto:anton.ingrish@gmail.com">anton.ingrish@gmail.com</a>
