{{/*
        Made by Maverick Wolf (549820835230253060)
        Edited by Ranger (765316548516380732)
        Credit to LemmeCry (664118444739919882)

    Trigger Type: `Regex`
    Trigger: `\A(-|<@!?204255221017214977>\s*)(cases|allcase)(\s+|\z)`

Repo: https://github.com/Maverick-Wolf/yagpdb-mave
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
            "color" 0x36393f
            )}}
{{else}}
    {{$args := parseArgs 1 "correct usuage is `-cases <User/Mention> [Page]`" (carg "member" "target user") (carg "int" "page number")}}
    {{$user := ($args.Get 0).User}}
    {{$id := $user.ID}}
    {{$display := ""}}
    {{$page := 1}} {{/* Default page to start at */}}
    {{if ($args.IsSet 1)}}
        {{if ge ($args.Get 1) 1}}
            {{$page = ($args.Get 1)}}
        {{else}}
            {{$page = 1}}
        {{end}}
    {{end}}
    {{$skip := mult (sub $page 1) 10}}
    {{$cases := dbTopEntries $id 10 $skip}}
    {{range $cases}}
        {{- $value := str .Value }} 
        {{- $display = joinStr "" $display "\n" $value -}} 
    {{else}}
        {{$display = "`No cases exist on this page`"}}
    {{end}}
    {{$id := sendMessageRetID nil (cembed
            "title" "Cases" 
            "author" (sdict "name" (print $user.Username " (" $user.ID ")") "icon_url" ($user.AvatarURL "512"))
            "description" $display
            "color" 0x36393f
            "footer" (sdict "text" (print "Page " $page))
            )}}
{{end}}