{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A!d\shelp(?:\s+|\z)`
©️ Ranger 2021
MIT License
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{execAdmin "clean" 1 302050872383242240}}
{{sendMessage nil (cembed
            "title" "DISBOARD: The public server list"
                "url" (print "https://disboard.org/server/" .Guild.ID)
            "description" (print "Hi! This is a custom command for Disboard! (https://disboard.org) 🤖\n\n__**COMMAND LIST**__\n\n`!d help`: This!\n`!d bump`: Bump this server\n`!d page`: Get the servers' page link\n`!d invite [channel`: Change the instant invite to this channel. If [channel] is specified, create instant invite for that channel (*Admin only*).\n\n**How do I add my server to DISBOARD?**\n\n1. Login to the DISBOARD website\n2. Go to Dashboard\n3. Click ”Click add new server”\nFill out your server and save it. You will be redirected to Discord's authorization screen. If not, click the ”Add bot” button on the server edit page.")
            "thumbnail" (sdict "url" "https://disboard.org/images/bot-command-image-thumbnail.png")
            "color" 4436910
            )}}
