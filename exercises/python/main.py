import os
import secrets
from flask import Flask, render_template_string, request, render_template


PORT = 3333


app = Flask(__name__)


def filter(user_input):
    A = ["{", "}", "%", "(", ")", "[", "]", "<", ">"]
    for i in A:
        user_input = user_input.replace(i, "")
    return user_input


@app.route("/", methods=["GET", "POST"])
def index():
    if request.method == "POST":
        username = filter(request.form.get("username"))
        comment = request.form.get("comment")
        if not username or not comment:
            return "Brak nazwy użytkownika lub komentarza"
        output = render_template_string(f"{username}: {comment}")
        return render_template("index.html", output=output)
    return render_template("index.html")


def main():
    app.config["SECRET_KEY"] = secrets.token_bytes(8).hex()
    os.environ["FLAG"] = f"SSTI{{{secrets.token_bytes(16).hex()}}}"
    print(f"server is running on port {PORT}")
    app.run(host="0.0.0.0", port=PORT)


if __name__ == "__main__":
    main()
