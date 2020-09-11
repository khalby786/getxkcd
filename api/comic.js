const fetch = require('node-fetch');

module.exports = async (request, response) => {
  
  const { num } = request.query;
  
  if (!num || num == "" || num == " " || num == "latest") {
    let req = await fetch("http://xkcd.com/info.0.json");
    let res = await req.json();
    response.status(200).json(res);
  } else {
    let req = await fetch(`http://xkcd.com/${num}/info.0.json`);
    let res = await req.json();
    response.status(200).json(res);
  }
}