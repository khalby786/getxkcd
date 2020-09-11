const fetch = require('node-fetch');

exports.handler = async (event) => {
  
  const { num } = event.queryStringParameters;
  
  if (!num || num == "" || num == " " || num == "latest") {
    let req = await fetch("http://xkcd.com/info.0.json");
    let res = await req.json();
      return {
      statusCode: 200,
      body: JSON.stringify(res)
    }
  } else {
    let req = await fetch(`http://xkcd.com/${num}/info.0.json`);
    let res = await req.json();
    return {
      statusCode: 200,
      body: JSON.stringify(res)
    }
  }
}