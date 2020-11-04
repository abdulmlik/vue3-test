const express = require("express");
const path = require("path");
var history = require("connect-history-api-fallback");
let app = express();
app.use(history());
app.use(express.static(path.join(__dirname, "client-side", "dist")));

app.get("*", (req, res) => {
  res.sendFile("index.html");
});

/**
 * @var {number} PORT - The PORT for where server run, it is node Variable.
 */
const port = process.env.PORT || 3000;

app.listen(port, err => {
  if (err) console.log(err);
  else console.log("Server listening on Port %s", port);
});
