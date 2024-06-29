---
title: sight AI Interview
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# sight AI Interview

Base URLs:

# Authentication

# Default

## POST gateway_compute

POST /compute

> Body Parameters

```json
{
  "numbers": [
    10,
    5,
    2
  ],
  "ast": {
    "type": "operation",
    "value": "+",
    "left": {
      "type": "number",
      "value": 10
    },
    "right": {
      "type": "operation",
      "value": "*",
      "left": {
        "type": "number",
        "value": 5
      },
      "right": {
        "type": "number",
        "value": 2
      }
    }
  }
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|

> Response Examples

> 200 Response

```json
{}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|Inline|

### Responses Data Schema

# Data Schema

