# Projector

I made projector to learn the basics of Go. It is a simple CLI tool that allows users to store environment variables to a given directory path.

## Usage

This CLI application supports the following 3 commands.

### Adding a Value
To add a value to the current working directory, type `projector add` followed by the variable name and the value to store


`projector add <key> <value>`

### Reading a Value
Once you have added a value, you can read its content by using `projector <key>` which will print only the given variable, or `projector` which will print all environment variables for the current wrking directory

### Removeing a Value
To remove a variable from the current working directory, type `projector rem <key>`
