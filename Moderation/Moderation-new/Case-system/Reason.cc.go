{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `RegEx`
    Trigger: `(-|<@!?204255221017214977>\s*)((edit|new)-?reason)`

©️ Ranger 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$roles := cslice }} {{/* Add your staff role ID's */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{$check := 0}}
{{range $roles}}
    {{if eq $check 0}}
        {{if hasRoleID .}}
        {{$check =1}}
        {{end}}
    {{end}}
{{end}}

{{if eq $check 0}}
    {{sendMessage nil (cembed
            "author" (sdict "name" (print .User.Username) "icon_url" (.User.AvatarURL "512"))
            "description" (print "<:Cross:817828050938363905> I'm sorry. You don't have permission to use this command.")
            "color" 3553599
            )}}
{{else}}
    {{$args := parseArgs 1 "`edit-reason <CaseID> <Reason:Text>` " (carg "int" "case number") (carg "string" "reason")}}
    {{$a := ""}}
    {{$reason := ""}}
    {{$userid := ""}}
    {{$action := ""}}
    {{$b := ($args.Get 0)}}
    {{with (dbGet $b "viewcase")}}	
        {{$a = sdict .Value}}
        {{$reason = $a.reason}}
        {{$userid = $a.userid}}
        {{$action = $a.action}}
        {{if ($args.Get 1)}}
            {{$newreason := (joinStr " " (slice $.CmdArgs 1))}}
            {{$Response := sendMessage nil (cembed
                "author" (sdict "icon_url" ($.User.AvatarURL "1024") "name" (print $.User.String))
                "description" (print "Successfully updated the reason for case: `" $b "`\n**Old reason :** " $reason "\n**New reason:** " $newreason )
                "color" 3553599
                "timestamp" currentTime
                )}}
            {{$sdict := (dbGet $b "viewcase").Value}}
            {{$sdict.Set "reason" $newreason}}
            {{dbSet $b "viewcase" $sdict}}
            {{dbSet $b $userid (print "Case # **" $b "**\t\t**| " (dbGet $b "cases").Value " Reason:** `" $newreason "`")}}
        {{else}}
        {{sendMessage nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" .User.Username)
            "description" (print "Please specify a reason")
            "color" 3553599
            )}}
        {{end}}
    {{else}}
        {{sendMessage nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" .User.Username)
            "description" (print "Could not find the case specified, Please make sure the case number is correct or the case has not been deleted")
            "color" 3553599
            )}}
    {{end}}
{{end}}