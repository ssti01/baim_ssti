import crypto from "crypto";
import fs from "fs";
import express from "express";
import handlebars from "handlebars";

const PORT = 4444;

const html = fs.readFileSync("index.html", "utf-8");
const app = express();

app.use(express.static("public"));

app.get("/", (req, res) => {
  const userInput = req.query.template ?? "Domyślny opis";
  const templateStr = html.replace("CHANGE ME", userInput);

  const template = handlebars.compile(templateStr);
  const result = template({});

  res.send(result);
});

process.env.FLAG = `SSTI{${crypto.randomBytes(16).toString("hex")}}`;

app.listen(PORT, () => {
  console.log(`server is running on port ${PORT}`);
});
