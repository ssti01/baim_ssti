import fs from "fs";
import express from "express";
import handlebars from "handlebars";

const hbs = fs.readFileSync("index.hbs", "utf-8");
const app = express();

app.use(express.static("public"));
app.use(express.urlencoded({ extended: false }));

app.get("/", (req, res) => {

  const UserInput = req.query.template || "Domyślny opis";
  const html = hbs.replace("{{ description }}", UserInput);
  
  const template = handlebars.compile(html);
  const result = template({});

  res.send(result);
});

app.listen(4444, () => {
  console.log("Serwer działa na porcie 4444");
});
