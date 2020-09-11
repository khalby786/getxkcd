const fetch = require('node-fetch');

module.exports = async (request, response) => {
  
  const { comic } = request.query;
  
  if (comic == "" || comic == " ") {
    let req = await fetch("http://xkcd.com/info.0.json");
    let res = await req.json();
    response.status(200).json(res);
  } else {
    let req = await fetch(`http://xkcd.com/${comic}/info.0.json`);
    let res = await req.json();
    response.status(200).json(res);
  }
}