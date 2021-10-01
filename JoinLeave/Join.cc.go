{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Join Message in channel`
©️ Ranger 2021
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
            "color" 3247335
            "timestamp" currentTime
            }}
    {{sendMessage $Log $logEmbed}}
    {{$welcomeEmbed := cembed
            "author" (sdict "url" (.User.AvatarURL "4096") "name" "User Joined!" "icon_url" (.User.AvatarURL "1024"))
            "description" (print "Hey there,  " .User.String "! Welcome to " .Guild.Name "!\nWe now have `" .Guild.MemberCount "` members!")
            "color" 65419
            "timestamp" currentTime
            }}
    {{sendMessage nil $welcomeEmbed}}
{{end}}
