{{/*
        Made by Ranger (765316548516380732)
        Modified by WickedWizard#3588 (719421577086894101)
        
    Trigger Type: `Regex`
    Trigger: `\A!d\sPage(?:\s+|\z)`
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
