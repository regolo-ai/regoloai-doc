# Regolo AI Docs

Welcome to the **Regolo AI Docs** for our project! This repository is structured to utilize **MkDocs** with the **MkDocs Material** theme for generating and deploying a beautiful static documentation site.

## Table of Contents

1. [Getting Started](#getting-started)
2. [Installation](#installation)
3. [Creating a Virtual Environment](#creating-a-virtual-environment)
4. [Building the Documentation](#building-the-documentation)
5. [Deploying the Documentation](#deploying-the-documentation)
6. [Adding New Document Pages](#adding-new-document-pages)
7. [Adding Images and Media](#adding-images-and-media)

## Getting Started

To work with this repository, you'll need to download or clone it from GitHub. Once you have it on your local machine, navigate to the `regolo-doc` directory which contains the `docs` folder. This is where you'll run all MkDocs commands.

## Installation

To set up your environment and install the necessary dependencies, please follow these steps:

1. **Clone the Repository:**

   ```bash
   git clone git@github.com:regolo-ai/regoloai-doc.git
   ```

2. **Navigate to the `regolo-doc` Directory:**

   ```bash
   cd regoloai-doc/regolo-doc
   ```

3. **Install Requirements:**

   It's recommended to use a virtual environment (as described below) before installing the requirements.

   ```bash
   pip install -r requirements.txt
   ```

   This will install MkDocs and the MkDocs Material theme.

## Creating a Virtual Environment

To ensure a clean and isolated environment for your documentation work, it's best to use Python's virtual environments. Here’s how you can create one:

1. **Create a Virtual Environment:**

   ```bash
   python -m venv env
   ```

2. **Activate the Virtual Environment:**

   - On Windows:

     ```bash
     .\env\Scripts\activate
     ```

   - On macOS and Linux:

     ```bash
     source env/bin/activate
     ```

3. **Once activated**, you can proceed with installing the requirements as mentioned above.

## Building the Documentation

With the virtual environment activated and dependencies installed, you can now build the documentation:

1. **Serve the Documentation Locally:**

   ```bash
   mkdocs serve
   ```

   This command will launch a local server at `http://127.0.0.1:8000` where you can view the documentation live. Any changes you make will automatically refresh.

## Deploying the Documentation

To deploy the documentation site to GitHub Pages, use the following command:

```bash
mkdocs gh-deploy
```

This command builds the site and pushes it to the `gh-pages` branch of your repository, making it accessible via GitHub Pages.

## Adding New Document Pages

To add new pages to the documentation:

1. **Create a Markdown File:**

   Add a new Markdown (`.md`) file inside the `docs` folder. For example, `new-page.md`.

2. **Reference the New Page in `mkdocs.yml`:**

   Open the `mkdocs.yml` file and add an entry under the `nav` section to reference your new page:

   ```yaml
   nav:
     - Home: index.md
     - New Page: new-page.md
   ```

   Ensure the indentation is correct and consistent.

## Adding Images and Media

To include images and other media files in your documentation:

1. **Place Media Files:**

   Put your images and media in the `img` folder inside the `docs` directory (e.g., `docs/img`).

2. **Reference Media in Markdown:**

   Use the following syntax in your Markdown files to reference images:

   ```markdown
   ![Alt text](img/your-image-file.png)
   ```

   Replace `your-image-file.png` with the actual file name.

By following these instructions, you can easily expand and enhance your documentation with new content and media. If you have any questions or encounter issues, feel free to reach out via the repository's issue tracker.