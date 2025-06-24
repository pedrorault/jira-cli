---
name: "example-task"
type: "Task"
parent: ""
summary: "Implement the new login page"
priority: "High"
reporter: "jira"
assignee: "jira"
labels:
  - "frontend"
  - "login"
components:
  - "Web"
  - "Auth"
fix-versions:
  - "v1.2.0"
affects-versions:
  - "v1.2.0"
original-estimate: "4h"
custom-fields:
  environment: "staging"
debug: false
---
# Awesome Title
## Awesome subtitle

> Blockquote: A wise man once said, WTF!!!!

[A link](http://example.com/)

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

![](https://go.dev/images/gophers/ladder.svg|width=60,height=188)

|a  |b  |c  |
|---|---|---|
|1  |2  |3  |
|4  |5  |6  |

We can also use Jira flavored markdown to create jira specific entities like panel.

{panel:bgColor=#e3fcef}
Success panel!

**strong text**
~~strikethrough text~~ 
{panel}

{panel:title=Error Lists|bgColor=#ffebe6}
1. Item A
2. Item B
3. Item C
{panel}

