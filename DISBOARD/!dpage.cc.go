{{/*
        Made by Rhyker (779096217853886504)

    Trigger Type: `Exact`
    Trigger: `\A!d\sbump(?:\s+|\z)`
©️ Dynamic 2021
MIT License
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{execAdmin "clean" 1}}
{{sendMessage nil (cembed
            "title" "DISBOARD: The public server list"
                "url" (print "https://disboard.org/server/" .Guild.ID)
            "description" (print "https://disboard.org/server/" .Guild.ID)
            "thumbnail" (sdict "url" "https://disboard.org/images/bot-command-image-thumbnail.png")
            "color" 4436910
            )}}
