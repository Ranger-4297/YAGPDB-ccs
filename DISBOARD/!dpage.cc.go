{{/*
        Made by Rhyker (779096217853886504)
        Modified by WickedWizard#3588
        
    Trigger Type: `Exact`
    Trigger: `!d page`
©️ Dynamic 2021
MIT License
*/}}


{{/* Only edit below if you know what you're doing (: rawr */}}

{{execAdmin "clean" 1 302050872383242240}}
{{sendMessage nil (cembed
            "title" "DISBOARD: The public server list"
                "url" (print "https://disboard.org/server/" .Guild.ID)
            "description" (print "https://disboard.org/server/" .Guild.ID)
            "thumbnail" (sdict "url" "https://disboard.org/images/bot-command-image-thumbnail.png")
            "color" 4436910
            )}}
