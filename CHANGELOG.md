# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Released]

## [1.3.0] - 2023-03-09

### Changed

- Added go_console.Command to simplify complex command
- Renaming of go_console.Cli to go_console.Script
- Struct definition of go_console.Script and go_console.Command

## [1.2.0] - 2023-03-09

### Changed

- Added Table rendering
- Added minimalistic documentation

### known limitations

- rowspan not supported yet 

## [1.1.0] - 2023-03-03

### Changed

- Added Question Helper for direct user input
- Added minimalistic documentation

### known limitations

- Question Helper does not support windows yet (due to hidden input)

## [1.0.0] - 2019-04-20

### Changed

- Added Null, Buffered, Chan and Console Output Interface
- Added minimalistic documentation

### Fixed

- Fix Styling on deep nested files