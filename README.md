# BluePrint

BluePrint CLI

## Description

BluePrint is a command-line tool written in Go that simplifies the process of creating GitHub repositories with predefined directory structures and files. It allows developers to generate blueprint repositories directly from the command line, saving time and effort in project setup.

With BluePrint, you can define custom blueprints that serve as templates for different types of projects. These blueprints can include directories, files, and even pre-configured settings such as licenses, README templates, and more. By utilizing BlueprintCLI, developers can quickly create new repositories with consistent structures and configurations, ensuring project consistency and accelerating the initial setup process.

## Features

- Generate GitHub repositories with predefined directory structures and files.
- Create custom blueprints for different project types or frameworks.
- Include commonly-used files such as README templates, license files, or configuration files in the blueprints.
- Define placeholder variables in blueprints to dynamically populate values during repository creation.
- Interact with the GitHub API to create repositories directly from the command line.
- Simple and intuitive command-line interface.

## Installation

To install BlueprintCLI, follow these steps:

Ensure you have Go version 1.16 or higher installed on your system.
Open a terminal and run the following command to download and install BlueprintCLI:
Copy

```
go install github.com/mrbomber0x001/blueprintcli
```

## Usage

To use BlueprintCLI, run the following command:

Copy
blueprintcli create [repositoryName] --blueprint=[blueprintName]
[repositoryName] is the name of the GitHub repository you want to create.
[blueprintName] is the name of the blueprint to use for repository creation.

## Documentation

Complete documentation for BlueprintCLI can be found <a href="#">here</a>.

## Contributing

To contribute to `BluePrint`, follow these steps:

1. Fork the repository on GitHub.
2. Clone the forked repository to your local machine.
3. Set up a development environment by following the instructions in the README.
4. Make your changes and additions.
5. Run tests using go test and ensure all tests pass.
6. Commit your changes and push them to your forked repository.
7. Create a pull request on the main repository.

## License

BlueprintCLI is distributed under the MIT License.

## Authors

- Yousef Meska - <a href="mailto:yousefmeska123@gmail.com">yousefmeska123@gmail.com</a>

## Acknowledgments

Optional section to acknowledge any external resources, libraries, or individuals who have contributed to the project or influenced its development.

## TODO and Log

- [x] Creating private/public repos
- [ ] Blueprinting repo with default files required for your project depending on the tech stack
- [ ] check the token -> cache, for the first time, to set in `.env` files
- [ ] Bundles the application in single command line tool in the system
- [ ] Clone the repo after creating it
- [ ] Make it as downloadable from git
