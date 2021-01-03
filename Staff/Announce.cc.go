{{/*
        Made by Rhyker/Ranger (779096217853886504)

    Trigger Type: `Command`
    Trigger: `Announce`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$announcementChannel := 784132357002625047}} {{/* Channel to send the announcement */}}
{{$announcementRole := 784132355379036196}} {{/* Announcement role ping */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{sendMessageNoEscape $announcementChannel (complexMessage
    "content" (mentionRoleID $announcementRole)
        "embed" (cembed
            "title" (print .User.Username " Has made an announcement!")
                "url" (print "https://discord.com/channels/" .Guild.ID "/" $announcementChannel)
            "description" (print .StrippedMsg)
            "thumbnail" (sdict "url" (print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon ".gif"))
            "image" (sdict "url" "https://cdn.discordapp.com/attachments/784132357002625047/789191521689534494/1.png")
            "color" 16761035 ))
            }}
