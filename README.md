# getxkcd

[![Deploy with Vercel](https://vercel.com/button)](https://vercel.com/new/clone?repository-url=https%3A%2F%2Fgithub.com%2Fkhalby786%2Fgetxkcd)


### get the latest comic

https://getxkcd.now.sh/latest<br />
https://getxkcd.now.sh/0<br />

#### example response

```json
{
  "month": "1",
  "num": 2886,
  "link": "",
  "year": "2024",
  "news": "",
  "safe_title": "Fast Radio Bursts",
  "transcript": "",
  "alt": "Dr. Petroff has also shown that the Higgs boson signal was actually sparks from someone microwaving grapes, the EHT black hole photo was a frozen bagel someone left in too long, and the LIGO detection was just someone slamming the microwave door too hard.",
  "img": "https://imgs.xkcd.com/comics/fast_radio_bursts.png",
  "title": "Fast Radio Bursts",
  "day": "26"
}
```

### get a comic by number

https://getxkcd.now.sh/1481

#### example response

```json
{
  "month": "2",
  "num": 1481,
  "link": "",
  "year": "2015",
  "news": "",
  "safe_title": "API",
  "transcript": "((This is a faux-screenshot of a technical document))\n[[A figure sits at a computer upon a desk, apparently engrossed in the document which we now see before us.]]\nTITLE: API GUIDE\nRequest URL Format: domain\nuser\nitem\nServer will return an XML document which contains the requested data and documentation describing how the data is organized spatially.\nAPI KEYS: To obtain API access, contact the x.509-authenticated server and request an ECDH-RSA TLS key...\n\nCaption: If you do things right, it can take people a while to realize that your \"API documentation\" is just instructions for how to look at your website.\n\n{{Title text: ACCESS LIMITS: Clients may maintain connections to the server for no more than 86,400 seconds per day. If you need additional time, you may contact IERS to file a request for up to one additional second.}}",
  "alt": "ACCESS LIMITS: Clients may maintain connections to the server for no more than 86,400 seconds per day. If you need additional time, you may contact IERS to file a request for up to one additional second.",
  "img": "https://imgs.xkcd.com/comics/api.png",
  "title": "API",
  "day": "2"
}
```

## notes

- you can also get the latest comic by sending a request to https://getxkcd.now.sh/, but the request header should have `Accept` as `application/json`


  ```bash
  curl -H "Accept: application/json" "https://getxkcd.now.sh/""
  ```

- the API is [backwards compatible](https://xkcd.com/1172/) with the previous versions which used https://getxkcd.now.sh/api/comic?num=latest to get the latest comic and https://getxkcd.now.sh/api/comic?num=1172 to get a particular comic
- if you get an HTML response from the API, you have to set your API client to handle redirects. alternatively, replace `.now.sh` domains with `.vercel.app` to avoid this pitfall in places like `curl`
- `vercel.json` handles all the routing, so make sure you copy it verbatim if you want good things to happen
- in an effort to learn and use a new programming language, i have used go which is unnecessarily complicating things but *✨ performance ✨* 

## closing thoughts

made by [khaleel](https://khaleelgibran.com) when he was feeling lonely