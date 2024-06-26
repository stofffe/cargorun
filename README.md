# cargorun

cargorun is a simple command line tool for running cargo projects with shell completions

## Shell completion setup

### Zsh
Add the following command to your ```.zshrc``` (Recommended)  
Alternatively it can be ran straight in the terminal

```zsh
source <(cargorun completion zsh)
```

### Other
Completions can be generated for ```bash```, ```powershell``` and ```fish```, the steps should be similar to the ```zsh``` setup above

## Usage
### Run
To run examples use
```zsh
cargorun example <Name>
```
To run binaries use
```zsh
cargorun bin <Name>
```

### Additional arguments
Additional arguments can be passed using ```--```
```zsh
cargorun example <Name> -- <Arg1> <Arg2>
```

### Examples
```zsh
cargorun example calculator -- add 1 2
-->
cargo run --example calculator add 1 2
```
```zsh
cargorun bin game -- --release
-->
cargo run --bin game --release
```
# cargorun
