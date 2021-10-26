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
    {{$msgid := ""}}
    {{$channel := ""}}
    {{$channel2 := ""}}
    {{$logs := ""}}
    {{$time := ""}}
    {{$b := ($args.Get 0)}}
    {{with (dbGet $b "viewcase")}}	
        {{$a = sdict .Value}}
        {{$reason = $a.reason}}
        {{$userid = $a.userid}}
        {{$action = $a.action}}
        {{$msgid = $a.msgid}}
        {{$channel = $a.channel}}
        {{$channel2 = $a.channel2}}
        {{$logs = $a.logs}}
        {{$time = $a.time}}
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
            {{if dbGet $b $userid}}
                {{dbSet $b $userid (print "Case # **" $b "**\t\t**| " (dbGet $b "cases").Value " Reason:** `" $newreason "`")}}
            {{end}}
            {{$embed := index (getMessage $channel $msgid).Embeds 0| structToSdict}}
            {{$embed.Set "description" (print "<:TextChannel:800978104105304065> **Case number:** " $b "\n<:Management:788937280508657694> **Who:** <@!" $userid "> `(ID " $userid ")`\n<:Metadata:788937280508657664> **Action:** `Warn`\n<:Assetlibrary:788937280554926091> **Channel:** <#" $channel2 ">\n<:Manifest:788937280579698728> **Reason:** " $newreason "\n<:Edit:800978104272683038> **Message Logs:** [Click Here](" $logs ")\n:clock12: **Time:** <t:" $time ":f>")}}
            {{editMessage $channel $msgid (cembed $embed)}}
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