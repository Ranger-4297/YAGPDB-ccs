{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Join Message in channel`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$Log:= 765316548516380732}} {{/* Join log channel */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{if .UsernameHasInvite}}
{{$silent := (execAdmin "ban" .User.ID "ad blocked")}}
{{else}}
{{$logEmbed := cembed
            "author" (sdict "url" (.User.AvatarURL "4096") "name" "User Joined" "icon_url" (.User.AvatarURL "1024"))
            "description" (print  .User.Mention "\n[" .User.String "]" " : " "`" .User.ID "`" "\nAccount created " "**" currentUserAgeHuman "** ago" "\nWe now have" "** " .Guild.MemberCount " **" "members")
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3247335
            }}
{{sendMessage $Log $logEmbed}}
{{$welcomeEmbed := cembed
            "author" (sdict "url" (.User.AvatarURL "4096") "name" "User Joined!" "icon_url" (.User.AvatarURL "1024"))
            "description" (print "Hello " .User.String "! Welcome to " .Guild.Name "! Check the DM I sent you!\nWe now have `" .Guild.MemberCount "` members!")
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 65419
            }}
{{sendMessage nil $welcomeEmbed}}
{{end}}
