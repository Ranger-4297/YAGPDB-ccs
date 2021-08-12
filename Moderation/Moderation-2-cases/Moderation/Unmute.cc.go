{{/* 
        Case system by Maverick Wolf (549820835230253060)
        Made & modified by Ranger (765316548516380732)

    **Tools & Util > Moderation > Mute > Mute DM**
©️ Dynamic 2021
MIT License
*/}}


{{/*        Notes
    Please be aware that even though the original custom commands had the time, I removed it from this for NOW. They may be added back if I can get it to show the current time, not the time in GMT.
*/}}

{{/* Configuration values start */}}
{{$LogChannel := 784132358085017604}} {{/* Modlog channel */}}
{{$dm := 1}} {{/* change to 0 if you don't wanna DM the offender about the moderation action */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}


{{$icon := ""}}
{{$name := printf "%s (%d)" .Guild.Name .Guild.ID}}
{{if .Guild.Icon}}
	{{$ext := "webp"}}
	{{if eq (slice .Guild.Icon 0 2) "a_"}} {{$ext = "gif"}} {{end}}
	{{$icon = printf "https://cdn.discordapp.com/icons/%d/%s.%s" .Guild.ID .Guild.Icon $ext}}
{{end}}

{{$case_number := (toInt (dbIncr 77 "cv" 1))}}
{{$action := .ModAction.Prefix}}
{{$a := 0}}

{{if eq $action "Muted" "Unmuted"}}
    {{$a = (sub (len .ModAction.Prefix) 1)}}
    {{else}}
    {{$a = (sub (len .ModAction.Prefix) 2)}}
{{end}}

{{$title := (slice .ModAction.Prefix 0 $a)}}
{{$id := .User.ID}}

{{$channel := $LogChannel}}
{{if .Channel.ID}}
    {{$channel = .Channel.ID}}
{{end}}

{{if $dm}}
    {{$UnmuteDM := cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print .User.String " (ID " .User.ID ")"))
            "description" (print "**Server:** " .Guild.Name "\n**Action:** `Unmute`\n**Reason: **" (joinStr " " (split (reReplace `Automoderator:` .Reason "<:Bot:787563190221406259>:") "\n")))
            "thumbnail" (sdict "url" $icon)
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3553599
            }}
    {{sendDM $UnmuteDM}}
{{end}}

{{$x := sendMessageRetID $LogChannel (cembed
            "author" (sdict "icon_url" (.Author.AvatarURL "1024") "name" (print .Author.String " (ID " .Author.ID ")"))
            "description" (print "<:TextChannel:800978104105304065> **Case number** " $case_number "\n<:Management:788937280508657694> **Who:** " .User.Mention " `(ID " .User.ID ")`\n<:Metadata:788937280508657664> **Action:** `Unmute`\n<:Assetlibrary:788937280554926091> **Channel:** <#" $channel ">\n<:Manifest:788937280579698728> **Reason:** " (joinStr " " (split (reReplace `Automoderator:` .Reason "<:Bot:787563190221406259>:") "\n")) "\n:clock12: **Time:** " ( joinStr " " (( currentTime.Add 0).Format "15:04 GMT")))
            "thumbnail" (sdict "url" (.User.AvatarURL "256"))
            "color" 6473311
            )}}

{{$Response := sendMessageRetID nil (cembed
            "author" (sdict "icon_url" (.User.AvatarURL "1024") "name" (print "Case type: Unmute"))
            "description" (print .Author.Mention " Has successfully unmuted " .User.Mention " :thumbsup:")
            "footer" (sdict "text" " ")
            "timestamp" currentTime
            "color" 3553599
            )}}
{{deleteMessage nil $Response 5}}

{{/*for viewcase*/}}
{{dbSet $case_number "viewcase" (sdict "name" .Author.Username "warnname" .User.Username "avatar" (.Author.AvatarURL "512") "reason" .Reason "userid" $id "action" (.ModAction.Prefix) "channel" $channel "msgid" $x "userdiscrim" .User.Discriminator)}}

{{/*for per user case viewing*/}}
{{dbSet $case_number $id (print "Case # **" $case_number "**\t\t**| " $title " Reason:** `" .Reason "`")}}

{{/* for delete case*/}}
{{dbSet $case_number "userid" (str $id)}}