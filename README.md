<h1 align="center">
  GETXKCD
</h1>

An XKCD API built using serverless functions with CORS enabled, hosted on Vercel.

### Latest Comic

```
https://getxkcd.now.sh/api/comic?num=latest
```

### Specific Comic

```
https://getxkcd.now.sh/api/comic?num=COMIC_NUMBER
```

```
https://getxkcd.now.sh/api/comic?num=2355
```

### Sample Response

GETXKCD fetches data from the [original XKCD JSON API](https://xkcd.com/json.html) and returns the result. *Because the original XKCD JSON API does not support CORS*.

The JSON object response received is not always the same, properties may be added or not present depending on the content of the comic.

**Example response (#2355)**

```
{
  "month": "9",
  "num": 2355,
  "link": "",
  "year": "2020",
  "news": "",
  "safe_title": "University COVID Model",
  "transcript": "",
  "alt": "I admit this is an exaggeration, since I can think of at least three parties I attended while doing my degree, and I'm probably forgetting several more.",
  "img": "https://imgs.xkcd.com/comics/university_covid_model.png",
  "title": "University COVID Model",
  "day": "4"
}
```

