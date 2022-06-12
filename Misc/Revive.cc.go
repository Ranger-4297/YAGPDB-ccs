{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Revive`

    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}

{{/* Configuration values start */}}
{{$activityPing := 799160232823226368}} {{/* RoleID of the role to ping when reviving chat */}}
{{$Cooldown := 3600}} {{/* Cooldown must be in seconds https://www.convertworld.com/en/time/seconds.html */}}
{{/* Configuration values end */}}

{{if $cooldown := dbGet 0 "cooldown" }}
    {{sendMessage nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print "Sorry, you're on cooldown"))
            "description" (print ":warning: This command is still on cooldown for: " (humanizeDurationSeconds ($cooldown.ExpiresAt.Sub currentTime)) )
            "color" 14043208
            )}}
{{else}}
	{{dbSetExpire 0 "cooldown" "cooldown" $Cooldown}}
	{{sendMessageNoEscape nil (complexMessage
            "content" (print "<@&" $activityPing ">")
            "embed" (cembed
            "description"  (print .User.Mention " summons you with an activity ping!")
            "color" 4436910
            ))}}
{{end}}
