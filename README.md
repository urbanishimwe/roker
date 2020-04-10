<h1 align="center"><code>roker</code></h1>

<div align="center">
  <sub>Created by <a href="https://github.com/urbanishimwe">Urban Ishimwe</a></sub>
</div>

Route ngrok URL to static URL


## Why

"I am overwhelmed by configuring URL on API service like Slack everytime I start my Ngrok server".
Then, this might be the solution.

## How?
1. Fork this repository
2. create an account on Heroku

## Setup
1. deploy the forked version on your Heroku
2. Config your URL by setting an heroku environment: `URL`=`secure-ngrok-url`
3. Config your SECRET by setting an heroku environment variable: e.g `SECRET`=`IFFEnGFS4H`

*you might need to restart your app everytime you change this environments*

## Final

Now, you can just use your Heroku link that will forward to your ngrok dynamic address.

If you want to change your url without changing the URL env variable and restarting your app. You can use the route:

```
https://HEROKU_URL/update_ngrok&secret=SECRET&url=NGROK_URL
```


You can, for example, open with your browser: https://myheroku/update_ngrok&secret=IFFEnGFS4H&url=https://b1dd4477.ngrok.io/

or 
From your terminal

```bash
curl "your_heroku/update_ngrok?secret=IFFEnGFS4H&url=https://b1dd4477.ngrok.io"
```

## Contributions

Feel free to contribute, PR and issues are very welcome 🙏🙏🙏🙏

