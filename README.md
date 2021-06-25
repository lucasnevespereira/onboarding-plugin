# Onboarding Mattermost Plugin

![Logo](assets/logo.png)

# Description

#### FR

Projet de workshop réalisé par 5 étudiants de l'EEMI. Réalisation d'un plugin pour Mattermost permettant un suivi optimal des projets par le biai d'un chatbot. Le chatbot permet un onboarding automatisé, grâce à cela le manager pourra s'assurer de l'avancée de son équipe et de mettre en évidence les points bloquants.

Notre projet a pour particularité d'être très flexible et adadptable à toutes les situations ou projets. Il peut ainsi être utilisé de façon ludique ou éducative.

#### EN

Workshop project done by 5 students at [EEMI](https://eemi.com). We've created a [Mattermost](https://mattermost.com/) plugin that increases the daily workflow of team members through a bot. This bot has onboarding automated conversations, thanks to this plugin teams will be able to ensure their progress and highlight the blocking points.

Our project has the particularity of being quite flexible and adaptable to all situations. It can also be used in a fun or educational way.

![Onboarding GIF](assets/demo.gif)

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

-   <b>Integrate Webhook call from plugin internal code </b>

## Example Webhook Request

To send a bot message (webhook) every 5 seconds

```
cd demo
watch -n 5 ./webhook-request.sh
```

### Prerequesites

```
brew install watch
```

-   <b>Use custom question and answers from `settings` in `plugin.json`</b>

-   <b>Integrate internal scheduler or cronjob to the plugin</b>

-   <b>Add test coverage for the plugin.</b> <br>
    <small>You can start [here](server/plugin_test.go)</small>

# Support

-   [Mattermost Userguide](https://docs.mattermost.com/guides/user.html)
-   [Developer Contact](mailto:lucas.nevespereira@eemi.com)

# License

This project is under [BSD-3 License](LICENSE)
