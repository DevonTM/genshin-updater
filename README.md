# Genshin Impact Patch Downloader

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/DevonTM/genshin-updater/blob/main/LICENSE)

## Description

Genshin Impact Patch Downloader is a Go application for efficiently downloading Genshin Impact update patches using aria2.

## Features

- Fetches update information from a JSON API
- Allows selection and downloading of specific update files
- Multi-threaded download for faster speed by using aria2
- Easy-to-use command-line interface

## Compiling

1. Install Go (if not already installed) by following the official installation instructions: [https://golang.org/doc/install](https://golang.org/doc/install)
2. Clone the repository: `git clone https://github.com/DevonTM/genshin-updater.git`
3. Navigate to the project directory: `cd genshin-updater`
4. Build the project: `make build`

## Usage

1. Download the latest release from the repository's [Releases](https://github.com/DevonTM/genshin-updater/releases) page.
2. Extract the downloaded archive to a directory of your choice.
3. Navigate to the extracted directory and run the `genshin-updater` executable.
4. Follow the on-screen prompts to select the desired update files.
5. After the download completes, move the downloaded files in the `patch` directory to the Genshin Impact installation directory where the `GenshinImpact.exe` file is located.

## License

This project is licensed under the [MIT License](LICENSE).

## Contributing

Contributions to this project are welcome. Feel free to fork the repository and submit a pull request.

## Acknowledgements

This application was inspired by the need for faster game updates in Genshin Impact.
