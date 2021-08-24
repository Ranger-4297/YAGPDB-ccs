{{/* 
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `whenbump`
©️ Dynamic 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$bumpChannel := 787957898872487956 }} {{/* Channel to mention in the notification for users to have a quick portal */}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{ $embed := cembed
            "title" "DISBOARD: The public server list"
                "url" (print "https://disboard.org/server/" .Guild.ID)
            "description" (print "You can bump us in <#" $bumpChannel "> using `!d bump` in " (((dbGet 0 "bump").ExpiresAt.Sub currentTime).Round .TimeSecond) "!\nWhy not leave us a review __[here](https://disboard.org/dashboard/reviews)__!")
            "thumbnail" (sdict "url" "https://disboard.org/images/bot-command-image-thumbnail.png")
            "color" 4436910
            }}

{{if dbGet 0 "cooldown" }}
    {{/* Leave blank */}}
{{else}}
    {{dbSetExpire 0 "cooldown" "cooldown" 600}}
    {{sendMessage nil $embed }}
{{end}}
