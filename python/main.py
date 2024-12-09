from flask import Flask, request, render_template_string, render_template

app = Flask(__name__)

with open("index.html") as f:
    template = f.read()


@app.route("/", methods=["GET", "POST"])
def index():
    if request.method == "POST":
		username = request.form.get("username")
		comment = request.form.get("comment")
		if not username or not comment:
            		return "Brak nazwy uzytkownika lub komentarza"
		output = render_template_string(f"{username}: {comment}")
		return render_template("index.html", output=output)
	
	return render_template("index.html")

if __name__ == "__main__":
    app.run(debug=True)
