from flask import Flask, request, render_template_string

app = Flask(__name__)

with open("index.html") as f:
    template = f.read()


@app.route("/")
def home():
    if request.args.get("c"):
        return render_template_string(
            template.replace("CHANGE ME", request.args.get("c"))
        )
    else:
        return "Hello, send someting inside the param 'c'!"


if __name__ == "__main__":
    app.run("0.0.0.0", 3333)
