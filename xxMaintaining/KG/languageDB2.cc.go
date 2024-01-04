{{$languageDB := (dbGet 0 "languageDB").Value}}
{{$languages := sdict
"arabic" (dict
	1 "لا يوجد تحالف؟ اكتب \"skip\""
	2 "مذهل! <@!> لقد استلمت للتو علامة الخادم الخاصة بك! انتقل الآن إلى <#> <1/5>"
	3 "<@!> أدخل علامة التحالف الخاصة بك هنا <2/5>"
	4 "ان هذا رائع! <@!> لديك الآن علامة تحالف! تابع وانتقل إلى <#> <2/5>"
	5 "<@!> اكتب اسم لعبة الشخصية الخاصة بك هنا! <3/5>"
	6 "ممتاز! <@!> لقد قمت بتحديث اسم العرض الخاص بك إلى اسم لعبة الشخصية الخاصة بك! خطوتان أخريان، تابع إلى <#> <3/5>"
	7 "<@!> حدد رتبة تحالفك داخل اللعبة. <4/5>"
	8 "لا يمكن إضافة رد فعل للمستخدم الذي قام بحظر الروبوت. تم تحديث اللقب"
	9 "الرجاء إدخال علامة مكونة من 3-4 أرقام"
	10 "الرجاء إدخال علامة رقمية"
	11 "الرجاء إدخال علامة تحالف مكونة من 3-4 أحرف"
	12 "الرجاء عدم استخدام أحرف خاصة"
	13 "الرجاء إدخال اسم مستخدم مكون من 3 إلى 15 حرفًا"
	14 "لنبدأ <@!>، انتقل إلى <#> <0/5>"
	15 "حسنًا <@!>، دعنا نجهزك للانضمام والتحقق!"
	16 "تفضل <@!>، لقد بدأنا للتو! <0/5>"
	17 "أخيراً!! <@!> أنت هنا!! الخطوة الأخيرة إلى الجانب الآخر!! <5/5>"
	18 "ها ها ها ها! لقد فعلتها! <@!> مرحبًا بك في الجانب الآخر!"
	19 "أحسنت! <@!> استعد الآن للخطوة الأخيرة⁠ <#> <5/5>"
	20 "### مرحبًا بك <@!> في KG World - WOR 🌀\n> يرجى اتباع خطوات التحقق الخمس، المشفرة بواسطة <@!765316548516380732>\n> - سيستغرق هذا أقل من دقيقة.\n> - للتغيير اللغة 👉 <#L>\n\n~ *الجانب الآخر في انتظارك* 💫\n### **👇 انقر على ✅ للبدء.**"
)
"korean" (dict
	1 "얼라이언스가 없나요? \"skip\" 를 입력하세요."
	2 "엄청난! <@!> 방금 서버 태그를 받았습니다! 이제 <#>로 이동하세요 <1/5>"
	3 "<@!> 여기에 동맹 태그를 입력하세요 <2/5>"
	4 "이것은 멋지다! <@!> 이제 동맹 태그를 가지게 되었습니다! 계속해서 <#> 로 이동하세요 <2/5>"
	5 "<@!> 여기에 캐릭터 게임 이름을 입력하세요! <3/5>"
	6 "완벽한! <@!> 표시 이름을 캐릭터 게임 이름으로 업데이트했습니다! 두 단계만 더 진행하면 <#> 로 이동합니다 <3/5>"
	7 "<@!> 게임 내 동맹 순위를 선택하세요 <4/5>"
	8 "봇을 차단한 사용자에게는 반응을 추가할 수 없습니다. 닉네임이 업데이트되었습니다"
	9 "3~4자리 태그를 입력해주세요"
	10 "숫자 태그를 입력해주세요"
	11 "연맹 태그를 3~4자로 입력해주세요"
	12 "특수문자는 사용하지 마세요"
	13 "3~15자의 사용자 이름을 입력하세요"
	14 "<@!>를 시작하고 <#>로 진행하세요. <0/5>"
	15 "좋습니다 <@!> 님, 등록 및 인증을 받을 수 있도록 준비하겠습니다!"
	16 "<@!> 님, 이제 막 시작했습니다! <0/5>"
	17 "마지막으로!! <@!> 여기 계시네요!! 반대편으로 가는 마지막 단계!! <5/5>"
	18 "하하하! 네가 해냈어! <@!> 반대편에 오신 것을 환영합니다!"
	19 "잘하셨어요! <@!>님은 이제 마지막 단계를 준비하세요⁠ <#> <5/5>"
	20 "### KG World에 <@!> 오신 것을 환영합니다 - WOR 🌀\n> <@!765316548516380732>로 코딩된 5가지 확인 단계를 따르십시오.\n> - 이 작업은 1분도 채 걸리지 않습니다.\n> - 변경하려면 언어 👉 <#L>\n\n~ *다른 쪽이 당신을 기다립니다* 🌟\n### **👇 시작하려면 ✅를 클릭하세요.**"
)
"german" (dict
	1 "Kein Bündnis? Geben Sie \"skip\" ein"
	2 "Eindrucksvoll! <@!> Sie haben gerade Ihr Server-Tag erhalten! Fahren Sie nun mit <#> fort <1/5>"
	3 "<@!> Geben Sie hier Ihr Allianz-Tag ein <2/5>"
	4 "Das ist cool! <@!> Sie haben jetzt ein Allianz-Tag! Fahren Sie fort und fahren Sie mit <#> fort <2/5>"
	5 "<@!> Geben Sie hier Ihren Charakter-Spielnamen ein! <3/5>"
	6 "Perfekt! <@!> Sie haben Ihren Anzeigenamen auf den Namen Ihres Charakterspiels aktualisiert! Zwei weitere Schritte, fahren Sie mit <#> fort <3/5>"
	7 "<@!> wähle deinen Allianzrang im Spiel aus <4/5"
	8 "Dem Benutzer, der den Bot blockiert hat, kann keine Reaktion hinzugefügt werden. Spitzname aktualisiert"
	9 "Bitte geben Sie ein 3-4-stelliges Tag ein"
	10 "Bitte geben Sie ein numerisches Tag ein"
	11 "Bitte geben Sie einen Allianz-Tag mit 3-4 Zeichen ein"
	12 "Bitte verwenden Sie keine Sonderzeichen"
	13 "Bitte geben Sie einen 3-15 Zeichen langen Benutzernamen ein"
	14 "Beginnen wir mit <@!> und fahren mit <#> <0/5> fort"
	15 "Alles klar, <@!>, bereiten wir dich auf das Onboarding und die Verifizierung vor!"
	16 "Los geht's <@!>, wir haben gerade erst angefangen! <0/5>"
	17 "Endlich!! <@!> Du bist hier!! Letzter Schritt auf die andere Seite!! <5/5>"
	18 "Hahaha! Du hast es geschafft! <@!> Willkommen auf der anderen Seite!"
	19 "Gut gemacht! <@!> Machen Sie sich jetzt bereit für den letzten Schritt⁠ <#> <5/5>"
	20 "### Willkommen <@!> bei KG World – WOR 🌀\n> Bitte folgen Sie den 5 Verifizierungsschritten, codiert durch <@!765316548516380732>\n> – Dies dauert weniger als eine Minute.\n> – Zum Ändern Sprache 👉 <#L>\n\n~ *Die andere Seite erwartet Sie* 💫\n### **👇 Klicken Sie auf ✅, um loszulegen.**"
)
"vietnamese" (dict
	1 "Không có liên minh? Gõ \"skip\""
	2 "Tuyệt vời! <@!> Bạn vừa nhận được Thẻ máy chủ của mình! Bây giờ hãy chuyển sang <#> <1/5>"
	3 "<@!> Nhập thẻ liên minh của bạn vào đây <2/5>"
	4 "Cái này hay đấy! <@!> bạn hiện có Thẻ Liên minh! Tiếp tục và tiến tới <#> <2/5>"
	5 "<@!> nhập Tên trò chơi nhân vật của bạn vào đây! <3/5>"
	6 "Hoàn hảo! <@!> bạn đã cập nhật tên hiển thị thành Tên trò chơi nhân vật của mình! Hai bước nữa, tiến tới <#> <3/5>"
	7 "<@!> chọn thứ hạng liên minh trong trò chơi của bạn <4/5>"
	8 "Không thể thêm phản ứng cho người dùng đã chặn bot. Đã cập nhật biệt hiệu"
	9 "Vui lòng nhập thẻ 3-4 chữ số"
	10 "Vui lòng nhập thẻ số"
	11 "Vui lòng nhập thẻ liên minh 3-4 ký tự"
	12 "Vui lòng không sử dụng ký tự đặc biệt"
	13 "Vui lòng nhập tên người dùng 3-15 ký tự"
	14 "Hãy bắt đầu <@!>, tiến tới <#> <0/5>"
	15 "Được rồi <@!>, hãy chuẩn bị sẵn sàng để bạn tham gia và xác minh!"
	16 "Của bạn đây <@!>, chúng ta vừa mới bắt đầu! <0/5>"
	17 "Cuối cùng!! <@!> Bạn đang ở đây!! Bước cuối cùng sang phía bên kia!! <5/5>"
	18 "Hahaha! Bạn đã thực hiện nó! <@!> Chào mừng bạn đến với phía bên kia!"
	19 "Làm tốt! <@!> bây giờ hãy sẵn sàng cho bước cuối cùng⁠ <#> <5/5>"
	20 "### Chào mừng <@!> đến với KG World - WOR 🌀\n> Vui lòng làm theo 5 bước xác minh, được mã hóa bởi <@!765316548516380732>\n> - Quá trình này sẽ mất chưa đầy một phút.\n> - Để thay đổi ngôn ngữ 👉 <#L>\n\n~ *Phía bên kia đang chờ bạn* 💫\n### **👇 Nhấp vào ✅ để bắt đầu.**"
)
"japanese" (dict
	1 "同盟はありませんか？ \"skip\" ップ」と入力してください"
	2 " 素晴らしい！ <@!> サーバー タグを受け取りました。 <#> に進みます <1/5>"
	3 "<@!> ここに同盟タグを入力してください <2/5>"
	4 "これはカッコいい！ <@!> さんは同盟タグを獲得しました! 続行して <#> に進みます <2/5>"
	5 "<@!> ここにキャラクター ゲーム名を入力してください! <3/5>"
	6 "完璧！ <@!> 表示名をキャラクター ゲーム名に更新しました! さらに 2 つの手順を実行して、<#> に進みます <3/5>"
	7 "<@!> はゲーム内の同盟ランクを選択します <4/5>"
	8 "ボットをブロックしたユーザーにリアクションを追加することはできません。 ニックネームを更新しました"
	9 "3-4桁のタグを入力してください"
	10 "数字タグを入力してください"
	11 "3-4文字の同盟タグを入力してください"
	12 "特殊文字は使用しないでください"
	13 "3-15 文字のユーザー名を入力してください"
	14 "始めましょう <@!>、<#> <0/5> に進みます"
	15 "わかりました <@!>、オンボーディングと認証の準備をしましょう!"
	16 "どうぞ <@!>、まだ始まったばかりです! <0/5>"
	17 "ついに！！ <@!> ここにいます!! 向こう側への最終段階!! <5/5>"
	18 "ははは！ やった！ <@!> 向こう側へようこそ!"
	19 "よくやった！ <@!> は最終ステップの準備をしてください⁠ <#> <5/5>"
	20 "### <@!> KG World へようこそ - WOR 🌀\n> <@!765316548516380732> でコード化された 5 つの確認手順に従ってください。\n> - これには 1 分もかかりません。\n> - 言語を変更するには 👉 <# L>\n\n~ *反対側があなたを待っています* 💫\n### **👇 ✅ をクリックして開始してください。**"
)
}}
{{range $k, $v := $languages}}
	{{$languageDB.Set $k $v}}
{{end}}
{{dbSet 0 "languageDB" $languageDB}}