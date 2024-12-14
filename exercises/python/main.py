from flask import Flask, render_template_string, request, render_template
import os


app = Flask(__name__)
app.config["SECRET_KEY"] = os.urandom(24)


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
            return "Brak nazwy uzytkownika lub komentarza"
        output = render_template_string(f"{username}: {comment}")
        return render_template("index.html", output=output)
    return render_template("index.html")


if __name__ == "__main__":
    app.run(debug=True)
