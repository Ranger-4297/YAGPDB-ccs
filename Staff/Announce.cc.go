{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Announce`
©️ Ranger 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$announcementChannel := 784132357002625047}} {{/* Channel to send the announcement */}}
{{$announcementRole := 784132355379036196}} {{/* Announcement role ping */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$icon := ""}}
{{$name := printf "%s (%d)" .Guild.Name .Guild.ID}}
{{if .Guild.Icon}}
    {{$ext := "webp"}}
    {{if eq (slice .Guild.Icon 0 2) "a_"}}
        {{$ext = "gif"}}
    {{end}}
        {{$icon = printf "https://cdn.discordapp.com/icons/%d/%s.%s" .Guild.ID .Guild.Icon $ext}}
{{end}}

{{sendMessageNoEscape $announcementChannel (complexMessage
    "content" (mentionRoleID $announcementRole)
        "embed" (cembed
            "title" (print .User.Username " Has made an announcement!")
                "url" (print "https://discord.com/channels/" .Guild.ID "/" $announcementChannel)
            "description" (print .StrippedMsg)
            "image" (sdict "url" "https://cdn.discordapp.com/attachments/784132357002625047/789191521689534494/1.png")
            "thumbnail" (sdict "url" $icon)
            "color" 16761035 ))
            }}
