{{/*
        Made by Maverick Wolf (549820835230253060)
        Edited by ranger_4297 (765316548516380732)
        Credit to LemmeCry (664118444739919882)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(delcase|deletecase|clearcase|dc|clc)(\s+|\z)`

Repo: https://github.com/Maverick-Wolf/yagpdb-mave
MIT License
*/}}


{{/* Configuration values start */}}
{{$roles := cslice }} {{/* Add your staff role ID's here */}}
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
    {{$args := parseArgs 1 "correct usuage is `-delcase <CaseID>`" (carg "int" "case number")}}
    {{$caseno := ($args.Get 0)}}
    {{$id := (toInt (dbGet $caseno "userid").Value)}}
    {{dbDel $caseno "viewcase"}}
    {{dbDel $caseno $id}}
    {{dbDel $caseno "userid"}}
    {{$embed := sendMessage nil (cembed
        "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" .User.Username)
        "description" (print "Deleted the case")
        "color" 0x36393f
        )}}
{{end}}