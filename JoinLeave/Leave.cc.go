{{/*
        Made by Ranger (779096217853886504)

    Trigger Type: `Leave Message in channel`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$Log:= 784132358085017602}} {{/* Join log channel */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing except for the 21st line. Add your ban appeal link at the end. (: rawr */}}

{{$logEmbed := cembed
            "author" (sdict "url" (.User.AvatarURL "4096") "name" "User Left" "icon_url" (.User.AvatarURL "1024"))
            "description" (print  .User.Mention "\n[" .User.String "]" " : " "`" .User.ID "`" "\nAccount created " "**" currentUserAgeHuman "** ago" "\nWe now have" "** " .Guild.MemberCount " **" "members")
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 16711680
            }}
{{sendMessage $Log $logEmbed}}
{{$leaveEmbed := cembed
            "author" (sdict "url" (.User.AvatarURL "4096") "name" "User left!" "icon_url" (.User.AvatarURL "1024"))
            "description" (print "Welp " .User.String " Left " .Guild.Name "~~ :( Guess they were cool enough~~\nWe now have `" .Guild.MemberCount "` members!")
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 16711680
            }}
{{sendMessage nil $leaveEmbed}}
