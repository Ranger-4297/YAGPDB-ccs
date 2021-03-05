{{/*
        Made by Ranger (779096217853886504)

    Trigger Type: `Unmute DM`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$LogChannel := 784132358085017604}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$icon := ""}}
{{$name := printf "%s (%d)" .Guild.Name .Guild.ID}}
{{if .Guild.Icon}}
	{{$ext := "webp"}}
	{{if eq (slice .Guild.Icon 0 2) "a_"}} {{$ext = "gif"}} {{end}}
	{{$icon = printf "https://cdn.discordapp.com/icons/%d/%s.%s" .Guild.ID .Guild.Icon $ext}}
{{end}}

{{if gt ( toInt ( currentTime.UTC.Format "15" ) ) 12}}
{{end}}

{{$channel := $LogChannel}}
{{if .Channel.ID}}
    {{$channel = .Channel.ID}}
{{end}}

{{$UnmuteDM := cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print .User.String " (ID " .User.ID ")"))
            "description" (print "**Server:** " .Guild.Name "\n**Action:** `Unmute`\n**Reason: **" .Reason ".")
            "thumbnail" (sdict "url" $icon)
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3553599
            }}
{{sendDM $UnmuteDM}}

{{$Log := cembed
            "author" (sdict "icon_url" (.Author.AvatarURL "1024") "name" (print .Author.String " (ID " .Author.ID ")"))
            "description" (print ":loud_sound: **Unmute:** *" .User.Username "#" .User.Discriminator "* `(ID " .User.ID ")`\n:receipt: **Channel:** <#" $channel ">\n:page_facing_up: **Reason:** " .Reason "\n:clock12: **Time:** " ( joinStr " " (( currentTime.Add 0).Format "15:04 GMT")))
            "thumbnail" (sdict "url" (.User.AvatarURL "256"))
            "color" 6473311
            }}
{{sendMessage $LogChannel $Log}}
