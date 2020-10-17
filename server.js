/**
 * @file representing a server side.
 *
 * run with node server.js
 */

const express = require("express");
const path = require("path");
const app = express();

/**  set the default path */
app.use(express.static(path.join(__dirname, "dist")));

/** all Route do this */
app.all("*", function(req, res) {
  const options = null; // like this { root: path.join(__dirname + "dist") }.

  res.setHeader("content-type", "text/html");
  res.sendFile("index.html", options, function(error, data) {
    if (error) {
      res.writeHead(404);
      res.write("File Not Found");
      // res.status(200).redirect(`/#${req.url}`);
    } else {
      res.write(data);
    }
    res.end();
  });
});

/**
 * @var {number} PORT - The PORT for where server run, it is node Variable.
 */
const port = process.env.PORT || 3001;

app.listen(port, (err) => {
  if (err) console.log(err);
  else console.log("Server listening on Port %s", port);
});
