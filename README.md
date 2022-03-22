# golain - Lain: CLI-to-Discord Utility written in Golang

## Install
- Prerequisite: [Install Go](https://go.dev/doc/install)
- Pull this repo: `git pull https://github.com/sabi/golain.git` or [Download the zip](https://github.com/sabi/golain/archive/refs/heads/main.zip)
- Enter the directory or unzip and change directory
- `go install` or `go build` to compile the binary

## Usage
- `golain msg -c DISCORD_CHANNEL "Your message goes here"`
- `golain msg -c sabi_engineers "You're doing a great job!"`

## Add New Discord Webhook
- `golain add -c CHANNEL_NAME -w WEBHOOK`

## How to Create a Discord Webhook
- Open your Server Settings, then go to Integrations
![image](https://user-images.githubusercontent.com/49737728/146336626-2e511660-dd73-4fe9-8b50-4a211d27d2a3.png)

- Create a new webhook for the appropriate channel, then copy the webhook to clipboard
![image](https://support.discord.com/hc/article_attachments/1500000455142/Screen_Shot_2020-12-15_at_4.45.52_PM.png)
