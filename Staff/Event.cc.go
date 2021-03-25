{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Event`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$eventAnnouncementChannel := 784132357002625047}} {{/* Channel to send the event announcement */}}
{{$eventRole := 784132355379036196}} {{/* Event announcement role ping */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{sendMessageNoEscape $eventAnnouncementChannel (complexMessage
    "content" (mentionRoleID $eventRole)
        "embed" (cembed
            "title" (print .User.Username " Has made an event notif!")
                "url" (print "https://discord.com/channels/" .Guild.ID "/" $eventChannel)
            "description" (print .StrippedMsg)
            "thumbnail" (sdict "url" (print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon ".gif"))
            "image" (sdict "url" "https://cdn.discordapp.com/attachments/784132357002625047/789191517704290324/2.png")
            "color" 16761035 ))
            }}
