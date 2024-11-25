import fs from "fs";
import express from "express";
import handlebars from "handlebars";

const html = fs.readFileSync("index.html", "utf-8");
const app = express();

app.use(express.static("public"));
app.use(express.urlencoded({ extended: false }));
app.get("/", (req, res) => {
  res.send(
    req.query.template
      ? handlebars.compile(html.replace("CHANGE ME", req.query.template))()
      : html
  );
});
app.listen(4444);
