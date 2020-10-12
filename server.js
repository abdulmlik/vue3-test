const express = require("express");
const app = express();

app.all("/*", function(req, res) {
  res.sendFile("index.html", { root: __dirname + "/dist" });
});

const port = process.env.PORT || 3000;

app.listen(port, () => {
  console.log("Server listening on Port %s", port);
});
