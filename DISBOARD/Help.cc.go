{{/*
        Made by Ranger (779096217853886504)
        Credits: WickedWizard#3588 (719421577086894101)

    Trigger Type: `Regex`
    Trigger: `\A!d\shelp(?:\s+|\z)`
©️ Dynamic 2021
MIT License
*/}}


{{execAdmin "clean" 1 302050872383242240}}
{{sendMessage nil (cembed
            "title" "DISBOARD: The public server list"
                "url" (print "https://disboard.org/server/" .Guild.ID)
            "description" "Hi! I am a bot for DISBOARD (https://disboard.org) 🤖\n\n__**COMMAND LIST**__\n\n`!d help`: This!\n`!d bump`: Bump this server\n`!d page`: Get server page link\n`!d invite [channel`: Change the instant invite to this channel. If [channel] is specified, create instant invite for that channel (*Admin only*).\n\n    **How do I add my server to DISBOARD?**\n\n1. Login to the DISBOARD website\n2. Go to Dashboard\n3. Click ”Click add new server”\nFill out your server and save it. You will be redirected to Discord's authorization screen. If not, click the ”Add bot” button on the server edit page.\n\n   *Support server*: https://https://discord.gg/kbQMsHZVwp"
            "thumbnail" (sdict "url" "https://disboard.org/images/bot-command-image-thumbnail.png")
            "color" 4436910
            )}}
