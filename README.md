# Building Slack Bots in Golang

This project is a Slackbot in Golang; it uses the [slack-go](https://github.com/slack-go/slack) library and communicates with slack using the [socket mode](https://api.slack.com/apis/connections/socket).

## Test the project

Create a file `test_slack.env` with the following variables:

```
SLACK_BOT_TOKEN=xoxb-xxxxxxxxxxx
SLACK_APP_TOKEN=xapp-1-xxxxxxxxx
```

Run the application

```
go run main.go
```

## Showcases

1. Greetings
2. App Home
3. Slack commands

### Greetings
![](./docs/assets/greeting.gif)

![](./out/controllers/greetingController/greetingController.png)

Tutorial 2 : [Properly Welcome Users in Slack with Golang using Socket Mode](https://levelup.gitconnected.com/properly-welcome-users-in-slack-with-golang-using-socket-mode-9a206d30a34a?sk=24fb8c44c1128cc3ef366d51d4a48812)

### App Home

![](./docs/assets/apphome_completed.gif)

![](./out/controllers/appHomeController/appHomeController.png)

Tutorial 1 : [Build Your Slack App Home in Golang Using Socket Mode](https://betterprogramming.pub/build-a-slack-app-home-in-golang-using-socket-mode-aff7b855bb31?sk=b8b1f7f3c03972793b26bec02dc3d2cc)

### Slash Commands

![](./docs/assets/slachcommandrocket.gif)

Tutorial 3: [Implement Slack Slash Command with Golang using Socket Mode](https://levelup.gitconnected.com/implement-slack-slash-command-in-golang-using-socket-mode-ac693e38148c?sk=33e90a65aded42cd4737ff6a137762cc)
## References
* [Building a home for your app 🏡](https://api.slack.com/tutorials/app-home-with-modal)
