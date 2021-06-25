# Onboarding Mattermost Plugin

![Logo](assets/logo.png)

# Description

#### FR

Projet de workshop réalisé par 5 étudiants de l'EEMI. Réalisation d'un plugin pour Mattermost permettant un suivi optimal des projets par le biai d'un chatbot. Le chatbot permet un onboarding automatisé, grâce à cela le manager pourra s'assurer de l'avancée de son équipe et de mettre en évidence les points bloquants.

Notre projet a pour particularité d'être très flexible et adadptable à toutes les situations ou projets. Il peut ainsi être utilisé de façon ludique ou éducative.

#### EN

Workshop project done by 5 students at [EEMI](https://eemi.com). We've created a [Mattermost](https://mattermost.com/) plugin that increases the daily workflow of team members through a bot. This bot has onboarding automated conversations, thanks to this plugin teams will be able to ensure their progress and highlight the blocking points.

Our project has the particularity of being quite flexible and adaptable to all situations. It can also be used in a fun or educational way.

![Onboarding GIF](demo/demo.gif)

## Usage

```
git clone https://github.com/lucasnevespereira/onboarding-plugin
```

```
cd onboarding-plugin
make dist
```

Find a <b>tar.gz</b> file in the `/dist` folder.

Now you have your plugin 🙂

#### Upload plugin to Mattermost

Log into your mattermost account as sysadmin

-   Go to system console
-   Under Plugin management upload the <b>tar.gz</b> file
-   Enable plugin

# Todo

-   <b>Fetch `settings` data from `plugin.json`</b>

-   <b>Configure internal cronjob to be customisable</b>

-   <b>Add test coverage for the plugin.</b> <br>
    <small>You can start [here](server/plugin_test.go)</small>

## Webhook Request

### Prerequesites

-   Create a incoming webhook in `Integrations`
-   Get the Webhook endpoint and add it to `demo/webhook-request.sh`
-   Install watch cmd

```
brew install watch
```

### Example request

To send a bot message (webhook) every 5 seconds

```
cd demo
watch -n 5 ./webhook-request.sh
```

# How to Contribute

If you want to contribute to this project please read the [Contribution Guide](CONTRIBUTING.md).

<hr>

# Support

-   [Mattermost Userguide](https://docs.mattermost.com/guides/user.html)
-   [Developer Contact](mailto:lucas.nevespereira@eemi.com)

# License

This project is under [BSD-3 License](LICENSE)
