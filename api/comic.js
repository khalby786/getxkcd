const fetch = require('node-fetch');

module.exports = async (req, res) => {
  const { comic } = req.query;
  
  if (comic == "" || comic == " ") {
    let req = await fetch("http://xkcd.com/info.0.json");
    let res = await req.json();
    res.sendStatus(200).json(res);
  } else {
    let req = await fetch(`http://xkcd.com/${comic}/info.0.json`);
    let res = await req.json();
    res.sendStatus(200).json(res);
  }
}