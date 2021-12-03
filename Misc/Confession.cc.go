{{/*
        Made by Ranger (765316548516380732)

    Trigger Type: `Command`
    Trigger: `Confess`
©️ Ranger 2021
MIT License
*/}}


{{/* Configuration values start */}}
{{$ConfessionChannel := 916414369716375552}}
{{$ConfessionLog := 916414381242339338}}
{{/* Configuration values end */}}

{{/* Only edit below if you know what you're doing (: rawr */}}

{{sendMessage $ConfessionChannel (cembed
            "author" (sdict "name" (print "Confession # " (dbIncr 1 "confessionsCount" 1)))
            "description" (print .StrippedMsg)
            "thumbnail" (sdict "url" "https://static.thenounproject.com/png/886364-200.png")
            "color" 3553599
            "timestamp" currentTime
            )}}
{{$EncryptedUser := (print (reReplace `([0-9])` (toString .User.ID) `($1[;s] $1-[74/)e$^`))}}
{{sendMessage $ConfessionLog (cembed
            "author" (sdict "name" (print "Confession # " (dbGet 1 "confessionsCount").Value))
            "description" (print .StrippedMsg "\n\nThe encrypted UserID is\n```" $EncryptedUser "```")
            "color" 3553599
            "timestamp" currentTime
            )}}
{{deleteTrigger 0}}
{{sendDM (print "Your confession has been sent! It has been logged with the encrypted userID\n```(7[;s] 7-[74/)e$^(6[;s] 6-[74/)e$^(5[;s] 5-[74/)e$^(3[;s] 3-[74/)e$^(1[;s] 1-[74/)e$^(6[;s] 6-[74/)e$^(5[;s] 5-[74/)e$^(4[;s] 4-[74/)e$^(8[;s] 8-[74/)e$^(5[;s] 5-[74/)e$^(1[;s] 1-[74/)e$^(6[;s] 6-[74/)e$^(3[;s] 3-[74/)e$^(8[;s] 8-[74/)e$^(0[;s] 0-[74/)e$^(7[;s] 7-[74/)e$^(3[;s] 3-[74/)e$^(2[;s] 2-[74/)e$^```")}}