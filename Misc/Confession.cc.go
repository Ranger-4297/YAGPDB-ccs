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
{{$LetterList := (shuffle (cslice "A" "a" "B" "b" "C" "c" "D" "d" "E" "e" "F" "f" "G" "g" "H" "h" "I" "i" "J" "j" "K" "k"))}}
{{$CharShuffle1 := (shuffle (cslice "[" "]" "{" "}"))}}
{{$CharShuffle2 := (shuffle (cslice "!" "£" "." "%" "@" ":" ";" "*" "^"))}}

{{$a := (randInt 11 98)}}
{{$b := (index $CharShuffle1 (randInt (len $CharShuffle1)))}}
{{$c := (index $CharShuffle2 (randInt (len $CharShuffle2)))}}
{{$d := (index $LetterList (randInt (len $LetterList)))}}
{{$e := (index $CharShuffle1 (randInt (len $CharShuffle1)))}}
{{$f := (print " ")}}
{{$g := (print "$1")}}
{{$h := (index $CharShuffle2 (randInt (len $CharShuffle2)))}}
{{$i := (index $LetterList (randInt (len $LetterList)))}}
{{$j := (randInt 1 9)}}
{{$k := (index $CharShuffle2 (randInt (len $CharShuffle2)))}}
{{$l := (index $LetterList (randInt (len $LetterList))) (index $LetterList (randInt (len $LetterList)))}}
{{$m := ""}}
{{$picker := randInt 1 3}}
{{if eq $picker 1}}
	{{$m = (index $CharShuffle2 (randInt (len $CharShuffle2)))}}
{{else}}
	{{$m = (index $CharShuffle2 (randInt (len $CharShuffle2))) (index $CharShuffle2 (randInt (len $CharShuffle2)))}}
{{end}}

{{$regex := (print $a $b $c $d $e $f $g $h $i $j $k $l $m)}}
{{sendMessage $ConfessionLog (cembed
            "author" (sdict "name" (print "Confession # " (dbGet 1 "confessionsCount").Value))
            "description" (print .StrippedMsg "\n\nThe encrypted UserID is\n```" (reReplace `([0-9])` (toString .User.ID) $regex) "```")
            "color" 3553599
            "timestamp" currentTime
            )}}
{{deleteTrigger 0}}
{{sendDM (print "Your confession has been sent! It has been logged with the encrypted userID\n```" (reReplace `([0-9])` (toString .User.ID) $regex) "```")}}