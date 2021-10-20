{{/* 
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `unban`
©️ Ranger 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$roles := cslice 784132355534880855}} {{/* Add your staff role ID's */}}
{{$LogChannel := 784132358085017604}} {{/* Modlog channel */}}
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
    {{$args := parseArgs 1 "```Unban <User:ID> [Reason:Text]```" (carg "string" "target") (carg "string" "optional")}}
    {{$target := $args.Get 0}}
    {{execAdmin "Unban" ($args.Get 0)}}
    {{$case_number := (toInt (dbIncr 77 "cv" 1))}}
    {{$action := "Unbanned"}}
    {{$id := ($args.Get 0)}}
    {{$username := "Unknown_user"}}
    {{$userdiscrim := "0000"}}
    {{$channel := $LogChannel}}
    {{$a := 0}}
    {{if eq $action "Muted" "Unmuted"}}
        {{$a = (sub (len "Unbanned") 1)}}
    {{else}}
        {{$a = (sub (len "Unbanned") 2)}}
    {{end}}
    {{$title := (slice "Unbanned" 0 $a)}}
    {{/* If no reason */}}
    {{$reason := (print "No reason specified")}}
    {{/* If reason */}}
    {{if ($args.Get 1)}}
        {{$reason = (joinStr " " (slice .CmdArgs 1))}}
    {{end}}
    {{$x := sendMessageRetID $LogChannel (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print .User.String " (ID " .User.ID ")"))
            "description" (print "<:TextChannel:800978104105304065> **Case number** " $case_number "\n<:Management:788937280508657694> **Who:** <@" $target "> `(ID " $target ")`\n<:Metadata:788937280508657664> **Action:** `Unban`\n<:Assetlibrary:788937280554926091> **Channel:** <#" .Channel.ID ">\n<:Manifest:788937280579698728> **Reason:** " $reason)
            "color" 6473311
            )}}
    {{$Response := sendMessageRetID nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print "Case type: Unban"))
            "description" (print .User.Mention " Has successfully Unbanned <@" $target "> :thumbsup:")
            "color" 3553599
            "timestamp" currentTime
			)}}
    {{deleteMessage nil $Response 3}}
    {{/*for viewcase*/}}
    {{dbSet $case_number "viewcase" (sdict "name" .User.Username "warnname" $username "avatar" (.User.AvatarURL "512") "reason" $reason "userid" $id "action" $action "channel" $channel "msgid" $x "userdiscrim" $userdiscrim)}}
    {{/*for per user case viewing*/}}
    {{dbSet $case_number $id (print "Case # **" $case_number "**\t\t**| " $title " Reason:** `" $reason "`")}}
    {{/* for delete case*/}}
    {{dbSet $case_number "userid" (str $id)}}
{{end}}