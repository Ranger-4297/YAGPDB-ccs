{{/*
        Made by Ranger (779096217853886504)

    Trigger Type: `Warn DM`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$LogChannel := 784132358085017604}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{if gt ( toInt ( currentTime.UTC.Format "15" ) ) 12}}
{{end}}

{{$WarnDM := cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print .User.String " (ID " .User.ID ")"))
            "description" (print "**Server:** " .Guild.Name "\n**Action:** `Warn`\n**Reason: **" .Reason ".")
            "thumbnail" (sdict "url" (print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon ".gif"))
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3553599
            }}
{{sendDM $WarnDM}}

{{$Log := cembed
            "author" (sdict "icon_url" (.Author.AvatarURL "1024") "name" (print .Author.String " (ID " .Author.ID ")"))
            "description" (print ":warning: **Warned:** *" .User.Username "#" .User.Discriminator "* `(ID " .User.ID ")`\n:receipt: **Channel:** <#" .Channel.ID ">\n:page_facing_up: **Reason:** " .Reason "\n:clock12: **Time:** " ( joinStr " " (( currentTime.Add 0).Format "15:04 GMT")))
            "thumbnail" (sdict "url" (.User.AvatarURL "256"))
            "color" 16556627
            }}
{{sendMessage $LogChannel $Log}}
