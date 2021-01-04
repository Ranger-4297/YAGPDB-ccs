{{/*
        Made by DZ#6669 (438789314101379072) (Removed due to change in ToS by Disboard)
        Made by WickedWizard#3588 (Removed due to change in ToS by Disboard)
        Updated by Rhyker (779096217853886504)

    Trigger Type: `RegEx`
    Trigger: `\A!d\sbump(?:\s+|\z)`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$bumpNotifChannel := 787957757022175232}} {{/* channel to send the message notifying users it's possible to bump again */}}
{{$bumpPing := "<@&784132355379036194>"}} {{/* role to ping when it's possible to bump once again */}}
{{$bumpChannel := 787957898872487956}} {{/* Channel to mention in the notification for users to have a quick portal */}}
{{$bumpLog := 794365711614345267}}
{{/* Configuration values end *}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{if gt ( toInt ( currentTime.UTC.Format "15" ) ) 12 }}
{{end}}

{{if .ExecData}}
    {{sendMessageNoEscape $bumpNotifChannel (complexMessage 
            "content" $bumpPing
            "embed" (cembed
            "title" "DISBOARD: The public server list"
                "url" (print "https://disboard.org/server/" .Guild.ID)
            "description" (print "You can now bump us in <#" $bumpChannel "> using `!d bump`!\nWhy not leave us a review __[here](https://disboard.org/dashboard/reviews)__!")
            "thumbnail" (sdict "url" "https://disboard.org/images/bot-command-image-thumbnail.png")
            "color" 4436910
            ))}}
{{else}}
    {{if not (dbGet 0 "bump")}}
        {{dbSetExpire 0 "bump" 1 7200}}
        {{execCC .CCID nil 7200 "data"}}
            {{execAdmin "clean" 1 302050872383242240}}
        {{sendMessage nil (cembed
            "title" "DISBOARD: The public server list"
                "url" (print "https://disboard.org/server/" .Guild.ID)
            "description" (print .User.Mention ",\nBump done :thumbsup:\nCheck it on DISBOARD: [https://disboard.org/](https://disboard.org/server/" .Guild.ID ")")
            "image" (sdict "url" "https://disboard.org/images/bot-command-image-bump.png")
            "color" 4436910
            )}}
        {{sendMessage $bumpLog (cembed
            "title" "DISBOARD: The public server list"
                "url" (print "https://disboard.org/server/" .Guild.ID)
            "description" (print .User.Mention "Has bumped our server\nDate is " (joinStr " " ((currentTime.Format "02/01/2006"))) "\nTime is " ( joinStr " " (( currentTime.Add 0).Format "15:04 GMT")))
            "image" (sdict "url" "https://disboard.org/images/bot-command-image-bump.png")
            "color" 4436910
            )}}
    {{else}}
            {{execAdmin "clean" 1 302050872383242240}}
        {{sendMessage nil (cembed
            "title" "DISBOARD: The public server list"
                "url" (print "https://disboard.org/server/" .Guild.ID)
            "description" (print .User.Mention ", Please wait another " (((dbGet 0 "bump").ExpiresAt.Sub currentTime).Round .TimeSecond) " until you can bump the server again")
            "thumbnail" (sdict "url" "https://disboard.org/images/bot-command-image-thumbnail-error.png")
            "color" 13849427
            )}}
    {{end}}
{{end}}
