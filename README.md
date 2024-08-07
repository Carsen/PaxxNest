<h1>PaxxNest</h1>

PaxxNest is a universal package manager written in Golang. It iterates through a list of package managers and attempts to install, remove, or list packages. This tool aims to simplify package management across different systems.

<h4>Please Note: </h4>

<p>PaxxNest is very new and the development team is limited. It is expected that some functions may not work as intended. Please have patience as we continue to build. Thanks!</p>

<h2>Features:</h2>

- Universal Package Management: Supports multiple package managers.

- Easy to Use: Simple commands to install, remove, or list packages.

- Cross-Platform: Works on various operating systems.

<h2>Installation:</h2>

1. Ensure GO is installed. If not, install it here: https://go.dev

2. Clone the Repository:

`git clone <https://github.com/Carsen/PaxxNest.git>`

3. Navigate to the Source Directory:

`cd PaxxNest/src`

4. Run the Program:

`go run .`

<h2>Usage:</h2>

<h4>List/Install/Remove</h4>

Follow the in-app prompts to select your operation and package.

<h4>Config</h4>

Configuration is done directly in src/main.go. There are helpful comments that indicate which line allows each package manager (All are allowed by default).

To add or remove a package manager from the list, simply comment out the line for that manager:

`mgr.AddManager("brew", Manager.BrewMan{})` :Active <--> Inactive: `// mgr.AddManager("brew", Manager.BrewMan{})`

<h2>Contributing</h2>

Contributions are welcome! Please fork the repository and submit a pull request.
