{{/*
        Made by Rhyker (779096217853886504)

    Trigger Type: `Kick DM`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$LogChannel := 784132358085017604}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{if gt ( toInt ( currentTime.UTC.Format "15" ) ) 12}}
{{end}}

{{$member := .Member}}
{{$user := .User}}
{{$args := parseArgs 0 " " (carg "member" "target")}}
{{if $args.IsSet 0}}
    {{$member = $args.Get 0}}
    {{$user = $member.User}}
{{end}}


{{$KickDM := cembed
            "author" (sdict "icon_url" ($user.AvatarURL "1024") "name" (print $user.String " (ID " $user.ID ")"))
            "description" (print "**Server:** " .Guild.Name "\n**Action:** `Kick`\n**Reason: **" .Reason ".")
            "thumbnail" (sdict "url" (print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon ".gif"))
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3553599
            }}
{{sendDM $KickDM}}

{{$Log := cembed
            "author" (sdict "icon_url" (.Author.AvatarURL "1024") "name" (print .Author.String " (ID " .Author.ID ")"))
            "description" (print ":boot: **Kicked:** *" $user.Username "#" $user.Discriminator "* `(ID " $user.ID ")`\n:receipt: **Channel:** <#" .Channel.ID ">\n:page_facing_up: **Reason:** " .Reason "\n:clock12: **Time:** " ( joinStr " " (( currentTime.Add 0).Format "15:04 GMT")))
            "thumbnail" (sdict "url" ($user.AvatarURL "256"))
            "color" 7372944
            }}
{{sendMessage $LogChannel $Log}}
