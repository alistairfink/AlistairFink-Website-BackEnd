你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
# Personal Website Backend

This is probably around version 4 or 5 of the back end for this website. I created this based off my original go based backend for my website which I used as a POC combined it with a stronger knowledge of web architecture and design patterns. I was unhappy with the original stack I had chosen for my website as the architecture was a mess and it was very poorly written and not very maintable. I was planning on replacing this it at some point but had never really gotten around to it. Eventually my VPS provider went backrupt without warning and I lost my entire database. This prompted me to simply get around to redoing the entire stack using go, a language I've grown to love in the past 6 months. In the end the new stack is a React front end, go back end, and Postgres DB.

## Architecture
* React Frontend
* Go/Chi Backend (Running in Docker Container)
* Postgres Database (Running in Docker Container)

Note: Before you ask the stuff in the config file has been changed in production so don't even try.
