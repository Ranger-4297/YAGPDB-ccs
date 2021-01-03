{{/*
        Made by Rhyker/Ranger (779096217853886504)

    Trigger Type: `Join Message in channrl`
©️ Dynamic 2021
MIT License
*/}}


{{if .UsernameHasInvite}}
{{ $silent := (execAdmin "ban" .User.ID "ad blocked") }}
{{else}}
{{$embed := cembed
            "author" (sdict "url" (.User.AvatarURL "4096") "name" "User Joined" "icon_url" (.User.AvatarURL "1024"))
            "description" (print  .User.Mention "\n[" .User.String "]" " : " "`" .User.ID "`" "\nAccount created " "**" currentUserAgeHuman "** ago" "\nWe now have" "** " .Guild.MemberCount " **" "members")
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3247335
            }}
{{sendMessage nil $embed}}
{{end}}
