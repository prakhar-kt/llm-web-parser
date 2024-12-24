# LLM Web Parser

A Python-based tool that combines web scraping and AI-powered summarization to generate concise, markdown-formatted summaries of web content.

## Features

- Web scraping using BeautifulSoup4
- Content cleaning and preprocessing
- AI-powered summarization using OpenAI's GPT models
- Markdown-formatted output
- Jupyter notebook implementation

## Prerequisites

- Python 3.11+
- OpenAI API key
- Required Python packages:
  - requests
  - beautifulsoup4
  - python-dotenv
  - openai
  - IPython

## Setup

1. Clone this repository
2. Create a conda env using "conda env create -f environment.yml"
3. Activate the env using "conda activate llms"
4. Create a `.env` file in the project root and add your OpenAI API key:
```
OPENAI_API_KEY=your-api-key-here
```

## Usage

The main functionality is implemented in `web_parser.ipynb`. The notebook provides two main functions:

- `summarize(url)`: Returns a markdown-formatted summary of the webpage
- `display_summary(url)`: Displays the formatted summary in a Jupyter notebook

Example usage:
```python
# Get summary as string
summary = summarize("https://example.com")

# Display formatted summary in notebook
display_summary("https://example.com")
```

## How It Works

1. The `Website` class handles web scraping:
   - Fetches webpage content using requests
   - Removes irrelevant elements (scripts, styles, images)
   - Extracts clean text content using BeautifulSoup

2. The summarization process:
   - Uses OpenAI's API to generate concise summaries
   - Formats output in markdown
   - Preserves important structural elements of the content

