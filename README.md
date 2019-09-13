[] To install binary executable into the bin and run the program. Follow these comands:

- Switch to bash terminal by running: `bash`
- Set up the GOBIN Path by running: `source $HOME/.bash_profile`
- Then switch back to fish terminal by running: `fish`
- Run binary file by typing the exacutable's name into the terminal: ex, `Server`

[] To set up GOBIN with fish and set up Path to run executable from terminal. Follow these comands:

- Nav and open "config.fish" ex: "~/.config/fish/config.fish". Now add and save this line to the file: `set -x GOBIN ~/go/bin`
- Set fish PATH by running: `echo "set -gx PATH \$PATH ~/go/bin" >> ~/.config/fish/config.fish`

