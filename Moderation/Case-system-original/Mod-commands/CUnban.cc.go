{{/*
        Made by ranger_4297 (765316548516380732)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(c(ase)?-?unban)(\s+|\z)`

    ©️ RhykerWells 2020-Present
    GNU, GPLV3 License
    Repository: https://github.com/Ranger-4297/YAGPDB-ccs
*/}}


{{/* Configuration values start */}}
{{$roles := cslice 885177438135517214}} {{/* Add your staff role ID's */}}
{{$LogChannel := 838432051094880306}} {{/* Modlog channel */}}
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
            "color" 0x36393f
            )}}
{{else}}
    {{$args := parseArgs 1 "` -CUnban <CaseID> [Reason:Tex]` " (carg "int" "case number") (carg "string" "reason")}}
    {{$a := ""}}
    {{$userid := ""}}
    {{$uavatar := ""}}
    {{$action := ""}}
    {{$b := ($args.Get 0)}}
    {{with (dbGet $b "viewcase")}}	
        {{$a = sdict .Value}}
        {{$userid = $a.userid}}
        {{$uavatar = $a.uavatar}}
        {{$action = $a.action}}
        {{if eq $action "Banned"}}
            {{/* If no reason */}}
            {{$reason := (print "No reason specified")}}
            {{/* If reason */}}
            {{if ($args.Get 1)}}
                {{$reason = (joinStr " " (slice $.CmdArgs 1))}}
            {{end}}
            {{/* Log & DM messages */}}
            {{$LogEmbed := (cembed
            "author" (sdict "icon_url" ($.User.AvatarURL "1024") "name" (print $.User.String " (ID " $.User.ID ")"))
            "description" (print "<:Management:788937280508657694> **Who:** <@" $userid "> `(ID " $userid ")`\n<:Metadata:788937280508657664> **Action:** `Unban`\n<:Assetlibrary:788937280554926091> **Channel:** <#" $.Channel.ID ">\n<:Manifest:788937280579698728> **Reason:** " $reason)
            "thumbnail" (sdict "url" $uavatar)
            "color" 6473311
            )}}
            {{sendMessage $LogChannel $LogEmbed}}
            {{$Response := sendMessage nil (cembed
            "author" (sdict "icon_url" $uavatar "name" (print "Case type: Unban"))
            "description" (print  $.User.Mention " Has successfully Unbanned <@!" $userid "> :thumbsup:")
            "color" 3553599
            "timestamp" currentTime
            )}}
            {{deleteMessage nil $Response 3}}
            {{$silent := execAdmin "Unban" $userid}}
        {{else}}
            {{sendMessage nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" .User.Username)
            "description" (print "Could not find the case specified. Please make sure the case number is correct, that the case is a `type: ban` case or that it's not been deleted.")
            "color" 3553599
            )}}
        {{end}}
    {{else}}
        {{sendMessage nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" .User.Username)
            "description" (print "Could not find the case specified. Please make sure the case number is correct or the case has not been deleted")
            "color" 3553599
            )}}
    {{end}}
{{end}}