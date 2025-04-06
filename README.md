# Personal Portfolio Website

This repository contains the source code for my personal portfolio website, built using **Go** and **HTMX**. The project showcases my work, skills, and projects in a clean and interactive design.

## Features

- **Dynamic Content Rendering**: Templates written in Go (`.templ` files) are used to dynamically generate HTML pages.
- **HTMX Integration**: Enables seamless interactivity without requiring full-page reloads.
- **JSON Data Handling**: Portfolio cards and other data are managed using JSON files.
- **Static Assets**: Includes images and JavaScript files for enhanced functionality.

## Project Structure

```
portfolio/
├── Body_templ.go       # Go code for rendering the body template
├── Body.templ          # HTML template for the body
├── Card_templ.go       # Go code for rendering individual cards
├── Card.templ          # HTML template for portfolio cards
├── cardsData.json      # JSON file containing data for portfolio cards
├── go.mod              # Go module file
├── go.sum              # Go dependency file
├── index_templ.go      # Go code for rendering the index page
├── index.templ         # HTML template for the index page
├── main.go             # Entry point for the application
├── portfolio           # Compiled binary (ignored in .gitignore)
├── readjson.go         # Utility for reading JSON data
├── Struct.go           # Struct definitions for the project
├── assets/             # Static assets
│   ├── htmx.min.js     # HTMX JavaScript library
│   └── images/         # Image assets
│       └── test.jpg    # Example image
├── handlers/           # HTTP handlers
│   └── test.go         # Example handler
└── tmp/                # Temporary files (ignored in .gitignore)
    └── main            # Temporary binary
```

## Prerequisites

- **Go**: Ensure you have Go installed on your system. You can download it from [golang.org](https://golang.org).
- **HTMX**: The project uses HTMX for interactive web components. The library is included in the `assets/` directory.

## Getting Started

1. Clone the repository:

   ```sh
   git clone https://github.com/imrun10/imrun10.github.io.git
   cd imrun10/imrun10.github.io/portfolio
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Run the application:

   ```sh
   go run main.go
   ```

4. Open your browser and navigate to `http://localhost:8080` to view the portfolio.

## Customization

- **Templates**: Modify the `.templ` files in the `portfolio/` directory to customize the HTML structure and design.
- **Data**: Update `cardsData.json` to change the content displayed on the portfolio cards.
- **Assets**: Add or replace images in the `assets/images/` directory and update the JavaScript in `assets/htmx.min.js` as needed.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [HTMX](https://htmx.org) for enabling interactive web components.
- The Go programming language for its simplicity and performance.
